package v1

import (
	fCtls "csat-servay/internal/adapter/fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

type V1 struct {
	Auth fCtls.AuthCtls
	User fCtls.UserCtls
}

func IndexRoute(app fiber.Router, ctls *V1) {
	// NOTE: grouping api v1
	v1 := app.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.Post("/one", ctls.Auth.OneAuthEndpoint)
			auth.Post("/refresh", nil)
			auth.Post("/logout", nil)
		}

		user := v1.Group("/users")
		{

			user.Post("/", ctls.User.CreateUserEndpoint)
			user.Post("/bulnk/ra", ctls.User.DumpRaUserEndpoint)
			user.Get("/", ctls.User.GetAllUserEndpoint)
		}

	}

}
