package main

import (
	"fmt"
	"time"
)

func main() {

	//Get current date time
	t1 := time.Now()
	fmt.Print("Time is ", t1)

	//Create custom date time object
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

	//Parse date time from string
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
}
