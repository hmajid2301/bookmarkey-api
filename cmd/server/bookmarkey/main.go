// Package main starts our web application and configures everything such as logger and sentry
package main

import (
	"os"

	beeline "github.com/honeycombio/beeline-go"
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/rs/zerolog/log"

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

	err := app.Bootstrap()
	if err != nil {
		log.Fatal().Err(err)
	}
	beeline.Init(beeline.Config{
		WriteKey:    os.Getenv("HONEYCOMB_API_KEY"),
		Dataset:     os.Getenv("SERVICE_NAME"),
		ServiceName: os.Getenv("SERVICE_NAME"),
	})
	middleware.ApplyMiddleware(app)
	bookmarks.AddHandlers(app)

	defer beeline.Close()

	log.Info().Msg("starting bookmarkey API service")
	if err := app.Start(); err != nil {
		log.Fatal().Err(err)
	}
}
