package call_api

import (
	"context"
	p "csat-servay/internal/core/port"

	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type jsonplaceholder struct {
	tracer trace.Tracer
	client resty.Client
}

func NewJsonplaceholderApi(tracer trace.TracerProvider, client resty.Client) p.JsonplaceholderApi {
	return jsonplaceholder{
		tracer: tracer.Tracer("calls"),
		client: client,
	}
}

func (call jsonplaceholder) TestGetEndpoint(ctx context.Context) error {
	_, span := call.tracer.Start(ctx, "call.TestGetEndpoint")
	defer span.End()

	resp, err := call.client.R().
		Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return err
	}

	span.SetAttributes(attribute.String("response", resp.String()))

	return nil
}
