package services

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgbot-notes/internal/interfaces"
)

type TelegramService struct {
	telegram interfaces.Telegram
}

func NewTelegramService(telegram interfaces.Telegram) *TelegramService {
	return &TelegramService{telegram: telegram}
}

func (service *TelegramService) Send(message *tgbotapi.Message, text string) error {
	err := service.telegram.Send(message.Chat.ID, text)
	return err
}
