package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"password-manager/cmd/commands"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "get":
		commands.HandleGet(bot, update)
	case "set":
		commands.HandleSet(bot, update)
	case "del":
		commands.HandleDel(bot, update)
	case "menu":
		commands.HandleMenu(bot, update)
	case "start":
		commands.HandleStart(bot, update)
	default:
		handleUnknownCommand(bot, update)
	}
}

func handleUnknownCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	SendMessage(bot, update.Message.Chat.ID, "I don't understand this command.", update.Message.MessageID)
}

func SendMessage(bot *tgbotapi.BotAPI, chatID int64, text string, replyToMessageID int) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = replyToMessageID
	bot.Send(msg)
}
