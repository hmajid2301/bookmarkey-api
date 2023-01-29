package groups

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// AddHandlers sets up the http handlers
func AddHandlers(app core.App) {
	store := NewStore(app)
	transport := NewTransport(store)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.DELETE("/groups/:id", func(c echo.Context) error {
			return transport.DeleteGroup(c)
		},
			apis.ActivityLogger(app),
			apis.RequireRecordAuth(),
		)
		return nil
	})
}
