package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Get", "get"),
		tgbotapi.NewInlineKeyboardButtonData("Set", "set"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Del", "del"),
	),
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyMarkup = inlineKeyboard
		//if update.Message.IsCommand() && update.Message.Command() == "start" {
		//	handleStart(bot, update)
		//}
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	} else if update.CallbackQuery != nil {
		handleCallbackQuery(bot, update)
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

func handleCallbackQuery(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	//callbackQuery := update.CallbackQuery
	//switch callbackQuery.Data {
	//case "get":
	//	handleGet(bot, update)
	//case "set":
	//	handleSet(bot, update)
	//case "del":
	//	handleDel(bot, update)
	//default:
	//	handleUnknownCommand(bot, update)
	//}
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data+" command")
	if _, err := bot.AnswerCallbackQuery(callback); err != nil {
		panic(err)
	}

}
