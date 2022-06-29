package tokengenerator

import (
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/util"
)

type Params struct {
	Rows string `query:"rows"`
}

func TokenCreator(c *fiber.Ctx) error {
	var msg string
	r := new(Params)
	if err := c.QueryParser(r); err != nil {
		return err
	}

	num, err := strconv.ParseInt(r.Rows, 10, 32)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		util.FileGenerator(int(num))
	}()
	wg.Wait()

	msg = strconv.FormatInt(int64(num), 10) + " rows in generated in that file"
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": msg,
	})
}
