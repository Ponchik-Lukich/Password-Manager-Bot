package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"password-manager/cmd/database"
)

func handleDel(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendMessage(bot, update.CallbackQuery.Message.Chat.ID, "Enter service name to delete:")
	err := database.SetUserState(update.CallbackQuery.Message.Chat.ID, "del")
	if err != nil {
		log.Print(err)
	}
}

func handleDelService(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	err := database.DeleteService(update.Message.Text, update.Message.Chat.ID)
	if err != nil {
		sendMessage(bot, update.Message.Chat.ID, "Error deleting service ("+err.Error()+")")
		err := database.SetUserState(update.Message.Chat.ID, "wait")
		if err != nil {
			log.Print(err)
		}
		handleUnknownCommand(bot, update)
		return
	} else {
		sendMessage(bot, update.Message.Chat.ID, "Service deleted successfully")
		err := database.SetUserState(update.Message.Chat.ID, "wait")
		if err != nil {
			log.Print(err)
		}
		handleUnknownCommand(bot, update)
		return
	}
}
