package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random number
// between the min and max value
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string
// of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner calls the RandomString
// function to generate a random
// name of the owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney calls the RandomInt function
// and returns and random int64 value in
// the range 0,1000.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency selects a random currency
// out of the three predefined currencies.
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "INR"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
