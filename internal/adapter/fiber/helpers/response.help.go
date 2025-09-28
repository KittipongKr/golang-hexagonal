package helpers

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "",
		"message": "ok",
		"data":    data,
	})
}

func FailedResponse(c *fiber.Ctx, data interface{}, status int) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "fail",
		"message": data,
		"data":    nil,
	})
}
