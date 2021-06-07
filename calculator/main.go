package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculator 0.2")
	value1 := GetInput("Enter value 1: ")
	value2 := GetInput("Enter value 2: ")
	operation := GetOperation()
	calculator := Calculator{value1, value2, operation}
	result := calculator.Calculate()
	fmt.Printf("Result : %v", result)
}

type Calculator struct {
	Value1    float64
	Value2    float64
	Operation string
}

func (c Calculator) Calculate() float64 {
	result := 0.0
	switch c.Operation {
	case "-":
		result = c.subtract()
	case "*":
		result = c.multiply()
	case "/":
		result = c.divide()
	default:
		result = c.add()
	}
	return result
}

func (c Calculator) add() float64 {
	return c.Value1 + c.Value2
}

func (c Calculator) subtract() float64 {
	return c.Value1 - c.Value2
}

func (c Calculator) multiply() float64 {
	return c.Value1 * c.Value2
}

func (c Calculator) divide() float64 {
	return c.Value1 / c.Value2
}

func GetInput(message string) float64 {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	floatVal, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		fmt.Println("Value must be integer")
		panic("Unexpected value error.")
	}
	return floatVal
}

func GetOperation() string {
	fmt.Print("Select operation, +, -, *, / :default(+): ")
	reader := bufio.NewReader(os.Stdin)
	operation, _ := reader.ReadString('\n')
	return strings.TrimSpace(operation)
}
