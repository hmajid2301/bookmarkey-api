package bookmarks

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"gitlab.com/bookmarkey/api/internal/collections"
)

// AddHandlers sets up the http handlers
func AddHandlers(app core.App) {
	collStore := collections.NewStore(app)
	collSrv := collections.NewService(collStore)

	store := NewStore(app)
	srv := NewService(store, collSrv)
	transport := NewTransport(srv)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/collections/:id/bookmark", func(c echo.Context) error {
			return transport.CreateBookmark(c)
		},
			apis.ActivityLogger(app),
			apis.RequireRecordAuth(),
		)
		return nil
	})
}
