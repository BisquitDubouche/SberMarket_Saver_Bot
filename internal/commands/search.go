// commands/search.go
package commands

import (
	"time"

	"github.com/BisquitDubouche/SberMarket_Saver_Bot/internal"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleSearch(ctx internal.BotContext, chatID int64) tgbotapi.Chattable {
	lastTime, ok := ctx.LastSearchTime[chatID]
	if ok && time.Since(lastTime) < 1*time.Minute {
		msg := tgbotapi.NewMessage(chatID, "Пожалуйста, подожди 1 минуту или воспользуйся навигационной панелью отправленной ранее.")
		return msg
	}
	ctx.LastSearchTime[chatID] = time.Now()

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🖥 Электроника", "electronics"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🧥 Одежда", "dress"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Пожалуйста, выбери категорию:")
	msg.ReplyMarkup = inlineKeyboard

	return msg
}
