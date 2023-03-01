package main

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"gitlab.com/bookmarkey/api/internal/bookmarks"
	_ "gitlab.com/bookmarkey/api/migrations"
)

func main() {
	app := pocketbase.New()
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Environment:      os.Getenv("ENV"),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("failed to start Sentry: %s", err)
	}

	err = app.Bootstrap()
	if err != nil {
		log.Fatal(err)
	}

	bookmarks.AddHandlers(app)

	if err := app.Start(); err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}
}
