package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"password-manager/cmd/database"
	"password-manager/cmd/models"
)

func handleStart(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	err := database.AddUser(models.User{ChatID: chatID})
	if err != nil {
		log.Fatal(err)
	}
}
