//go:generate atomctl gen routes
//go:generate swag fmt
//go:generate swag init -ot json
package main

import (
	"log"

	"github.com/atom-apps/auth/modules/auth"
	"github.com/atom-providers/jwt"
	serviceHttp "github.com/atom-providers/service-http"
	"github.com/rogeecn/atom"
	"github.com/spf13/cobra"
)

func main() {
	providers := serviceHttp.Default(
		jwt.DefaultProvider(),
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
