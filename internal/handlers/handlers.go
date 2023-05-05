package handlers

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, I'm your bot!")
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
