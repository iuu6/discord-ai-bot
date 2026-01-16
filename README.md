# Discord AI Bot

A Discord bot with AI chat capabilities using OpenAI API.

## Setup

1. Copy `.env.example` to `.env`:
```bash
cp .env.example .env
```

2. Edit `.env` and fill in your credentials:
   - `DISCORD_TOKEN`: Your Discord bot token
   - `OPENAI_API_KEY`: Your OpenAI API key
   - `OPENAI_API_BASE`: OpenAI API base URL (default: https://api.openai.com/v1)
   - `OPENAI_MODEL`: Model to use (default: gpt-3.5-turbo)
   - `SYSTEM_PROMPT`: Optional system prompt for the AI

3. Install dependencies:
```bash
go mod download
```

4. Run the bot:
```bash
go run main.go
```

## Commands

- `/ping` - Bot replies with "pang"
- `/chat <message>` - Chat with AI

## Project Structure

- `/functions/` - Command handlers (each subfolder is a feature)
- `/utils/` - Utilities and middleware
  - `/utils/config/` - Configuration loader
  - `/utils/messenger/` - Message sending utility
