package fiber_routes

import (
	envCfgs "csat-servay/configs/env"
	v1 "csat-servay/internal/adapter/fiber/routes/v1"
	"csat-servay/internal/adapter/fiber/tracer"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.opentelemetry.io/otel/trace"
)

func FiberRoute(cfgs *envCfgs.FiberConfig, ctls *Controller, tp trace.TracerProvider) *FiberServer {
	app := fiber.New()
	app.Use(tracer.FiberTrace(tp, nil))

	var credentialBool = map[string]bool{
		"true":  true,
		"1":     true,
		"false": false,
		"0":     false,
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfgs.AllowOrigins,
		AllowHeaders:     cfgs.AllowHeaders,
		AllowMethods:     cfgs.AllowMethods,
		AllowCredentials: credentialBool[cfgs.AllowCredentials],
	}))

	app.Use(logger.New())

	api := app.Group("/api")
	{
		v1.IndexRoute(api, &ctls.V1)
	}

	return &FiberServer{
		address: Server{
			Host: cfgs.Host,
			Port: cfgs.Port,
		},
		application: app,
	}
}

func Launch(serv *FiberServer) error {
	return serv.application.Listen(serv.address.Host + ":" + serv.address.Port)
}
