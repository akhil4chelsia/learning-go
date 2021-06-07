package main

import (
	"fmt"
	"math"
)

func main() {
	circleRadius := 15.5
	circumferance := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f ", circumferance)
}
