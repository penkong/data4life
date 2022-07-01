package reporter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/db/pgdb"
	connectdb "github.com/penkong/data4life/gateway/pkg/connect_db"
)

type Params struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

func ReporterPG(c *fiber.Ctx) error {

	r := new(Params)
	if err := c.QueryParser(r); err != nil {
		return err
	}

	l := r.Limit | 50
	o := r.Offset | 0

	count, err := connectdb.Pdb.Queries.CountNonUnique(c.Context())
	if err != nil {
		return err
	}

	rows, err := connectdb.Pdb.Queries.ReadNonUnique(c.Context(), pgdb.ReadNonUniqueParams{
		Limit:  int32(l),
		Offset: int32(o),
	})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"count":   count,
		"payload": rows,
	})
}
