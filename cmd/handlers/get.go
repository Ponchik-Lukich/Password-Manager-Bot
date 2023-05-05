package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func handleGet(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendMessage(bot, update.Message.Chat.ID, "Get command")
}
