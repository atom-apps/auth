package controller

import (
	"github.com/atom-apps/auth/modules/auth/service"
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
		userMappingSvc *service.UserMappingService,
	) (*CallbackController, error) {
		obj := &CallbackController{
			door:           door,
			userMappingSvc: userMappingSvc,
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
		tenantSvc *service.TenantService,
	) (*TenantController, error) {
		obj := &TenantController{
			tenantSvc: tenantSvc,
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
