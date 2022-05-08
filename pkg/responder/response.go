package responder

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type errorResponse struct {
	Message string      `json:"message"`
	Fields  interface{} `json:"fields,omitempty"`
}

func Send(c *fiber.Ctx, obj interface{}) error {
	c.Set("Content-Type", "application/json")
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return c.Send(b)
}

func Error(c *fiber.Ctx, code int, message string) error {
	c.Set("Content-Type", "application/json")
	b, err := json.Marshal(&errorResponse{
		Message: message,
	})
	if err != nil {
		return err
	}

	c.JSON(b)
	return c.SendStatus(code)
}

func ValidationError(c *fiber.Ctx, code int, message string, fields interface{}) error {
	b, err := json.Marshal(&errorResponse{
		Message: message,
		Fields:  fields,
	})
	if err != nil {
		return err
	}

	c.JSON(b)
	return c.SendStatus(code)
}
