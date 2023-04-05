package api

import (
	"github.com/gofiber/fiber/v2"
)

func ListRooms(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
