package postgres_repo

import (
	"context"
	"database/sql"
	"tgbot-notes/internal/models"
)

type TableNotes struct {
	db *sql.DB
}

func NewTableNotes(db *sql.DB) *TableNotes {
	return &TableNotes{db: db}
}

func (tn *TableNotes) Create(ctx context.Context, note *models.Note) error {
	_, err := tn.db.ExecContext(ctx, "INSERT INTO notes (chat_id, note, date, status) VALUES ($1, $2, $3, $4)", note.GetChatID(), note.GetNote(), note.GetDate(), note.GetStatus())
	if err != nil {
		return err
	}

	return nil
}

func (tn *TableNotes) Get() {

}
