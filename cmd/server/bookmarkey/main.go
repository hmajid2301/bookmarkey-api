// Package main starts our web application and configures everything such as logger and sentry
package main

import (
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"gitlab.com/bookmarkey/api/internal/bookmarks"
	"gitlab.com/bookmarkey/api/internal/middleware"
	_ "gitlab.com/bookmarkey/api/migrations"
)

func main() {
	_ = godotenv.Load()
	app := pocketbase.New()
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	err := setupSentry()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start Sentry")
	}

	err = app.Bootstrap()
	if err != nil {
		log.Fatal().Err(err)
	}

	middleware.ApplyMiddleware(app)
	bookmarks.AddHandlers(app)

	defer sentry.Flush(2 * time.Second)
	log.Info().Msg("starting bookmarkey API service")
	if err := app.Start(); err != nil {
		sentry.CaptureException(err)
		log.Fatal().Err(err)
	}
}

func setupSentry() error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		Environment:      os.Getenv("ENV"),
		TracesSampleRate: 0.8,
		EnableTracing:    true,
	})
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
	return err
}
