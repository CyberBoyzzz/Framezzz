package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/CyberBoyzzz/Framezzz/client"
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
//	@Summary		Get all comics
//	@Tags			Comics
//	@Produce		json
//	@Success		200	{object} []model.GetComicResponse
//
// @Router			/api/comics [get]
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
		}).Fatal("Error when requesting /comics")

		panic(err)
	}
}

// GetComic godoc
//
//	@Summary		Get a specific comic
//	@Tags			Comics
//	@Produce		json
//	@Param			id path int	true "Comic ID"
//	@Success		200	{object} model.GetComicResponse
//
// @Router			/api/comic/{id} [get]
func (h *Handlers) GetComicHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Extract variables from the URL
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		h.Sender.JSON(w, http.StatusBadRequest, "Comic ID is required")
		return
	}

	// Convert ID to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Sender.JSON(w, http.StatusBadRequest, "Invalid comic ID format")
		return
	}

	// Verify if the comic exists in the database
	comicExists, err := h.Storage.VerifyComicExists(ctx, id)
	if err != nil {
		logger.Log.Error("Error verifying comic existence:", err)
		h.Sender.JSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	var comic model.Comic

	if comicExists {
		// Fetch the comic from the database
		comic, err = h.Storage.GetComic(ctx, id)
		if err != nil {
			logger.Log.Error("Error fetching comic from database:", err)
			h.Sender.JSON(w, http.StatusInternalServerError, "Internal server error")
			return
		}
	} else {
		comic, err = client.FetchComicFromAPI(ctx, id)
		if err != nil {
			logger.Log.Error("Error fetching comic from external API:", err)
			h.Sender.JSON(w, http.StatusBadGateway, "Failed to fetch comic from external service")
			return
		}

		comic := &model.UpdateComicRequest{
			ID:       comic.ID,
			Title:    comic.Title,
			CoverURL: comic.CoverURL,
			Likes:    0,
		}

		id, err = h.Storage.UpdateComic(ctx, *comic)
		if err != nil {
			logger.Log.Error("Error storing new comic in database:", err)
			h.Sender.JSON(w, http.StatusInternalServerError, "Internal server error")
			return
		}
	}

	comicResponse := model.GetComicResponse{
		ID:       comic.ID,
		Title:    comic.Title,
		CoverURL: comic.CoverURL,
		Likes:    comic.Likes,
	}

	err = h.Sender.JSON(w, http.StatusOK, comicResponse)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatalf("Error sending response for /comic/%d", comic.ID)
	}
}

// UpdateComic godoc
//
//	@Summary		Update a specific comic
//	@Tags			Comics
//	@Param			title body string false "Comic title"
//	@Param			coverUrl body string false "Comic coverUrl"
//	@Success		200	{object} model.IDResponse
//
// @Router			/api/comic/update [patch]
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
		h.Sender.JSON(w, http.StatusBadRequest, "Comic with id="+fmt.Sprint(comic.ID)+" not found")
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
