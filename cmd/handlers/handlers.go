package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"password-manager/cmd/database"
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
		if update.Message.IsCommand() && update.Message.Command() == "start" {
			handleStart(update)
		}
		user, err := database.GetUser(update.Message.Chat.ID)
		if err != nil {
			log.Fatal(err)
		}
		switch user.State {
		case "get":
			handleGetService(bot, update)
		case "set":
			handleSetService(bot, update)
		default:
			handleUnknownCommand(bot, update)
		}
	} else if update.CallbackQuery != nil {
		handleCallbackQuery(bot, update)
	}
}

func handleUnknownCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Here is my command list:")
	msg.ReplyMarkup = inlineKeyboard
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string, replyToMessageIDs ...int) {
	msg := tgbotapi.NewMessage(chatID, text)
	for _, replyToMessageID := range replyToMessageIDs {
		msg.ReplyToMessageID = replyToMessageID
	}
	bot.Send(msg)
}

func handleCallbackQuery(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	switch update.CallbackQuery.Data {
	case "get":
		handleGet(bot, update)
	case "set":
		handleSet(bot, update)
	case "del":
		handleDel(bot, update)
	default:
		handleUnknownCommand(bot, update)
	}
}
