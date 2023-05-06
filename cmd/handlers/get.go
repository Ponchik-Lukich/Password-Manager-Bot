package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"password-manager/cmd/database"
	"password-manager/cmd/utils"
)

func handleGet(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendMessage(bot, update.CallbackQuery.Message.Chat.ID, "Enter service name:")
	err := database.SetUserState(update.CallbackQuery.Message.Chat.ID, "get")
	if err != nil {
		log.Print("user state change error: ", err)
	}
}

func handleWaitDelete(bot *tgbotapi.BotAPI, update *tgbotapi.Update, messageID int) {
	deleteMessage(bot, update.Message.Chat.ID, messageID)
	err := database.SetUserState(update.Message.Chat.ID, "wait")
	if err != nil {
		log.Print("user state change error: ", err)
	}
	sendMessage(bot, update.Message.Chat.ID, "Credentials were hidden")
	handleUnknownCommand(bot, update)
}

func handleGetService(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	service, err := database.GetService(update.Message.Text, update.Message.Chat.ID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			sendMessage(bot, update.Message.Chat.ID, "Service not found")
		} else {
			sendMessage(bot, update.Message.Chat.ID, "Error retrieving service ("+err.Error()+")")
		}
		err = database.SetUserState(update.Message.Chat.ID, "wait")
		if err != nil {
			log.Print("user state change error: ", err)
		}
		handleUnknownCommand(bot, update)
		return
	} else {
		login, err := utils.Decrypt(service.Login)
		if err != nil {
			log.Print("decrypt error: ", err)
		}
		password, err := utils.Decrypt(service.Password)
		if err != nil {
			log.Print("decrypt error: ", err)
		}
		response := fmt.Sprintf("Your credentials:\nService: %s\nLogin: %s\nPassword: %s\nType any word to delete "+
			"this message", service.Name, login, password)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		sentMessage, _ := bot.Send(msg)
		err = database.SetUserState(update.Message.Chat.ID, "wait_delete", sentMessage.MessageID)
		if err == nil {
			log.Print("user state change error: ", err)
		}
		return
	}
}
