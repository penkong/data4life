package apirouters

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetTodos(ctx *fiber.Ctx) error {
	fmt.Println("333")
	// s, err := h.Repo.GetTodo(ctx.Context(), 3)
	// if err != nil {
	// 	return fmt.Errorf("cannot load config:", err)
	// }
	fmt.Println("hello")

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "cannot parse id",
	})
}