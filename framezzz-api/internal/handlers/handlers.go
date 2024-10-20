package handlers

import (
	"github.com/CyberBoyzzz/Framezzz/internal/storage"
	"github.com/CyberBoyzzz/Framezzz/pkg/httputils"
	"github.com/go-playground/validator/v10"
)

// Handlers implements all the handler functions and has the dependencies that they use (Sender, Storage).
type Handlers struct {
	Sender  *httputils.Sender
	Storage storage.StorageInterface
}

// Validate is a singleton that provides validation services for in handlers.
var Validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
