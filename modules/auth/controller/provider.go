package controller

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		door *casdoorsdk.Client,
	) (*AuthController, error) {
		obj := &AuthController{
			door: door,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		door *casdoorsdk.Client,
	) (*CallbackController, error) {
		obj := &CallbackController{
			door: door,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		door *casdoorsdk.Client,
	) (*RoleController, error) {
		obj := &RoleController{
			door: door,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		door *casdoorsdk.Client,
	) (*UserController, error) {
		obj := &UserController{
			door: door,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
