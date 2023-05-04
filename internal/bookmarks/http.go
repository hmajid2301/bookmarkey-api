// Package bookmarks provides ways to interact with the bookmarks, such as adding a new bookmark
package bookmarks

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
	"github.com/rs/zerolog"
)

// Servicer used to interact with the service module
type Servicer interface {
	Create(url, collectionID, userID string) error
	GetMetadata(url string) (*BookmarkMetaData, error)
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

// CreateBookmark used to create a new bookmark in the app
func (h Handler) CreateBookmark(c echo.Context) error {
	authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

	b := new(NewBookmark)
	if err := c.Bind(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(b); err != nil {
		return err
	}

	collectionID := c.PathParam(("id"))
	err := h.service.Create(b.URL, collectionID, authRecord.Id)

	l := zerolog.Ctx(c.Request().Context())
	if err != nil {
		l.Error().Err(err).Msg("failed to create bookmark")
		if errors.Is(err, ErrNotAuthorized) {
			return apis.NewForbiddenError(err.Error(), nil)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create bookmark")
	}

	return c.NoContent(http.StatusCreated)
}

// GetURLMetadata struct contains all the fields used to create a new Bookmark
type GetURLMetadata struct {
	URL string `json:"url" validate:"required,url"`
}

// GetURLMetadata gets the metadata associated with a URL
func (h Handler) GetURLMetadata(c echo.Context) error {
	l := zerolog.Ctx(c.Request().Context())
	u, err := url.ParseRequestURI(c.QueryParam(("url")))
	if err != nil || (u.Scheme == "" || u.Host == "") {
		if err != nil {
			l.Error().Err(err)
		}
		return echo.NewHTTPError(http.StatusBadRequest, "Expected valid URL in query parameter")
	}

	bookmarkMetadata, err := h.service.GetMetadata(u.String())
	if err != nil {
		l.Error().Err(err).Msg("failed to get metadata")
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get metadata")
	}

	return c.JSON(http.StatusOK, bookmarkMetadata)
}
