package storage

import (
	"context"

	"github.com/CyberBoyzzz/Framezzz/internal/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initializes the postgres driver
)

type StorageInterface interface {
	GetComic(ctx context.Context, id int) (model.Comic, error)
	GetComics(ctx context.Context) ([]model.Comic, error)
	UpdateComic(ctx context.Context, comic model.UpdateComicRequest) (int, error)
	VerifyComicExists(ctx context.Context, id int) (bool, error)
	//UserLogin(ctx context.Context, id int) (model.Comic, error)
	//UserRegister(ctx context.Context, id int) (model.Comic, error)
}

// Storage contains an SQL db. Storage implements the StorageInterface.
type Storage struct {
	db *sqlx.DB
}

func (s *Storage) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetDB() *sqlx.DB {
	return s.db
}
