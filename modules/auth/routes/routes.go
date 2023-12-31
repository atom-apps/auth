package routes

import (
	"github.com/atom-apps/auth/modules/auth/controller"
	"github.com/atom-providers/log"
	"github.com/gofiber/fiber/v2"
	"github.com/rogeecn/atom"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	return container.Container.Provide(newRoute, atom.GroupRoutes)
}

func newRoute(svc contracts.HttpService, roleController *controller.RoleController, userController *controller.UserController, callbackController *controller.CallbackController, authController *controller.AuthController) contracts.HttpRoute {
	engine := svc.GetEngine().(*fiber.App)
	group := engine.Group("auth")
	log.Infof("register route group: %s", group.(*fiber.Group).Prefix)

	routeAuthController(group, authController)
	routeCallbackController(group, callbackController)
	routeUserController(group, userController)
	routeRoleController(group, roleController)
	return nil
}
