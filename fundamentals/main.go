package main

import (
	"fmt"
)

func main() {
	var n1 int
	var n2 int
	var op string
	fmt.Println("Enter first number")
	fmt.Scanln(&n1)
	fmt.Println("Enter second number")
	fmt.Scanln(&n2)
	fmt.Println("Enter operation (+, -, *, /)")
	fmt.Scanln(&op)
	switch op {
	case "+":
		fmt.Printf("Sum of %v, %v is %v\n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("Difference of %v, %v is %v\n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("Product of %v, %v is %v\n", n1, n2, n1*n2)
	case "/":
		{
			if n2 == 0 {
				e := fmt.Errorf("Error: Division by zero")
				fmt.Println(e)
			} else {
				fmt.Printf("Quotient of %v, %v is %v\n", n1, n2, n1/n2)
			}
		}
	default:
		fmt.Printf("No operator selected\n")
	}
}
