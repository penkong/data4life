package apirouters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/db/pgdb"
)

type Handlers struct {
	Repo *pgdb.Repo
	*fiber.Ctx
}

func NewHandlers(repo *pgdb.Repo) *Handlers {
	return &Handlers{Repo: repo}
}

func Setup(app *fiber.App, pg *pgdb.Repo) {
	h := NewHandlers(pg)
	v1 := app.Group("/v1/api")
	SetUpTodo(v1, h)
}

func SetUpTodo(r fiber.Router, h *Handlers) {
	todosRoutes := r.Group("/todos")
	todosRoutes.Get("/", GetTodos)
}

func GetTodos(ctx *fiber.Ctx) {
	h.Repo.GetTodo(ctx, 3)
}
