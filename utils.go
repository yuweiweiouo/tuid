package tuid

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func NonZeroWeekday(t time.Time) int {
	weekday := t.Weekday()
	if weekday == 0 {
		return 7
	}
	return int(weekday)
}
