package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

var rng *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(src)
}

func RandomInt(min, max int) int {
	return rng.Intn(max-min) + min
}

func RandomDecimal(min, max float64) decimal.Decimal {
	random := min + rng.Float64()*(max-min)
	return decimal.NewFromFloat(random).Round(2)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return RandomString(6) + "@email.com"
}

func RandomPassword() string {
	return RandomString(6)
}

func RandomName() string {
	return RandomString(8)
}

func RandomRole() string {
	roles := []string{"admin", "event_organizer", "customer"}
	n := len(roles)
	return roles[rng.Intn(n)]
}

func RandomTicketQuota() int32 {
	return int32(RandomInt(0, 1000))
}

func RandomPrice() decimal.Decimal {
	return RandomDecimal(0, 1000)
}
