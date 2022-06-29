package util

import (
	"log"
	"os"
)

func Generate(f *os.File, l string) {
	_, err2 := f.WriteString(l + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}
}
