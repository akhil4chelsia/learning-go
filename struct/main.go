package main

import "fmt"

func main() {

	poodle := Dog{"Poodle", 10, "Bow!!"}
	fmt.Printf("Breed is %v \n", poodle.Breed)
	poodle.Speak()
	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

//Copy of Dog object is passed to below function, if you want to pass reference use pointers. (d *Dog)
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
