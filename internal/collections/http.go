package collections

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

// Repository is an interface for interacting with  the database
type Repository interface {
	GetByID(id string) (*models.Record, error)
	Add(name string, userID string) error
	Delete(record *models.Record) error
}

// Handler used to interact with collections
type Handler struct {
	repo Repository
}

// NewTransport returns a new Transport struct to call methods in this package
func NewTransport(repo Repository) Handler {
	return Handler{
		repo: repo,
	}
}

// DeleteCollection delete a users collections given a user ID
func (h Handler) DeleteCollection(c echo.Context) error {
	authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

	collection, err := h.repo.GetByID(c.PathParam("id"))
	if err != nil {
		return err
	}

	collectionOwner := collection.GetStringSlice("user")[0]
	if authRecord.Id != collectionOwner {
		return apis.NewForbiddenError("The user does not have permission to delete collection.", nil)
	}

	if err := h.repo.Delete(collection); err != nil {
		// TODO: Log error properly
		// return err
		return apis.NewApiError(http.StatusInternalServerError, "Failed to delete collection.", nil)
	}

	return c.JSON(http.StatusOK, map[string]string{"name": collection.GetString("name"), "message": "Successfully deleted collection."})
}
