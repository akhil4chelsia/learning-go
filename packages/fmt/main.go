package main

import "fmt"

func main() {
	//Print without new line
	fmt.Print("Welcome to Go.")

	//Print with new line
	fmt.Println("Hello")

	//Print string with value
	const age = 32
	fmt.Println("Age is ", age)

	//Print string with multiple values
	const a, b, c = 1, 2, 3
	fmt.Println("one", a, "two", b, "three", c)

	//Print a slice of data
	items := []int{1, 2, 3, 4}
	length, err := fmt.Println(items) //returns no of bytes and error
	fmt.Println(length, err)

	//Basic formating of decimal values.
	const temp = 34.56
	const date = 25
	fmt.Printf("Temperature is : %2f on %d April\n", temp, date) // float & decimal
	fmt.Printf("Temperature in Hexa decimal : %x\n", temp)       // Hexa decimal
	fmt.Printf("Is hot? %t\n", temp > 28)                        // boolean value

	//Explicit argument index
	fmt.Printf("%[2]d %[1]d \n", 1, 2) // will print 2,1

	//Arguments value can be used to print values repeatedly.
	fmt.Printf("%[1]d %[1]d \n", 1)

	type Circle struct {
		radius float64
		area   float64
	}

	circle := Circle{
		radius: 12,
		area:   32.4,
	}
	fmt.Printf("%v\n", circle) // %v prints only value
	fmt.Printf("%v\n", circle) //  %+v print field name with value
	fmt.Printf("%T", circle)   // Prints type of variable

	// fmt.Sprintf - Instead of printing the string, it returns formatted string.
	formatted := fmt.Sprintf("\nType is %[1]T , values are %+[1]v", circle)
	fmt.Println(formatted)

	//Floating point precisions
	number := 12.3456
	fmt.Printf("%.2f\n", number) // .2f - format to 2 decimal place.

	//Print with padding and precision
	fmt.Printf("%10.2f\n", number)

	//Print with padding and precision with + sign |   +12.35
	fmt.Printf("%+10.2f\n", number)

	//Print with padding 0 and precision | 0000012.35
	fmt.Printf("%010.2f\n", number)

}
