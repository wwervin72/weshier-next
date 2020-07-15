package util

import (
	"math/rand"
	"time"
)

// RandomString generate a random str
func RandomString(len int) string {
	bytes := make([]byte, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
