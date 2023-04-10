package bookmarks

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
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

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{ID: authRecord.Id})
	})

	b := new(NewBookmark)
	if err := c.Bind(b); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(b); err != nil {
		return err
	}

	collectionID := c.PathParam(("id"))
	err := h.service.Create(b.URL, collectionID, authRecord.Id)

	if err != nil {
		log.Println("failed to create bookmark: %w", err)
		sentry.CaptureException(err)
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
	authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{ID: authRecord.Id})
	})

	u, err := url.ParseRequestURI(c.QueryParam(("url")))
	if err != nil || (u.Scheme == "" || u.Host == "") {
		if err != nil {
			log.Println(err)
			sentry.CaptureException(err)
		}
		return echo.NewHTTPError(http.StatusBadRequest, "Expected valid URL in query parameter")
	}

	bookmarkMetadata, err := h.service.GetMetadata(u.String())
	if err != nil {
		log.Println("failed to get metadata: %w", err)
		sentry.CaptureException(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get metadata")
	}

	return c.JSON(http.StatusOK, bookmarkMetadata)
}
