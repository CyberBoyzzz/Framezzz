package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/CyberBoyzzz/Framezzz/internal/model"
	"github.com/CyberBoyzzz/Framezzz/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// GetComics godoc
//
//	@Summary		Get all books
//	@Tags			Books
//	@Produce		json
//	@Success		200	{object} []model.GetBookResponse
//
// @Router			/api/books [get]
func (h *Handlers) GetComicsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	comic, err := h.Storage.GetComics(ctx)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	comicResponse := []model.GetComicResponse{}
	for _, comic := range comic {
		comicResponse = append(comicResponse, model.GetComicResponse{
			ID:       comic.ID,
			Title:    comic.Title,
			CoverURL: comic.CoverURL,
		})

	}

	err = h.Sender.JSON(w, http.StatusOK, comicResponse)
	if err != nil {
		logger.OutputLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatal("Error when requesting /books")

		panic(err)
	}
}

// GetComic godoc
//
//	@Summary		Get a specific book
//	@Tags			Books
//	@Produce		json
//	@Param			id path int	true "Book ID"
//	@Success		200	{object} model.GetBookResponse
//
// @Router			/api/book/{id} [get]
func (h *Handlers) GetComicHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	var comic model.Comic
	comic, err = h.Storage.GetComic(ctx, id)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if (model.Comic{}) == comic {
		h.Sender.JSON(w, http.StatusBadRequest, "Book with id="+fmt.Sprint(comic.ID)+" not found")
		if err != nil {
			panic(err)
		}
		return
	}

	bookResponse := model.GetComicResponse{
		ID:       comic.ID,
		Title:    comic.Title,
		CoverURL: comic.CoverURL,
		Likes:    comic.Likes,
	}

	err = h.Sender.JSON(w, http.StatusOK, bookResponse)
	if err != nil {
		logger.OutputLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatal(fmt.Sprint("Error when requesting /comic/", comic.ID))

		panic(err)
	}
}

// UpdateComic godoc
//
//		@Summary		Update a specific book
//		@Tags			Books
//		@Produce		json
//	 	@Accept			json
//		@Param			title body string false "Book title"
//		@Param			author body string false "Book author"
//		@Param			coverUrl body string false "Book coverUrl"
//		@Param			postUrl body string false "Book post url"
//		@Success		200	{object} model.IDResponse
//
// @Router			/api/book/update [patch]
func (h *Handlers) UpdateComicHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var comic model.UpdateComicRequest

	err := json.NewDecoder(r.Body).Decode(&comic)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = Validate.Struct(comic)
	if err != nil {
		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, err.Field()+" "+err.Tag())
		}
		h.Sender.JSON(w, http.StatusBadRequest, strings.Join(errs, ", "))
		return
	}

	exists, err := h.Storage.VerifyComicExists(ctx, comic.ID)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		h.Sender.JSON(w, http.StatusBadRequest, "Book with id="+fmt.Sprint(comic.ID)+" not found")
		return
	}

	id, err := h.Storage.UpdateComic(ctx, comic)
	if err != nil {
		h.Sender.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := model.IDResponse{ID: id}
	err = h.Sender.JSON(w, http.StatusOK, response)
	if err != nil {
		panic(err)
	}
}

// cSpell:ignore godoc logrus
