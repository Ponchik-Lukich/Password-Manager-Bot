package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"password-manager/cmd/handlers"
)

func HandleSet(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	handlers.SendMessage(bot, update.Message.Chat.ID, "Set command", update.Message.MessageID)
}
