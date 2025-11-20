package controllers

import (
	h "csat-servay/internal/adapter/fiber/helpers"
	p "csat-servay/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type PingCtls struct {
	pingServ p.PingServ
}

func NewPingCtls(pingServ p.PingServ) PingCtls {
	return PingCtls{
		pingServ: pingServ,
	}
}

func (ctl *PingCtls) GetPongEndpoint(ctx *fiber.Ctx) error {
	resp, err := ctl.pingServ.GetPongService()
	if err != nil {
		return h.FailedResponse(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	if err := ctl.pingServ.GetJsonplaceholderService(); err != nil {
		return h.FailedResponse(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	h.SuccessResponse(ctx, resp)

	return nil
}

func (ctl *PingCtls) GetJsonplaceholderEndpoint(ctx *fiber.Ctx) error {
	if err := ctl.pingServ.GetJsonplaceholderService(); err != nil {
		return h.FailedResponse(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	h.SuccessResponse(ctx, "success")

	return nil
}
