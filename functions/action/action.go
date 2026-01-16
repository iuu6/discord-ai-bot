package action

import (
	"discord-ai-bot/utils/messenger"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.MessageReference == nil {
		return
	}

	content := m.Content
	var action string
	var reverse bool

	if strings.HasPrefix(content, "//") {
		action = strings.TrimPrefix(content, "//")
		reverse = false
	} else if strings.HasPrefix(content, "/") {
		action = strings.TrimPrefix(content, "/")
		reverse = false
	} else if strings.HasPrefix(content, "\\\\") {
		action = strings.TrimPrefix(content, "\\\\")
		reverse = true
	} else if strings.HasPrefix(content, "\\") {
		action = strings.TrimPrefix(content, "\\")
		reverse = true
	} else {
		return
	}

	action = strings.TrimSpace(action)
	if action == "" {
		return
	}

	refMsg, err := s.ChannelMessage(m.ChannelID, m.MessageReference.MessageID)
	if err != nil {
		return
	}

	var response string
	if reverse {
		response = m.Author.Mention() + " 被 " + refMsg.Author.Mention() + " " + action + "了 !"
	} else {
		response = m.Author.Mention() + " " + action + "了 " + refMsg.Author.Mention() + " !"
	}

	messenger.Send(s, m.ChannelID, response, m.Reference())
}
