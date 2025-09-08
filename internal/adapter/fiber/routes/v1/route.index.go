package v1

import (
	fCtls "csat-servay/internal/adapter/fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

type V1 struct {
	Ping fCtls.PingCtls
	User fCtls.UserCtls
}

func IndexRoute(app fiber.Router, ctls *V1) {
	// NOTE: grouping api v1
	v1 := app.Group("/v1")
	{
		ping := v1.Group("/ping")
		{
			ping.Get("/", ctls.Ping.GetPongEndpoint)
		}

		user := v1.Group("/users")
		{
			user.Post("/", ctls.User.CreateUserEndpoint)
			user.Get("/", ctls.User.GetAllUserEndpoint)
		}
	}

}
