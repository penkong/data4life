package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder

	l := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}
