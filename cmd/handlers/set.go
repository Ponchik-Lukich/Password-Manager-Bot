package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"password-manager/cmd/database"
	"password-manager/cmd/utils"
)

func handleSet(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendMessage(bot, update.CallbackQuery.Message.Chat.ID, "Enter service credentials in format:\n<service name>:<login>:<password>")
	err := database.SetUserState(update.CallbackQuery.Message.Chat.ID, "set")
	if err != nil {
		log.Fatal(err)
	}
}

func handleSetService(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	serviceCredentials := update.Message.Text
	service, err := utils.ValidateService(serviceCredentials, update.Message.Chat.ID)
	sendMessage(bot, update.Message.Chat.ID, "Validating service..."+service.Name+" "+service.Login+" "+service.Password)
	if err != nil {
		sendMessage(bot, update.Message.Chat.ID, err.Error())
		err := database.SetUserState(update.Message.Chat.ID, "wait")
		if err != nil {
			log.Fatal(err)
		}
		return
	} else {
		err = database.AddService(service)
		if err != nil {
			sendMessage(bot, update.Message.Chat.ID, "Error adding service ("+err.Error()+")")
			err := database.SetUserState(update.Message.Chat.ID, "wait")
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		sendMessage(bot, update.Message.Chat.ID, "Service added/updated successfully")
		err := database.SetUserState(update.Message.Chat.ID, "wait")
		if err != nil {
			log.Fatal(err)
		}
		return
	}
}
