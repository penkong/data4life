package writerto

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/penkong/data4life/gateway/pkg/connect_db"
	"log"
	"os"
	"strings"
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

	ch := make(chan string)
	// reading each line of file and put it in channel
	go func() {
		defer close(ch)
		// read file to scanner as bufio (buffed)
		bfsc := bufio.NewScanner(src)
		for bfsc.Scan() {
			ch <- bfsc.Text()
		}
	}()

	// -----------------------------------------------

	// with dictionary
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

	valueStrings := []string{}
	values := []interface{}{}

	wg2 := new(sync.WaitGroup)
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		ct := 0
		for k, v := range dict {
			values = append(values, k)
			values = append(values, v)
			valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d)", 2*ct+1, 2*ct+2))
			ct++
		}
	}()
	wg2.Wait()

	fmt.Println("values made!")

	// -----------------------------------------------

	tx, err := connectdb.Pdb.DB.BeginTx(c.Context(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	smt := `INSERT INTO token(name, occur) VALUES %s`
	// ON CONFLICT (name) DO UPDATE SET occur = token.occur + 1;`

	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))
	fmt.Println(smt)
	_, err = tx.Exec(smt, values...)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	// -----------------------------------------------

	// slow write to db
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
