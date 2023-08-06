package controller

import (
	"github.com/atom-apps/auth/modules/auth/dto"
	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
)

// @provider
type AuthController struct {
	jwt *jwt.JWT
}

// Login by user info
//
//	@Summary		login
//	@Description	login
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	 dto.AuthLoginResponse
//	@Router			/auths/login [post]
func (c *AuthController) Login(ctx *fiber.Ctx) (dto.AuthLoginResponse, error) {
	claim := c.jwt.CreateClaims(jwt.BaseClaims{
		UserID:   1,
		TenantID: 1,
		Role:     "user",
	})

	token, err := c.jwt.CreateToken(claim)
	if err != nil {
		return dto.AuthLoginResponse{}, err
	}

	return dto.AuthLoginResponse{
		Token:    token,
		ExpireAt: claim.ExpiresAt.Time,
	}, nil
}
