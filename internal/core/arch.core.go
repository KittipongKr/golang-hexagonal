package arch

import (
	envCfgs "csat-servay/configs/env"
	calls "csat-servay/internal/adapter/calls"
	cFiber "csat-servay/internal/adapter/fiber/controllers"
	rFiber "csat-servay/internal/adapter/fiber/routes"
	rFiberV1 "csat-servay/internal/adapter/fiber/routes/v1"
	mongo "csat-servay/internal/adapter/mongo"
	p "csat-servay/internal/core/port"
	serv "csat-servay/internal/core/service"

	"go.opentelemetry.io/otel/trace"
)

type Adaptor struct {
	MongoAdaptor mongo.Adaptor
	CallsAdaptor calls.Adaptor
}

type BusinessLogic struct {
	PingServ p.PingServ
}

func SetAdaptors(
	tp trace.TracerProvider,
	envCfgs envCfgs.EnvConfig,
	a Adaptor,
) BusinessLogic {
	return BusinessLogic{
		PingServ: serv.NewPingServ(
			tp,
			a.MongoAdaptor.PingRepo,
			a.CallsAdaptor.Jsonplaceholder,
		),
	}
}

type Handler struct {
	Router rFiber.Controller
}

func SetHandlers(
	tp trace.TracerProvider,
	logic BusinessLogic,
) Handler {
	return Handler{
		Router: rFiber.Controller{
			V1: rFiberV1.V1{
				Ping: cFiber.NewPingCtls(tp, logic.PingServ),
			},
		},
	}
}
