package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleDel(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendMessage(bot, update.Message.Chat.ID, "Del command")
}
