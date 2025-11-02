package interfaces

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Telegram interface {
	Send(message tgbotapi.MessageConfig) error
}
