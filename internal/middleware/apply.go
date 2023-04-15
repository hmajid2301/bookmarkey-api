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
		e.Router.Use(middleware.Recover())
		e.Router.Use(middleware.Logger())
		e.Router.Use(middleware.RequestID())
		e.Router.Use(NewSentry(SentryOptions{}))
		e.Router.Use(echo.WrapMiddleware(logger))
		// e.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		// 	LogURI:    true,
		// 	LogStatus: true,
		// 	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
		// 		l := zerolog.Ctx(c.Request().Context())
		// 		l.Info().
		// 			Time("start_time", v.StartTime).
		// 			Str("host", v.Host).
		// 			Int64("latency_ms", v.Latency.Milliseconds()).
		// 			Str("latency_human", v.Latency.String()).
		// 			Str("remote_ip", v.RemoteIP).
		// 			Str("user_agent", v.UserAgent).
		// 			Str("uri", v.URI).
		// 			Str("method", v.Method).
		// 			Int("status", v.Status).
		// 			Msg("received request")

		// 		return nil
		// 	},
		// }))

		e.Router.Validator = &CustomValidator{validator: validator.New()}
		return nil
	})
}
