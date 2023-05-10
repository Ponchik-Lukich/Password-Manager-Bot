package bot

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"os"
	"password-manager/cmd/handlers"
)

func RunBot(botToken string, w http.ResponseWriter, r *http.Request) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	webhookURL := os.Getenv("TELEGRAM_WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal("Webhook URL is required")
	}

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookURL))
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Printf("Error unmarshalling update: %v", err)
		return
	}

	handlers.HandleUpdate(bot, &update)
}

func RunBotLocal(botToken string) {
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
