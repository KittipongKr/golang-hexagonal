package tracer

import (
	"os"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)



func FiberTrace(tp trace.TracerProvider, next func(ctx *fiber.Ctx) bool)  fiber.Handler{
	h, _ := os.Hostname()
	return otelfiber.Middleware(
		otelfiber.WithTracerProvider(tp),
		otelfiber.WithCustomAttributes(func(ctx *fiber.Ctx) []attribute.KeyValue {
			attrs := []attribute.KeyValue{
				attribute.String("http.server.name", h),
				attribute.String("http.request.id", ctx.Get("requestid")),
				attribute.String("http.request.content-type", ctx.Get(fiber.HeaderContentType)),
			}
			return attrs
		}),
		otelfiber.WithNext(next),
	)
}
