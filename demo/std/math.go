package std

import (
	"fmt"
	"math/rand"
	"time"
)

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func DemoRand() bool {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(RandInt(100, 999))
	fmt.Printf("rand num: %d\n", num)

	return true
}
