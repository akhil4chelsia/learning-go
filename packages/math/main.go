package main

import (
	"fmt"
	"math"
)

func main() {
	//Pi constant
	fmt.Println(math.Pi)

	//floor and ceil
	fmt.Println(math.Ceil(math.Pi))
	fmt.Println(math.Floor(math.Pi))

	//Trancate will return int value of x
	fmt.Printf("%.2f \n", math.Trunc(math.Pi))

	//Max and Min
	fmt.Println(math.Max(math.Pi, math.Ln2))
	fmt.Println(math.Min(math.Pi, math.Ln2))

	//Mod operator
	fmt.Println(17 % 5)               // this can be only used with integers
	fmt.Println(math.Mod(17.23, 4.2)) // use math.Mod() for float remainders

	//Round & RoundToEven
	fmt.Println(math.Round(math.Pi))
	fmt.Println(math.RoundToEven(math.Pi))

	x := 10.3
	//Absolute value
	fmt.Println(math.Abs(x))
	fmt.Println(math.Abs(-x))
	//Pow function
	fmt.Println(math.Pow(2, 8))
	fmt.Println(math.E, math.Exp(10)) // e^10

	//Trignometric functions
	fmt.Println(math.Cos(math.Pi))
	fmt.Println(math.Sin(math.Pi / 2))
	fmt.Println(math.Log(10))

	//Square root and cube root
	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Cbrt(125))

	fmt.Println(math.Hypot(30, 40)) //pythogorous theorem
}
