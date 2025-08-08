package handler

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgbot-notes/internal/models"
	"tgbot-notes/internal/models/quotes"
	"tgbot-notes/internal/models/statuses"
	"tgbot-notes/internal/services"
	"time"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// Обработка команд, поступающих от пользователя
func (h *Handler) HandleCommands(ctx context.Context, telegramBot *services.TelegramService, message *tgbotapi.Message) error {
	var err error
	switch message.Command() {
	case "start":
		helloMsg := fmt.Sprintf("Привет %s!", message.Chat.UserName)
		text := helloMsg + quotes.HelloQuote
		err = telegramBot.Send(message, text)
	case "set_note":
		err = telegramBot.Send(message, quotes.SettingNoteQuoteSetText)
	case "help":
		err = telegramBot.Send(message, quotes.HelpQuote)
	}

	return err
}

// Обработка состояния диалога для различных команд
func (h *Handler) HandleDialog(ctx context.Context, telegramBot *services.TelegramService, message *tgbotapi.Message, note *models.Note, dialog *string) error {
	var err error
	switch *dialog {
	case "set_note":
		if note.GetNote() == "" {
			note.SetNote(message.Text)
			err = telegramBot.Send(message, quotes.SettingNoteQuoteSetTime)
			return err
		}
		if note.GetDate() != time.Now() {
			err = note.SetDate(message.Text)
			if err != nil {
				err = telegramBot.Send(message, quotes.SettingNoteQuoteSetTimeError)
				return err
			}
			note.SetChatID(message.Chat.ID)
			note.SetStatus(statuses.Uncompleted)
			err = telegramBot.SentNote(ctx, note)
			if err != nil {
				err = telegramBot.Send(message, quotes.SettingNoteQuoteSetError)
				*dialog = ""
				return err
			}
			err = telegramBot.Send(message, quotes.SettingNoteQuoteEnd)
			*dialog = ""
			return err
		}
	}

	return err
}
