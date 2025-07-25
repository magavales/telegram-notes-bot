package models

import "time"

type Note struct {
	text string
	t    *MyTime
}

func NewNote() *Note {
	return &Note{t: NewMyTime()}
}

func (n *Note) GetText() string {
	return n.text
}

func (n *Note) SetText(text string) {
	n.text = text
}

func (n *Note) GetTime() time.Time {
	return n.t.Get()
}

func (n *Note) SetTime(t string) error {
	err := n.t.Set(t)

	return err
}
