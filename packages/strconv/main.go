package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 100
	s := strconv.Itoa(number)
	fmt.Printf("%T %v \n", s, s)
	s2 := "200"
	num2, _ := strconv.Atoi(s2)
	fmt.Printf("%T %v \n", num2, num2)

	s3 := "false"
	r, _ := strconv.ParseBool(s3)
	fmt.Println(r)

	s4 := "10.234"
	r2, _ := strconv.ParseFloat(s4, 64)
	fmt.Println(r2)

	i := 25
	s5 := strconv.FormatInt(int64(i), 10)
	fmt.Println(s5)
}
