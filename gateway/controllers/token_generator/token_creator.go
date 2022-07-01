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

	// select waitgroup and worker tested , waitgroup is fastest one , seems .
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

	// Worker pattern

	// const numJobs = 5
	// jobs := make(chan string, n)
	// results := make(chan string, n)

	// for w := 1; w <= 10000; w++ {
	// 	go worker(w, jobs, results)
	// }

	// for j := 1; j <= n; j++ {
	// 	jobs <- util.RandomString(7)
	// }
	// close(jobs)

	// for a := 1; a <= n; a++ {
	// 	util.Generate(f, <-results)
	// }

	// ----------------------------------------------

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": strconv.FormatInt(int64(num), 10) + " rows in generated in that file",
	})
}

// func worker(id int, jobs <-chan string, results chan<- string) {
// 	for j := range jobs {
// 		results <- j
// 	}
// }
