package handler

import (
	"log"
	"net/http"
	"os"
	"password-manager/cmd/bot"
	"password-manager/internal/database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		panic("Bot token is required")
	}

	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	bot.RunBot(botToken, w, r)
}
