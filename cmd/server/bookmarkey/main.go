package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	"gitlab.com/bookmarkey/api/internal/collections"
	_ "gitlab.com/bookmarkey/api/migrations"
)

func main() {
	app := pocketbase.New()
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true,
	})

	err := app.Bootstrap()
	if err != nil {
		log.Fatal(err)
	}

	collections.AddHandlers(app)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
