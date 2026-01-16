package messenger

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Send(s *discordgo.Session, channelID string, content string) {
	_, err := s.ChannelMessageSend(channelID, content)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
