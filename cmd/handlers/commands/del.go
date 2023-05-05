package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"password-manager/cmd/handlers"
)

func HandleDel(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	handlers.SendMessage(bot, update.Message.Chat.ID, "Del command", update.Message.MessageID)
}
