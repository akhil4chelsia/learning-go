package main

import "fmt"

func main() {
	states := make(map[string]string)
	//insert items
	states["WA"] = "Washinton"
	states["TX"] = "Texas"

	//retrieve item
	fmt.Println(states["TX"])

	//remove item from map
	delete(states, "TX")
	fmt.Println(states)

	//iterate through map
	for k, v := range states {
		fmt.Printf("Key: %v , Value: %v \n", k, v)
	}
}
