package fiber_routes

import (
	v1 "csat-servay/internal/adapter/fiber/routes/v1"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	address     Server
	application *fiber.App
}

type Server struct {
	Host string
	Port string
}

type Controller struct {
	V1 v1.V1
}
