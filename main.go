package main

import (
	"discord-ai-bot/functions/chat"
	"discord-ai-bot/functions/ping"
	"discord-ai-bot/utils/config"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	cfg := config.Load()

	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	dg.AddHandler(messageHandler)

	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	err = dg.Open()
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}
	defer dg.Close()

	log.Println("Bot is running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "/ping") {
		ping.Handle(s, m)
	} else if strings.HasPrefix(m.Content, "/chat ") {
		chat.Handle(s, m)
	}
}
