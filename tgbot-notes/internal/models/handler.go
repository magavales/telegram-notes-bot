package models

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgbot-notes/internal/models/quotes"
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
func (h *Handler) HandleDialog(ctx context.Context, telegramBot *services.TelegramService, message *tgbotapi.Message, note *Note, dialog *string) error {
	var err error
	switch *dialog {
	case "set_note":
		if note.GetText() == "" {
			note.SetText(message.Text)
			err = telegramBot.Send(message, quotes.SettingNoteQuoteSetTime)
			break
		}
		if note.GetTime() != time.Now() {
			err = note.SetTime(message.Text)
			if err != nil {
				err = telegramBot.Send(message, quotes.SettingNoteQuoteSetTimeError)
				break
			}
			err = telegramBot.Send(message, quotes.SettingNoteQuoteEnd)
			*dialog = ""
			break
		}
	}

	return err
}
