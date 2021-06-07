package main

import "fmt"

func main() {

	//array example 1
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"
	fmt.Println(colors)
	fmt.Println(colors[1])

	//array example 2, can declare and initialize values in one line
	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(len(nums))

	//slices example, if you dont specify a number it becomes a slice.
	var s = []string{"red", "blue", "green"}
	s = append(s, "Purple")
	fmt.Println(s)

	//using make function
	i := make([]int, 5)
	i[0] = 1
	i[1] = 2
	i[2] = 3
	i[3] = 4
	i[4] = 5
	fmt.Println(i)

	//map example
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
}
