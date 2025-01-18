package model

type ExternalComicAPIResponse struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// Comic is the db schema for the books table
type Comic struct {
	ID       int    `json:"num" db:"id"`
	Title    string `json:"safe_title" db:"safe_title"`
	CoverURL string `json:"img_url" db:"img_url"`
	Likes    int    `json:"likes" db:"likes"`
	Base
}

// Below are the structures of the request/response structs in the comics handler.
type GetComicResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	CoverURL string `json:"img"`
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
