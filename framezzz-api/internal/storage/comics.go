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
	var comic model.Comic

	err := s.db.Get(&comic, `Select * from comics where id=$1`, id)
	if err != nil {
		return comic, err
	}

	return comic, nil
}

func (s *Storage) GetComics(ctx context.Context) ([]model.Comic, error) {
	var comic []model.Comic
	err := s.db.Select(&comic, `SELECT * from comics`)
	if err != nil {
		return nil, err
	}

	return comic, nil
}

func (s *Storage) UpdateComic(ctx context.Context, comic model.UpdateComicRequest) (int, error) {
	var columns []string
	var argCount = 1
	var args []interface{}

	if comic.Title != "" {
		columns = append(columns, fmt.Sprintf("title=$%d", argCount))
		args = append(args, comic.Title)
		argCount++
	}

	if comic.CoverURL != "" {
		columns = append(columns, fmt.Sprintf("cover_url=$%d", argCount))
		args = append(args, comic.CoverURL)
		argCount++
	}

	if comic.Likes > 0 {
		columns = append(columns, fmt.Sprintf("likes = likes + $%d", argCount))
		args = append(args, comic.Likes)
		argCount++
	}

	// Always update the updated_at timestamp
	columns = append(columns, fmt.Sprintf("updated_at=$%d", argCount))
	args = append(args, time.Now().UTC())
	argCount++

	if len(columns) == 0 {
		return 0, errors.New("No fields to update")
	}

	args = append(args, comic.ID)

	query := fmt.Sprintf(`UPDATE comics SET %s WHERE id=$%d RETURNING id`, strings.Join(columns, ", "), argCount)

	var id int
	err := s.db.Get(&id, query, args...)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) VerifyComicExists(ctx context.Context, id int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM comics WHERE id = $1)"
	err := s.db.GetContext(ctx, &exists, query, id)
	if err != nil {
		return false, err
	}
	return exists, nil
}
