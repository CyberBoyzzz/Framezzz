package model

// Book is the db schema for the books table
type Comic struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	CoverURL string `json:"coverUrl" db:"cover_url"`
	Likes    int    `json:"likes" db:"likes"`
	Base
}

// Below are the structures of the request/response structs in the books handler.
type GetComicResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	CoverURL string `json:"coverUrl"`
	Likes    int    `json:"likes"`
}

type UpdateComicRequest struct {
	ID       int    `json:"id" validate:"required"`
	Title    string `json:"title"`
	CoverURL string `json:"coverUrl"`
	Likes    int    `json:"likes"`
}

type IDResponse struct {
	ID int `json:"id"`
}
