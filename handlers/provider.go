package handlers

import (
	"github.com/atom-apps/auth/handlers/auth"
	"github.com/rogeecn/atom/container"
)

func Providers() container.Providers {
	return container.Providers{
		{Provider: auth.Provide},
	}
}
