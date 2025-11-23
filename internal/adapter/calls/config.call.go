package calls

import (
	envCfgs "csat-servay/configs/env"
	calls "csat-servay/internal/adapter/calls/api"
	"csat-servay/internal/core/port"

	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel/trace"
)

func SetRrestyClient() *resty.Client {
	return resty.New()
}

type Adaptor struct {
	Jsonplaceholder port.JsonplaceholderApi
}

func SetAdaptor(
	tp trace.TracerProvider,
	client *resty.Client,
	cfgs envCfgs.EnvConfig) Adaptor {
	return Adaptor{
		Jsonplaceholder: calls.NewJsonplaceholderApi(tp, *client),
	}
}
