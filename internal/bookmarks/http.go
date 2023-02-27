package bookmarks

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

// Servicer used to interact with the service module
type Servicer interface {
	Create(url, collectionID, userID string) error
}

// Handler are all the HTTP handlers
type Handler struct {
	service Servicer
}

// NewTransport returns a new Transport struct to call methods in this module
func NewTransport(service Servicer) Handler {
	return Handler{
		service: service,
	}
}

// NewBookmark struct contains all the fields used to create a new Bookmark
type NewBookmark struct {
	URL string `json:"url" validate:"required,url"`
}

// CustomValidator enables us to use the validator package
type CustomValidator struct {
	validator *validator.Validate
}

// Validate used to validate our HTTP payloads depending on the `validate` struct tags
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// CreateBookmark used to create a new bookmark in the app
func (h Handler) CreateBookmark(c echo.Context) error {
	authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

	// TODO: move to main function ?
	c.Echo().Validator = &CustomValidator{validator: validator.New()}
	b := new(NewBookmark)
	if err := c.Bind(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(b); err != nil {
		return err
	}

	collectionID := c.PathParam(("id"))
	err := h.service.Create(b.URL, collectionID, authRecord.Id)

	if errors.Is(err, ErrNotAuthorized) {
		return apis.NewForbiddenError(err.Error(), nil)
	}

	return c.NoContent(http.StatusCreated)
}
