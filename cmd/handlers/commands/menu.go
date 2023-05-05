package commands

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleMenu(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Get", "get"),
			tgbotapi.NewInlineKeyboardButtonData("Set", "set"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Del", "del"),
			tgbotapi.NewInlineKeyboardButtonData("Start", "start"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Choose a command:")
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
}
