package handler

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgbot-notes/internal/models"
	"tgbot-notes/internal/models/buttons"
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
		newMessage := tgbotapi.NewMessage(message.Chat.ID, text)
		err = telegramBot.Send(newMessage)
	case "set_note":
		newMessage := tgbotapi.NewMessage(message.Chat.ID, quotes.SettingNoteQuoteSetText)
		err = telegramBot.Send(newMessage)
	case "get_notes":
		notes, err := telegramBot.GetNotes(ctx, message.Chat.ID)
		if err != nil {
			newMessage := tgbotapi.NewMessage(message.Chat.ID, quotes.GettingNotesEmpty)
			err = telegramBot.Send(newMessage)
		}
		msg := ""
		for i, value := range notes {
			msg = msg + fmt.Sprintf("Задача номер %d\n", i+1) + value.String()
		}
		newMessage := tgbotapi.NewMessage(message.Chat.ID, msg)
		err = telegramBot.Send(newMessage)
	case "get_note_by_date":
		replyKeyboard := models.NewReplyKeyboard()
		replyKeyboard.CreateKeyboardGetNoteByDate()
		newMessage := tgbotapi.NewMessage(message.Chat.ID, "Выбери свой путь: ")
		newMessage.ReplyMarkup = replyKeyboard.GetKeyboard()
		err = telegramBot.Send(newMessage)
	case "help":
		newMessage := tgbotapi.NewMessage(message.Chat.ID, quotes.HelpQuote)
		err = telegramBot.Send(newMessage)
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
			newMessage := tgbotapi.NewMessage(message.Chat.ID, quotes.SettingNoteQuoteSetTime)
			err = telegramBot.Send(newMessage)
			return err
		}
		if note.GetDate() != time.Now() {
			ctx = context.WithValue(ctx, "set_note", "set_note")
			err = note.SetDate(ctx, message.Text)
			if err != nil {
				newMessage := tgbotapi.NewMessage(message.Chat.ID, quotes.SettingNoteQuoteSetTimeError)
				err = telegramBot.Send(newMessage)
				return err
			}
			note.SetChatID(message.Chat.ID)
			note.SetStatus(statuses.Uncompleted)
			err = telegramBot.SetNote(ctx, note)
			if err != nil {
				newMessage := tgbotapi.NewMessage(message.Chat.ID, quotes.SettingNoteQuoteSetError)
				err = telegramBot.Send(newMessage)
				*dialog = ""
				return err
			}
			newMessage := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(quotes.SettingNoteQuoteEnd+note.GetDate().Format("02.01.2006 15:04")))
			err = telegramBot.Send(newMessage)
			*dialog = ""
			ctx = context.WithValue(ctx, "set_note", "")
			return err
		}

	}

	return err
}

func (h *Handler) HandleCallback(ctx context.Context, telegramBot *services.TelegramService, callback *tgbotapi.CallbackQuery) error {
	switch callback.Data {
	case buttons.Tomorrow:
		notes, err := telegramBot.GetNotesByDate(ctx, callback.Message.Chat.ID, buttons.Tomorrow)
		if err != nil {
			newMessage := tgbotapi.NewMessage(callback.Message.Chat.ID, quotes.GettingNotesEmpty)
			err = telegramBot.Send(newMessage)
		}
		msg := ""
		for i, value := range notes {
			msg = msg + fmt.Sprintf("Задача номер %d\n", i+1) + value.String()
		}
		newMessage := tgbotapi.NewMessage(callback.Message.Chat.ID, msg)
		err = telegramBot.Send(newMessage)

		return err
	}

	return nil
}
