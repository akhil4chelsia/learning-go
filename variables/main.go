package main

import "fmt"

func main() {
	// string type
	var aString string = "Hello"
	fmt.Println(aString)
	fmt.Printf("The variable type is %T\n", aString)

	// integer type
	var anInteger int = 42
	fmt.Println(anInteger)
	fmt.Printf("The variable type is %T\n", anInteger)

	// Default value for int is 0
	var defaultInt int
	fmt.Printf("Default value for int is %d", defaultInt)

	// Variable types are infered while assignment
	var anotherString = "This is another string" // not specified as string type but go will infer as string.
	fmt.Println(anotherString)
	fmt.Printf("The variable type infered is %T\n", anotherString)

	var anotherInt = 50
	fmt.Println(anotherInt)
	fmt.Printf("The variable type infered is %T\n", anotherInt)

	var afloatValue = 25.5
	fmt.Println(afloatValue)
	fmt.Printf("The variable type infered is %T\n", afloatValue)

	const myValue = "a string"
	fmt.Println(myValue)
}
