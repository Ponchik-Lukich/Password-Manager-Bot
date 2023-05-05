package handler

import (
	"net/http"
	"os"
	"password-manager/cmd/bot"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		panic("Bot token is required")
	}

	bot.RunBot(botToken)
}
