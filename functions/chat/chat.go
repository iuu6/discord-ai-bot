package chat

import (
	"context"
	"discord-ai-bot/utils/config"
	"discord-ai-bot/utils/messenger"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	cfg := config.Load()

	query := strings.TrimPrefix(m.Content, "/chat ")
	if query == "" {
		messenger.Send(s, m.ChannelID, "Please provide a message after /chat")
		return
	}

	client := openai.NewClient(
		option.WithAPIKey(cfg.OpenAIKey),
		option.WithBaseURL(cfg.OpenAIBase),
	)

	messages := []openai.ChatCompletionMessageParamUnion{}

	if cfg.SystemPrompt != "" {
		messages = append(messages, openai.SystemMessage(cfg.SystemPrompt))
	}

	messages = append(messages, openai.UserMessage(query))

	completion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Model:    openai.F(cfg.OpenAIModel),
		Messages: openai.F(messages),
	})

	if err != nil {
		log.Printf("Error calling OpenAI: %v", err)
		messenger.Send(s, m.ChannelID, "Error communicating with AI")
		return
	}

	if len(completion.Choices) > 0 {
		messenger.Send(s, m.ChannelID, completion.Choices[0].Message.Content)
	}
}
