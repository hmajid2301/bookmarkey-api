package main

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"gitlab.com/bookmarkey/api/internal/bookmarks"
	"gitlab.com/bookmarkey/api/internal/middleware"
	_ "gitlab.com/bookmarkey/api/migrations"
)

func main() {
	godotenv.Load()
	app := pocketbase.New()
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Environment:      os.Getenv("ENV"),
		TracesSampleRate: 1.0,
		EnableTracing:    true,
	})
	if err != nil {
		log.Fatalf("failed to start Sentry: %s", err)
	}

	err = app.Bootstrap()
	if err != nil {
		log.Fatal(err)
	}

	middleware.ApplyMiddleware(app)
	bookmarks.AddHandlers(app)

	defer sentry.Flush(2 * time.Second)
	if err := app.Start(); err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}
}
