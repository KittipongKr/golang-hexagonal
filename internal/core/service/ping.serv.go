package service

import (
	"context"
	"csat-servay/internal/core/port"

	"go.opentelemetry.io/otel/trace"
)

type pingServ struct {
	tracer             trace.Tracer
	pingRepo           port.PingRepo
	jsonplaceholderApi port.JsonplaceholderApi
}

func NewPingServ(
	tracer trace.TracerProvider,
	pingRepo port.PingRepo,
	jsonplaceholderApi port.JsonplaceholderApi,
) port.PingServ {
	return pingServ{
		tracer:             tracer.Tracer("services"),
		pingRepo:           pingRepo,
		jsonplaceholderApi: jsonplaceholderApi,
	}
}

func (s pingServ) GetPongService(ctx context.Context) (string, error) {
	_, span := s.tracer.Start(ctx, "service.GetPongService")
	defer span.End()

	resp, err := s.pingRepo.GetPongRepo()
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (s pingServ) GetJsonplaceholderService(ctx context.Context) error {
	hCtx, span := s.tracer.Start(ctx, "service.GetJsonplaceholderService")
	defer span.End()

	if err := s.jsonplaceholderApi.TestGetEndpoint(hCtx); err != nil {
		return err
	}

	return nil
}
