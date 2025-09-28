package controllers

import (
	hp "csat-servay/internal/adapter/fiber/helpers"
	m "csat-servay/internal/adapter/fiber/models"
	d "csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"
	errs "csat-servay/pkg/errs"
	json "csat-servay/pkg/json"
	logs "csat-servay/pkg/logs"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

type UserCtls struct {
	userServ p.UserServ
}

func NewUserCtls(userServ p.UserServ) UserCtls {
	return UserCtls{
		userServ: userServ,
	}
}

// NOTE: create user controllers
func (u UserCtls) CreateUserEndpoint(c *fiber.Ctx) error {
	user := m.User{}
	if err := c.BodyParser(&user); err != nil {
		logs.Error(err)
		return hp.FailedResponse(c, errs.ErrInvalidInput.Message, fiber.StatusBadRequest)
	}

	resp := &d.User{}
	copier.Copy(&resp, &user)

	userInserted, err := u.userServ.CreateNewService(resp)
	if err != nil {
		logs.Error(err)
		return hp.FailedResponse(c, errs.ErrInternalServer.Message, fiber.StatusInternalServerError)
	}

	copier.Copy(&user, &userInserted)

	return hp.SuccessResponse(c, user)
}

// NOTE: get user controllers
func (u UserCtls) GetAllUserEndpoint(c *fiber.Ctx) error {
	user, err := u.userServ.GetAllService()
	if err != nil {
		logs.Error(err)
		return hp.FailedResponse(c, errs.ErrInternalServer.Message, fiber.StatusInternalServerError)
	}

	resp := []m.User{}
	if err := json.JsoniterMarshalIndent(user, &resp); err != nil {
		logs.Error(err)
		return hp.FailedResponse(c, errs.ErrInternalServer.Message, fiber.StatusInternalServerError)
	}

	return hp.SuccessResponse(c, resp)
}
