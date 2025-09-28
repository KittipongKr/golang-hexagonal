package controllers

import (
	h "csat-servay/internal/adapter/fiber/helpers"
	m "csat-servay/internal/adapter/fiber/models"
	d "csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"
	errs "csat-servay/pkg/errs"
	"csat-servay/pkg/logs"

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
		logs.Error(err)
		return h.FailedResponse(ctx, errs.ErrInvalidInput.Message, fiber.StatusBadRequest)
	}

	tokens, err := c.authServ.LoginOneIdService(&d.OneAuthResb{
		Username: authReqb.Username,
		Password: authReqb.Password,
	})
	if err != nil {
		logs.Error(err)
		return h.FailedResponse(ctx, errs.ErrInternalServer.Message, fiber.StatusInternalServerError)
	}

	resp := m.AuthResp{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}

	return h.SuccessResponse(ctx, resp)
}
