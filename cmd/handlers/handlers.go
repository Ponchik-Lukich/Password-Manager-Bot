package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"password-manager/cmd/database"
	"password-manager/cmd/models"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		if update.Message.Command() == "start" {
			database.AddUser(models.User{ChatID: update.Message.Chat.ID})
		}
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		_, err := bot.Send(msg)
		if err != nil {
			return
		}
	}

}
