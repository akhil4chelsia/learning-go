package main

import "fmt"

func main() {
	doSomething()
	sum := add(1, 2)
	fmt.Println(sum)

	allsum, length := addAll(1, 2, 3, 4, 5)
	fmt.Printf("Sum: %v, Length: %v", allsum, length)
}

func doSomething() {
	fmt.Println("Doing something...")
}

// if function accepts same type of argument, then only need to declare once.
func add(value1, value2 int) int {
	return value1 + value2
}

//Function can accept variable aruments (...) and can return multiple values.
func addAll(values ...int) (int, int) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, len(values)
}
