package groups

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

// Repository is an interface for interacting with  the database
type Repository interface {
	GetByID(id string) (*models.Record, error)
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

// DeleteGroup delete a users groups given a user ID
func (h Handler) DeleteGroup(c echo.Context) error {
	authRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)

	group, err := h.repo.GetByID(c.PathParam("id"))
	if err != nil {
		return err
	}

	groupOwner := group.GetStringSlice("user")[0]
	if authRecord.Id != groupOwner {
		return apis.NewForbiddenError("The user does not have permission to delete group.", nil)
	}

	if err := h.repo.Delete(group); err != nil {
		// TODO: Log error properly
		// return err
		return apis.NewApiError(http.StatusInternalServerError, "Failed to delete group.", nil)
	}

	return c.JSON(http.StatusOK, map[string]string{"name": group.GetString("name"), "message": "Successfully deleted group."})
}
