package tokengenerator

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/util"
)

type Params struct {
	Rows string `query:"rows"`
}

func TokenCreator(c *fiber.Ctx) error {
	r := new(Params)
	if err := c.QueryParser(r); err != nil {
		return err
	}

	num, err := strconv.ParseInt(r.Rows, 10, 64)
	if err != nil {
		return err
	}

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	n := int(num)

	// ---------------------------------------------

	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- util.RandomString(7)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for l := range ch {
				util.Generate(f, l)
			}
		}()
	}
	wg.Wait()

	// ----------------------------------------------

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": strconv.FormatInt(int64(num), 10) + " rows in generated in that file",
	})
}
