package writerto

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	pg "github.com/lib/pq"
	connectdb "github.com/penkong/data4life/gateway/pkg/connect_db"
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

	// reading file
	src, e := os.Open("data.txt")
	check(e)
	defer src.Close()

	// -----------------------------------------------

	// reading each line of file and put it in channel
	ch := make(chan string)
	go func() {
		defer close(ch)
		// read file to scanner as bufio (buffed)
		bfsc := bufio.NewScanner(src)
		for bfsc.Scan() {
			ch <- bfsc.Text()
		}
	}()

	// -----------------------------------------------

	// Dictionary usage , listen on channel and make dictionary with inputs
	dict := make(map[string]int64)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for s := range ch {
			if _, exist := dict[s]; exist {
				dict[s] = dict[s] + 1
			} else {
				dict[s] = 1
			}
		}
	}()
	wg.Wait()

	fmt.Println("dictionary made")

	// -----------------------------------------------

	// prepare dictionary for bulk insert with 2 slices
	keys := []interface{}{}
	values := []interface{}{}

	wg2 := new(sync.WaitGroup)
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for k, v := range dict {
			keys = append(keys, k)
			values = append(values, int(v))
		}
	}()
	wg2.Wait()

	fmt.Println("values made!")

	// -----------------------------------------------

	// Transactional Bulk insert with unnest pattern
	tx, err := connectdb.Pdb.DB.BeginTx(c.Context(), &sql.TxOptions{})
	check(err)
	defer tx.Rollback()

	query := `
  INSERT INTO token
    (name, occur)
    (
      select * from unnest($1::TEXT[], $2::INT[])
    )`

	_, err = tx.Query(query, pg.Array(keys), pg.Array(values))
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	// -----------------------------------------------

	// slow write to db - bad one
	// gNum := 80
	// wg := new(sync.WaitGroup)
	// wg.Add(gNum)
	// for i := 0; i < gNum; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		for line := range ch {
	// 			connectdb.Pdb.Queries.WriteToken(c.Context(), line)
	// 		}
	// 	}()
	// }
	// wg.Wait()

	// -----------------------------------------------

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "wrote to pg done!!!",
	})
}
