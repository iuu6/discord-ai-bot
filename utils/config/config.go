package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken   string
	OpenAIKey      string
	OpenAIBase     string
	OpenAIModel    string
	SystemPrompt   string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DiscordToken:   os.Getenv("DISCORD_TOKEN"),
		OpenAIKey:      os.Getenv("OPENAI_API_KEY"),
		OpenAIBase:     os.Getenv("OPENAI_API_BASE"),
		OpenAIModel:    os.Getenv("OPENAI_MODEL"),
		SystemPrompt:   os.Getenv("SYSTEM_PROMPT"),
	}
}
