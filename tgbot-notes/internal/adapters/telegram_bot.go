package adapters

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)

	return &TelegramBot{bot: bot}, err
}

func (bot *TelegramBot) GetUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.bot.GetUpdatesChan(u)

	return updates
}

func (bot *TelegramBot) Send(chatID int64, text string) error {
	message := tgbotapi.NewMessage(chatID, text)
	_, err := bot.bot.Send(message)
	return err
}
