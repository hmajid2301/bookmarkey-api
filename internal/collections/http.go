package collections

import (
	"encoding/json"
	"errors"
	"fmt"
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

// AddCollection add a users collections
func (h Handler) AddCollection(c echo.Context) error {
	authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

	type NewCollection struct {
		Name string `json:"collection_name"`
	}

	newCollection := NewCollection{}
	err := json.NewDecoder(c.Request().Body).Decode(&newCollection)
	if err != nil {
		return apis.NewApiError(http.StatusBadRequest, "Failed to decode payload when trying to add collection.", nil)
	}
	if newCollection.Name == "" {
		return apis.NewApiError(http.StatusBadRequest, "Missing `collection_name` field in request.", nil)
	}

	err = h.repo.Add(newCollection.Name, authRecord.Id)
	if err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			return apis.NewApiError(http.StatusConflict, fmt.Sprintf("A collection already exists with name the %s.", newCollection.Name), nil)

		}
		// return err
		return apis.NewApiError(http.StatusInternalServerError, "Failed to add collection.", nil)
	}

	return c.JSON(http.StatusOK, map[string]string{"name": newCollection.Name, "message": "Successfully created collection."})
}
