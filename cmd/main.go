package main

import (
	"log"
	"os"
	"time"

	"github.com/BisquitDubouche/SberMarket_Saver_Bot/internal"
	"github.com/BisquitDubouche/SberMarket_Saver_Bot/internal/commands"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TG_API_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Авторизован как %s", bot.Self.UserName)

	updates, err := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))
	if err != nil {
		log.Fatal(err)
	}

	ctx := internal.BotContext{
		Bot:            bot,
		LastSearchTime: make(map[int64]time.Time),
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		switch update.Message.Text {
		case "/start":
			msg := commands.HandleStart(ctx, chatID)
			bot.Send(msg)

		case "🔍 Найти товары":
			msg := commands.HandleSearch(ctx, chatID)
			bot.Send(msg)

		case "🔥 Предложение дня":
			msg := commands.HandleOffer(ctx, chatID)
			bot.Send(msg)

		default:
			msg := tgbotapi.NewMessage(chatID, "Неизвестная команда.")
			commands.HandleSearch(ctx, chatID)
			bot.Send(msg)
		}
	}
}
