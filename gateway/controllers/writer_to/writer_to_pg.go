package writerto

import (
	"bufio"
	"log"
	"os"
	"sync"

	// "strings"
	// "sync"

	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/pkg/connect_db"
	// "github.com/penkong/data4life/gateway/pkg/connect_db"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriterToPG(c *fiber.Ctx) error {
	var gNum = 5
	src, e := os.Open("data.txt")
	check(e)
	defer src.Close()
	wg := new(sync.WaitGroup)
	wg.Add(gNum)
	bfsc := bufio.NewScanner(src)
	for i := 0; i < gNum; i++ {
		for bfsc.Scan() {
			go func() {
				connectdb.Pdb.Queries.WriteToken(c.Context(), bfsc.Text())
				defer connectdb.Pdb.Queries.Close()
				defer wg.Done()
			}()
		}
	}
	wg.Wait()
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "I am writer to PG",
	})
}

// jobs := make(chan string)
// results := make(chan int)

// wg := new(sync.WaitGroup)

// go func() {
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 		jobs <- scanner.Text()
// 	}
// 	close(jobs)
// }()

// // Collect all the results...
// // First, make sure we close the result channel when everything was processed
// go func() {
// 	wg.Wait()
// 	close(results)
// }()

// // Now, add up the results from the results channel until closed
// for v := range results {
// 	fmt.Println(v)
// }

// connectdb.Pdb.Queries.WriteToken()
