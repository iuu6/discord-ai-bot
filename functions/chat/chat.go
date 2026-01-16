package chat

import (
	"bytes"
	"context"
	"discord-ai-bot/utils/config"
	"discord-ai-bot/utils/messenger"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	cfg := config.Load()

	query := strings.TrimPrefix(m.Content, "/chat ")
	if query == "" {
		messenger.Send(s, m.ChannelID, "Please provide a message after /chat", m.Reference())
		return
	}

	payload := map[string]interface{}{
		"model":      cfg.AnthropicModel,
		"max_tokens": 1024,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": query,
			},
		},
	}

	if cfg.SystemPrompt != "" {
		payload["system"] = []map[string]interface{}{
			{
				"type": "text",
				"text": cfg.SystemPrompt,
			},
		}
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling request: %v", err)
		messenger.Send(s, m.ChannelID, "Error communicating with AI", m.Reference())
		return
	}

	url := fmt.Sprintf("%s/v1/messages", strings.TrimSuffix(cfg.AnthropicBase, "/"))
	req, err := http.NewRequestWithContext(context.Background(), "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		messenger.Send(s, m.ChannelID, "Error communicating with AI", m.Reference())
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", cfg.AnthropicKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error calling Anthropic: %v", err)
		messenger.Send(s, m.ChannelID, "Error communicating with AI", m.Reference())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		messenger.Send(s, m.ChannelID, "Error communicating with AI", m.Reference())
		return
	}

	if resp.StatusCode != 200 {
		log.Printf("API error %d: %s", resp.StatusCode, string(body))
		messenger.Send(s, m.ChannelID, "Error communicating with AI", m.Reference())
		return
	}

	var response struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Printf("Error parsing response: %v", err)
		messenger.Send(s, m.ChannelID, "Error communicating with AI", m.Reference())
		return
	}

	if len(response.Content) > 0 {
		messenger.Send(s, m.ChannelID, response.Content[0].Text, m.Reference())
	}
}
