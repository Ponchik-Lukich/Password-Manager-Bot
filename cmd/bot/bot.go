package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"password-manager/internal/handlers"
)

func RunBot(botToken string) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		handlers.HandleUpdate(bot, &update)
	}
}
