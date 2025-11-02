package models

import (
	"context"
	"fmt"
	"time"
)

type Note struct {
	id     int64
	chatID int64
	note   string
	date   *MyTime
	status string
}

func NewNote() *Note {
	return &Note{date: NewMyTime()}
}

func (n *Note) GetID() int64 {
	return n.id
}

func (n *Note) SetID(id int64) {
	n.id = id
}

func (n *Note) GetChatID() int64 {
	return n.chatID
}

func (n *Note) SetChatID(chatID int64) {
	n.chatID = chatID
}

func (n *Note) GetNote() string {
	return n.note
}

func (n *Note) SetNote(text string) {
	n.note = text
}

func (n *Note) GetDate() time.Time {
	return n.date.Get()
}

func (n *Note) SetDate(ctx context.Context, t string) error {
	err := n.date.Set(ctx, t)

	return err
}

func (n *Note) GetStatus() string {
	return n.status
}

func (n *Note) SetStatus(status string) {
	n.status = status
}

func (n *Note) String() string {
	msg := fmt.Sprintf("Задача: %s\nКогда напомнить: %s\nСтатус задачи: %s\n\n", n.note, n.date.String(), n.status)
	return msg
}
