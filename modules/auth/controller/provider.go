package controller

import (
	"github.com/atom-providers/jwt"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		jwt *jwt.JWT,
	) (*AuthController, error) {
		obj := &AuthController{
			jwt: jwt,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
