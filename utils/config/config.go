package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken   string
	AnthropicKey   string
	AnthropicBase  string
	AnthropicModel string
	SystemPrompt   string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DiscordToken:   os.Getenv("DISCORD_TOKEN"),
		AnthropicKey:   os.Getenv("ANTHROPIC_API_KEY"),
		AnthropicBase:  os.Getenv("ANTHROPIC_API_BASE"),
		AnthropicModel: os.Getenv("ANTHROPIC_MODEL"),
		SystemPrompt:   os.Getenv("SYSTEM_PROMPT"),
	}
}
