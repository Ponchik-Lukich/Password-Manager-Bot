package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "get":
		handleGet(bot, update)
	case "set":
		handleSet(bot, update)
	case "del":
		handleDel(bot, update)
	case "menu":
		handleMenu(bot, update)
	case "start":
		handleStart(bot, update)
	default:
		handleUnknownCommand(bot, update)
	}
}

func handleUnknownCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendMessage(bot, update.Message.Chat.ID, "I don't understand this command.", update.Message.MessageID)
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string, replyToMessageID int) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = replyToMessageID
	bot.Send(msg)
}
