package services

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgbot-notes/internal/interfaces"
	"tgbot-notes/internal/models"
	"tgbot-notes/internal/repository"
)

type TelegramService struct {
	telegram interfaces.Telegram
	repo     *repository.Postgres
}

func NewTelegramService(telegram interfaces.Telegram, repo *repository.Postgres) *TelegramService {
	t := &TelegramService{
		telegram: telegram,
		repo:     repo,
	}

	return t
}

func (service *TelegramService) Send(message *tgbotapi.Message, text string) error {
	err := service.telegram.Send(message.Chat.ID, text)
	return err
}

func (service *TelegramService) SentNote(ctx context.Context, note *models.Note) error {
	err := service.repo.TableNotes.Create(ctx, note)
	return err
}
