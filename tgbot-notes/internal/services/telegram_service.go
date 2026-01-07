package services

import (
	"context"
	"tgbot-notes/internal/interfaces"
	"tgbot-notes/internal/models"
	"tgbot-notes/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func (service *TelegramService) Send(message tgbotapi.MessageConfig) error {
	err := service.telegram.Send(message)
	return err
}

func (service *TelegramService) SetNote(ctx context.Context, note *models.Note) (int64, error) {
	id, err := service.repo.TableNotes.Create(ctx, note)
	return id, err
}

func (service *TelegramService) GetNotes(ctx context.Context, chatID int64) ([]*models.Note, error) {
	notes, err := service.repo.TableNotes.Get(ctx, chatID)
	return notes, err
}

func (service *TelegramService) GetNotesByDate(ctx context.Context, chatID int64, button string) ([]*models.Note, error) {
	notes, err := service.repo.TableNotes.GetByDate(ctx, chatID, button)
	return notes, err
}
