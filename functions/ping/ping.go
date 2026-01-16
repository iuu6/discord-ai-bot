package ping

import (
	"discord-ai-bot/utils/messenger"

	"github.com/bwmarrin/discordgo"
)

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	messenger.Send(s, m.ChannelID, "pang", m.Reference())
}
