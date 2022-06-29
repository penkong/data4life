package writerto

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	// "strings"
	// "sync"

	"github.com/gofiber/fiber/v2"
	// "github.com/penkong/data4life/gateway/pkg/connect_db"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func WriterToPG(c *fiber.Ctx) error {
	// init needs
	var gNum = 8
	wg := new(sync.WaitGroup)

	ch := make(chan string)

	// reading file
	src, e := os.Open("data.txt")
	check(e)
	defer src.Close()

	wg.Add(gNum)

	// reading each line of file and put it in channel
	go func() {
		defer close(ch)
		// read file to scanner as bufio (buffed)
		bfsc := bufio.NewScanner(src)
		for bfsc.Scan() {
			ch <- bfsc.Text()
		}
	}()

	for i := 0; i < gNum; i++ {
		go func() {
			defer wg.Done()
			for line := range ch {
				// connectdb.Pdb.Queries.WriteToken(c.Context(), line)
				fmt.Println(line)
			}
		}()
	}

	wg.Wait()

	// for i := 0; i < gNum; i++ {
	// 	wg.Add(gNum)
	// 	go func() {
	// 		for bfsc.Scan() {
	// 			defer wg.Done()
	// 			connectdb.Pdb.Queries.WriteToken(c.Context(), bfsc.Text())
	// 		}
	// 	}()
	// }
	// wg.Wait()
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
