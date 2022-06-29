package util

import (
	"log"
	"os"
	"sync"
)

func FileGenerator(record int) {
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var wg sync.WaitGroup

	wg.Add(record)
	for i := 0; i < record; i++ {
		go func() {
			defer wg.Done()
			generate(f)
		}()
		// fmt.Println(strconv.Itoa(runtime.NumGoroutine()))
	}
	wg.Wait()
}

func generate(f *os.File) {
	_, err2 := f.WriteString(RandomString(7) + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}
}
