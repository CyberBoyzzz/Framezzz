package storage

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/CyberBoyzzz/Framezzz/internal/model"
)

func (s *Storage) GetComic(ctx context.Context, id int) (model.Comic, error) {
	var book model.Comic

	err := s.db.Get(&book, `Select * from comics where id=$1`, id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *Storage) GetBooks(ctx context.Context) ([]model.Comic, error) {
	var books []model.Comic
	err := s.db.Select(&books, `SELECT * from comics`)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *Storage) UpdateBook(ctx context.Context, book model.UpdateComicRequest) (int, error) {
	var columns []string
	var argCount = 1
	var args []interface{}

	if book.Title != "" {
		columns = append(columns, fmt.Sprintf("title=$%d", argCount))
		args = append(args, book.Title)
		argCount++
	}

	if book.CoverURL != "" {
		columns = append(columns, fmt.Sprintf("cover_url=$%d", argCount))
		args = append(args, book.CoverURL)
		argCount++
	}

	columns = append(columns, fmt.Sprintf("updated_at=$%d", argCount))
	args = append(args, time.Now().UTC())
	argCount++

	if len(columns) == 0 {
		return 0, errors.New("No fields to update")
	}

	args = append(args, book.ID)

	query := fmt.Sprintf(`UPDATE comics SET %s WHERE id=$%d RETURNING id`, strings.Join(columns, ", "), argCount)

	var id int
	err := s.db.Get(&id, query, args...)
	if err != nil {
		return 0, err
	}
	return id, nil
}
