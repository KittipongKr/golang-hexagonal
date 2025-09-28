package arch

import (
	envCfgs "csat-servay/configs/env"
	calls "csat-servay/internal/adapter/calls"
	cFiber "csat-servay/internal/adapter/fiber/controllers"
	rFiber "csat-servay/internal/adapter/fiber/routes"
	rFiberV1 "csat-servay/internal/adapter/fiber/routes/v1"
	mongo "csat-servay/internal/adapter/mongo"
	p "csat-servay/internal/core/port"
	"csat-servay/internal/core/service"
)

type Adaptor struct {
	MongoAdaptor mongo.Adaptor
	CallsAdaptor calls.Adaptor
}

type BusinessLogic struct {
	Auth p.AuthServ
	User p.UserServ
}

func SetAdaptors(envCfgs envCfgs.EnvConfig, a Adaptor) BusinessLogic {
	return BusinessLogic{
		Auth: service.NewAuthServ(
			envCfgs.App,
			envCfgs.Jwt,
			a.MongoAdaptor.User,
			a.CallsAdaptor.One,
		),
		User: service.NewUserServ(
			a.MongoAdaptor.User,
		),
	}
}

type Handler struct {
	Router rFiber.Controller
}

func SetHandlers(
	logic BusinessLogic,
) Handler {
	return Handler{
		Router: rFiber.Controller{
			V1: rFiberV1.V1{
				Auth: cFiber.NewAuthCtls(logic.Auth),
				User: cFiber.NewUserCtls(logic.User),
			},
		},
	}
}
