package controllers

import (
	m "csat-servay/internal/adapter/fiber/models"
	"csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type AuthCtls struct {
	authServ p.AuthServ
}

func NewAuthCtls(
	authServ p.AuthServ,
) AuthCtls {
	return AuthCtls{
		authServ: authServ,
	}
}

func (c *AuthCtls) OneAuthEndpoint(ctx *fiber.Ctx) error {
	authReqb := m.AuthReqb{}
	if err := ctx.BodyParser(&authReqb); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	tokens, err := c.authServ.LoginOneIdService(&domain.OneAuthResb{
		Username: authReqb.Username,
		Password: authReqb.Password,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	resp := m.AuthResp{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
