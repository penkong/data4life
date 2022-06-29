package writerto

import (
	"bufio"
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/pkg/connect_db"
	"log"
	"os"
	"sync"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriterToPG(c *fiber.Ctx) error {
	// init needs
	var gNum = 80
	wg := new(sync.WaitGroup)

	ch := make(chan string)

	// reading file
	src, e := os.Open("data.txt")
	check(e)
	defer src.Close()

	// reading each line of file and put it in channel
	go func() {
		defer close(ch)
		// read file to scanner as bufio (buffed)
		bfsc := bufio.NewScanner(src)
		for bfsc.Scan() {
			ch <- bfsc.Text()
		}
	}()

	wg.Add(gNum)
	for i := 0; i < gNum; i++ {
		go func() {
			defer wg.Done()
			for line := range ch {
				connectdb.Pdb.Queries.WriteToken(c.Context(), line)
			}
		}()
	}

	wg.Wait()

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "I am writer to PG",
	})
}
