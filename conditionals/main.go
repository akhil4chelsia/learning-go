package main

import "fmt"

func main() {
	var value int = 42
	var result string

	if value < 0 {
		result = "Negative"
	} else if value == 0 {
		result = "Zero"
	} else {
		result = "Positive"
	}

	if value := -42; value < 0 {
		result = "Negative"
	}

	fmt.Println(result)
}
