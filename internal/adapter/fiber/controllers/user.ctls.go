package controllers

import (
	m "csat-servay/internal/adapter/fiber/models"
	d "csat-servay/internal/core/domain"
	p "csat-servay/internal/core/port"
	"csat-servay/pkg/json"

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	dUser := &d.User{}
	copier.Copy(&dUser, &user)

	userInserted, err := u.userServ.CreateNewService(dUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	copier.Copy(&user, &userInserted)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

// NOTE: get user controllers
func (u UserCtls) GetAllUserEndpoint(c *fiber.Ctx) error {
	user, err := u.userServ.GetAllService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get user",
		})
	}

	resp := []m.User{}
	if err := json.JsoniterMarshalIndent(user, &resp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User retrieved successfully",
		"user":    resp,
	})
}
