package postgres_repo

import (
	"context"
	"database/sql"
	"tgbot-notes/internal/models"
	"tgbot-notes/internal/models/buttons"
	"time"
)

type TableNotes struct {
	db *sql.DB
}

func NewTableNotes(db *sql.DB) *TableNotes {
	return &TableNotes{db: db}
}

func (tn *TableNotes) Create(ctx context.Context, note *models.Note) error {
	_, err := tn.db.ExecContext(ctx, "INSERT INTO notes (chat_id, note, date, status)  VALUES ($1, $2, $3, $4)", note.GetChatID(), note.GetNote(), note.GetDate(), note.GetStatus())
	if err != nil {
		return err
	}

	return nil
}

func (tn *TableNotes) Get(ctx context.Context, chatID int64) ([]*models.Note, error) {
	var (
		id     int64
		chat   int64
		text   string
		date   string
		status string
		notes  []*models.Note
		note   *models.Note
	)
	notes = make([]*models.Note, 0)
	rows, err := tn.db.QueryContext(ctx, "SELECT * FROM notes WHERE chat_id = $1", chatID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		note = models.NewNote()
		err := rows.Scan(&id, &chat, &text, &date, &status)
		if err != nil {
			return nil, err
		}

		note.SetID(id)
		note.SetChatID(chat)
		note.SetNote(text)
		note.SetDate(ctx, date)
		note.SetStatus(status)
		notes = append(notes, note)
	}

	return notes, nil
}

func (tn *TableNotes) GetByDate(ctx context.Context, chatID int64, button string) ([]*models.Note, error) {
	var (
		err     error
		id      int64
		chat    int64
		text    string
		date    string
		status  string
		notes   []*models.Note
		note    *models.Note
		newTime time.Time
		rows    *sql.Rows
	)
	switch button {
	case buttons.Tomorrow:
		rows, err = tn.db.QueryContext(ctx, "SELECT * FROM notes WHERE chat_id = $1", chatID)
		for rows.Next() {
			note = models.NewNote()
			err = rows.Scan(&id, &chat, &text, &date, &status)
			if err != nil {
				return nil, err
			}

			newTime, err = time.Parse(time.RFC3339, date)
			if err != nil {
				return nil, err
			} else {
				if newTime.Day() == time.Now().Day()+1 && newTime.Month() == time.Now().Month() && newTime.Year() == time.Now().Year() {
					note.SetID(id)
					note.SetChatID(chat)
					note.SetNote(text)
					note.SetDate(ctx, date)
					note.SetStatus(status)
					notes = append(notes, note)
				} else {
					continue
				}
			}
		}
	}

	return notes, err
}
