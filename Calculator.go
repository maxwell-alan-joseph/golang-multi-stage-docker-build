package main

import (
	"fmt"
	"os"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b != 0 {
		return a / b
	} else {
		fmt.Println("Error: Division by zero is undefined.")
		os.Exit(1)
	}
	return 0 // this line will never be reached but is needed for the compiler
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Welcome to the Go Calculator!")
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
  fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", add(num1, num2))
	case "-":
		fmt.Printf("Result: %.2f\n", subtract(num1, num2))
	case "*":
		fmt.Printf("Result: %.2f\n", multiply(num1, num2))
	case "/":
		fmt.Printf("Result: %.2f\n", divide(num1, num2))
	default:
		fmt.Println("Invalid operator. Please use one of the following: +, -, *, /")
	}
}
