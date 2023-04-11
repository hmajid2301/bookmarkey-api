// Package middleware provides different middlewares than can be used by the http transports
package middleware

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/pocketbase/pocketbase/core"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

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

// ApplyMiddleware adds all the generic middleware to our app
func ApplyMiddleware(app core.App) {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(middleware.Logger())
		e.Router.Use(NewSentry(SentryOptions{}))
		e.Router.Validator = &CustomValidator{validator: validator.New()}
		return nil
	})
}
