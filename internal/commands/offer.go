// commands/offer.go
package commands

import (
	"github.com/BisquitDubouche/SberMarket_Saver_Bot/internal"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleOffer(ctx internal.BotContext, chatID int64) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(chatID, "Вы выбрали '🔥 Предложение дня'.")
	return msg
}
