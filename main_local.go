package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"password-manager/cmd/bot"
	"password-manager/cmd/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		panic("Bot token is required")
	}

	err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = database.LocalInit()
	if err != nil {
		log.Fatalf("Failed to initialize tables: %v", err)
	}

	bot.RunBotLocal(botToken)
}
