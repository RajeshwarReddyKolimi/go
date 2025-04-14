package main

import "fmt"

type Calculator interface {
	add(a, b float64) float64
	subtract(a, b float64) float64
	multiply(a, b float64) float64
	divide(a, b float64) float64
}

func add(a, b float64) float64 {
	sum := a + b
	fmt.Println("The sum is ", sum)
	return sum
}

func subtract(a, b float64) float64 {
	difference := a - b
	fmt.Println("The difference is ", difference)
	return difference
}

func multiply(a, b float64) float64 {
	product := a * b
	fmt.Println("The product is ", product)
	return product
}

func divide(a, b float64) float64 {
	if b == 0 {
		e := fmt.Errorf("Error: Division by zero")
		fmt.Println(e)
		return 0
	}
	quotient := a / b
	fmt.Println("The quotient is ", quotient)
	return quotient
}

func Exercise1() {
	var n1 float64
	var n2 float64
	var op string
	fmt.Println("Enter first number")
	fmt.Scanln(&n1)
	fmt.Println("Enter second number")
	fmt.Scanln(&n2)
	fmt.Println("Enter operation (+, -, *, /)")
	fmt.Scanln(&op)
	switch op {
	case "+":
		add(n1, n2)
	case "-":
		subtract(n1, n2)
	case "*":
		multiply(n1, n2)
	case "/":
		divide(n1, n2)
	default:
		fmt.Printf("No operator selected\n")
	}
}
