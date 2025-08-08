package repository

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"tgbot-notes/internal/repository/postgres_repo"
)

type Postgres struct {
	db         *sql.DB
	TableNotes *postgres_repo.TableNotes
}

func NewPostgres(url string) (*Postgres, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return &Postgres{db: nil, TableNotes: postgres_repo.NewTableNotes(db)}, err
	}

	return &Postgres{db: db, TableNotes: postgres_repo.NewTableNotes(db)}, nil
}
