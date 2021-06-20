package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//Provide a seed value to initialize
	rand.Seed(time.Now().UnixNano())

	//Generate random numbers
	fmt.Println(rand.Int())     // int
	fmt.Println(rand.Intn(10))  // upto 10
	fmt.Println(rand.Float32()) // float32

	//Generate a random order to select items from slice
	items := []string{"apple", "orange", "kiwi", "grape", "melon", "pear"}
	indexes := rand.Perm(len(items))
	fmt.Println(indexes)
	for _, i := range indexes {
		fmt.Println(items[i])
	}

	//Generate random number between a and b
	a, b := 10, 50
	for i := 0; i < 10; i++ {
		n := a + rand.Intn(b-a+1)
		fmt.Println(n)
	}

	//Generate a random uppercase character
	for i := 0; i < 10; i++ {
		c := string('A' + rune(rand.Intn('Z'-'A'+1)))
		fmt.Println(c)
	}

	//Shuffle sequence of values
	str := "one two three four five six seven eight nine ten"
	words := strings.Fields(str)
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}
