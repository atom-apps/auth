package controller

import (
	"encoding/json"

	"github.com/atom-apps/auth/database/models"
	"github.com/atom-apps/auth/modules/auth/dto"
	"github.com/atom-apps/auth/modules/auth/service"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
)

// @provider
type WebhookController struct {
	door           *casdoorsdk.Client
	userMappingSvc *service.UserMappingService
}

// Signup
//
//	@Summary		signup
//	@Description	signup
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		301	{string}	string	""
//	@Router			/auth/webhook/signup [post]
func (c *WebhookController) Signup(ctx *fiber.Ctx, body *dto.WebhookForm) error {
	var data *dto.WebhookObject
	if err := json.Unmarshal([]byte(body.Object), data); err != nil {
		return err
	}

	user, err := c.door.GetUser(data.Name)
	if err != nil {
		return err
	}

	return c.userMappingSvc.CreateFromModel(ctx.Context(), &models.UserMapping{
		UUID: user.Id,
	})
}
