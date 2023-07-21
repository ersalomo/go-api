package utils

import (
	"math/rand"
	"strings"
	"time"
)

const letter = "abcdefghijklmnopqrstuvwyxyz"

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min) + 1
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(letter)

	for i := 0; i <= n; i++ {
		c := letter[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"IDR", "UER", "CAD", "SGD"}

	length := len(currencies)
	return currencies[rand.Intn(length)]
}
