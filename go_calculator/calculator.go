package main

import (
	"fmt"
	"os"
	"strconv"
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
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run calculator.go <num1> <operator> <num2>")
		fmt.Println("Operators: +, -, *, /")
		return
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operator := os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid numbers")
		return
	}

	var result float64

	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result = divide(num1, num2)
	default:
		fmt.Println("Error: Invalid operator. Use +, -, *, /")
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}