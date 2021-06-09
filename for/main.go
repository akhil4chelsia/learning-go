package main

import "fmt"

func main() {
	states := []string{"LA", "TX", "NY"}

	//Normal for loop with index variable
	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}

	//Advanced forEach loop
	for _, value := range states {
		fmt.Println(value)
	}

	// Eqalent to while loop
	value := 1
	for value <= 10 {
		fmt.Println(value)
		value++
		if value == 5 {
			break // break and continue
		}
	}

	// Equalent to while true loop
	value := 1
	for {
		if value == 10 {
			break
		}
		value += 1
	}

	// Goto statement
	value2 := 1
	for value2 <= 10 {
		fmt.Println(value2)
		value2++
		if value2 == 5 {
			goto theLabel
		}
	}

theLabel:
	fmt.Println("theLabel:", value2)
}
