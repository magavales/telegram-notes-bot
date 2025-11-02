package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgbot-notes/internal/models/buttons"
)

type ReplyKeyboard struct {
	keyboard tgbotapi.InlineKeyboardMarkup
}

func NewReplyKeyboard() *ReplyKeyboard {
	return &ReplyKeyboard{
		keyboard: tgbotapi.NewInlineKeyboardMarkup(),
	}
}

func (rk *ReplyKeyboard) GetKeyboard() *tgbotapi.InlineKeyboardMarkup {
	return &rk.keyboard
}

func (rk *ReplyKeyboard) CreateKeyboardGetNoteByDate() {
	tomorrow := tgbotapi.NewInlineKeyboardButtonData(buttons.Tomorrow, buttons.Tomorrow)
	thisWeek := tgbotapi.NewInlineKeyboardButtonData(buttons.ThisWeek, buttons.ThisWeek)
	nextWeek := tgbotapi.NewInlineKeyboardButtonData(buttons.NextWeek, buttons.NextWeek)
	thisMonth := tgbotapi.NewInlineKeyboardButtonData(buttons.ThisMonth, buttons.ThisMonth)
	nextMonth := tgbotapi.NewInlineKeyboardButtonData(buttons.NextMonth, buttons.NextMonth)
	rowOne := tgbotapi.NewInlineKeyboardRow(tomorrow, thisWeek, nextWeek)
	rowTwo := tgbotapi.NewInlineKeyboardRow(thisMonth, nextMonth)
	rk.keyboard.InlineKeyboard = append(rk.keyboard.InlineKeyboard, rowOne, rowTwo)
}
