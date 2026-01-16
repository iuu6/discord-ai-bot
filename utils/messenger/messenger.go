package messenger

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Send(s *discordgo.Session, channelID string, content string, reference *discordgo.MessageReference) {
	_, err := s.ChannelMessageSendReply(channelID, content, reference)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
