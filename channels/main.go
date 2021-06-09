package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) //Creating unbuffered channel
	//ch := make(chan int, 5) //Creating buffered channel with capacity 5

	go func() {
		for i := 0; i < 15; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch) // close the channel
	}()

	// for i := 0; i < 5; i++ {
	// 	val := <-ch
	// 	fmt.Printf("Received from channel: %d\n", val)
	// }

	// get values from channel using range function of channel
	for val := range ch {
		fmt.Printf("Received from channel: %d\n", val)
	}
}
