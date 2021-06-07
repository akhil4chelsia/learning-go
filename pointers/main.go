package main

import "fmt"

func main() {
	anInt := 40
	var p = &anInt
	fmt.Println("Value of p: ", *p)
}
