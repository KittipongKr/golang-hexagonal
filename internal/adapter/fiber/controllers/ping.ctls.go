package controllers

import "github.com/gofiber/fiber/v2"

type PingCtls struct {
}

func NewPingCtl() PingCtls {
	return PingCtls{}
}

func (c *PingCtls) GetPongEndpoint(ctx *fiber.Ctx) error {
	return ctx.SendString("Pong")
}
