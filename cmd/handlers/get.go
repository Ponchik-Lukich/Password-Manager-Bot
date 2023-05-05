package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"password-manager/cmd/database"
)

func handleGet(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendMessage(bot, update.CallbackQuery.Message.Chat.ID, "Enter service name:")
	err := database.SetUserState(update.CallbackQuery.Message.Chat.ID, "get")
	if err != nil {
		log.Fatal(err)
	}
}

func handleGetServiceName(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	serviceName := update.Message.Text
	err := database.SetUserState(update.Message.Chat.ID, "wait")
	if err == nil {
		log.Fatal(err)
	}
	service, err := database.GetService(serviceName)
	if err != nil {
		if err.Error() == "not found" {
			sendMessage(bot, update.Message.Chat.ID, "Service not found")
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		sendMessage(bot, update.Message.Chat.ID, "Error retrieving service ("+err.Error()+")")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	response := fmt.Sprintf("Service: %s\nLogin: %s\nPassword: %s", service.Name, service.Login, service.Password)
	sendMessage(bot, update.Message.Chat.ID, response)
}
