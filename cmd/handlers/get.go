package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"password-manager/cmd/database"
	"time"
)

func handleGet(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendMessage(bot, update.CallbackQuery.Message.Chat.ID, "Enter service name:")
	err := database.SetUserState(update.CallbackQuery.Message.Chat.ID, "get")
	if err != nil {
		log.Print(err)
	}
}

func handleWaitDelete(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	sendMessage(bot, update.CallbackQuery.Message.Chat.ID, "Credentials were hidden")
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
			log.Print(err)
		}
		handleUnknownCommand(bot, update)
		return
	} else {
		response := fmt.Sprintf("Your credentials:\nService: %s\nLogin: %s\nPassword: %s\nType any word to delete "+
			"this message", service.Name, service.Login, service.Password)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		sentMessage, _ := bot.Send(msg)
		err = database.SetUserState(update.Message.Chat.ID, "wait_delete")
		if err == nil {
			log.Print(err)
		}
		time.AfterFunc(time.Second*1, func() {
			deleteMessage(bot, update.Message.Chat.ID, sentMessage.MessageID)
			err := database.SetUserState(update.CallbackQuery.Message.Chat.ID, "wait")
			if err != nil {
				log.Print(err)
			}
			handleUnknownCommand(bot, update)
		})
		return
	}
}
