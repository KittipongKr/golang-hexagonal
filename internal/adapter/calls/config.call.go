package calls

import (
	envCfgs "csat-servay/configs/env"
	calls "csat-servay/internal/adapter/calls/api"
	p "csat-servay/internal/core/port"

	"github.com/go-resty/resty/v2"
)

func SetRrestyClient() *resty.Client {
	return resty.New()
}

type Adaptor struct {
	One p.OneCall
}

func SetAdaptor(
	client *resty.Client,
	cfgs envCfgs.EnvConfig) Adaptor {
	return Adaptor{
		One: calls.NewOneCall(client, cfgs.One),
	}
}
