//go:generate atomctl gen routes
//go:generate swag fmt
//go:generate swag init -ot json
package main

import (
	"log"

	_ "github.com/go-gormigrate/gormigrate/v2"

	"github.com/atom-apps/auth/database/query"
	"github.com/atom-apps/auth/modules/auth"
	"github.com/atom-providers/casdoor"
	"github.com/atom-providers/cert"
	databasePostgres "github.com/atom-providers/database-postgres"
	serviceHttp "github.com/atom-providers/service-http"
	"github.com/rogeecn/atom"
	"github.com/spf13/cobra"
)

func main() {
	providers := serviceHttp.Default(
		query.DefaultProvider(),
		cert.DefaultProvider(),
		databasePostgres.DefaultProvider(),
		casdoor.DefaultProvider(),
	).With(
		auth.Providers(),
	)
	opts := []atom.Option{
		atom.Name("auth"),
		atom.RunE(func(cmd *cobra.Command, args []string) error {
			return serviceHttp.Serve()
		}),
	}

	if err := atom.Serve(providers, opts...); err != nil {
		log.Fatal(err)
	}
}
