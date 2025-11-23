package controllers

import (
	h "csat-servay/internal/adapter/fiber/helpers"
	p "csat-servay/internal/core/port"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

type PingCtls struct {
	tracer   trace.Tracer
	pingServ p.PingServ
}

func NewPingCtls(tracer trace.TracerProvider, pingServ p.PingServ) PingCtls {
	return PingCtls{
		tracer:   tracer.Tracer("handlers"),
		pingServ: pingServ,
	}
}

func (ctl *PingCtls) GetPongEndpoint(ctx *fiber.Ctx) error {
	hCtx, span := ctl.tracer.Start(ctx.UserContext(), "controller.GetPongEndpoint")
	defer span.End()

	resp, err := ctl.pingServ.GetPongService(hCtx)
	if err != nil {
		return h.FailedResponse(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	if err := ctl.pingServ.GetJsonplaceholderService(hCtx); err != nil {
		return h.FailedResponse(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	h.SuccessResponse(ctx, resp)

	return nil
}

func (ctl *PingCtls) GetJsonplaceholderEndpoint(ctx *fiber.Ctx) error {
	hCtx, span := ctl.tracer.Start(ctx.UserContext(), "controller.GetPongEndpoint")
	defer span.End()

	if err := ctl.pingServ.GetJsonplaceholderService(hCtx); err != nil {
		return h.FailedResponse(ctx, err.Error(), fiber.StatusInternalServerError)
	}

	h.SuccessResponse(ctx, "success")

	return nil
}
