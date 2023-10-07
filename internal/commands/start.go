package commands

import (
	"github.com/BisquitDubouche/SberMarket_Saver_Bot/internal"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleStart(ctx internal.BotContext, chatID int64) tgbotapi.Chattable {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔍 Найти товары"),
			tgbotapi.NewKeyboardButton("🔥 Предложение дня"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Привет, я бот для поиска выгодных предложений на СберМаркете. Снизу должны быть кнопочки для выбора действия")
	msg.ReplyMarkup = keyboard

	return msg
}
