package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	var result string
	switch dow {
	case 1:
		result = "It's Monday"
		//fallthrough - use to fallthrough case down
	case 2:
		result = "It's Tuesday"
	default:
		result = "It's some other day."
	}

	fmt.Println(result)
}
