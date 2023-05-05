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

	//err := database.Connect()
	//if err != nil {
	//	panic("Could not connect to database")
	//}

	bot.RunBot(botToken, w, r)
}
