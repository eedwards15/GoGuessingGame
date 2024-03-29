package Helpers

import (
	"math/rand"
	"time"
)

func GenerateRandom(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
