// context.go
package internal

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
    "time"
)

type BotContext struct {
    Bot            *tgbotapi.BotAPI
    LastSearchTime map[int64]time.Time
}
