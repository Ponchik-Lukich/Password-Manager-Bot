package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"password-manager/cmd/database"
	"password-manager/cmd/models"
)

func HandleStart(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	err := database.AddUser(models.User{ChatID: chatID})
	if err != nil {
		log.Fatal(err)
	}

	HandleMenu(bot, update)
}
