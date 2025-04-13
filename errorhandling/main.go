package main

import (
	"flag"
	"fmt"
	"strconv"
)

func calculateSum() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
		}
	}()
	flag.Parse()
	arr := flag.Args()
	sum := 0
	if len(arr) == 0 {
		panic("No numbers provided")
	}
	for _, str := range arr {
		n, er := strconv.Atoi(str)
		if er != nil {
			panic("Provided input is not a number")
		}
		sum += n
	}
	fmt.Println(sum)

	defer func() {
		fmt.Println("This function executes if there is no panic")
	}()
}
func main() {
	calculateSum()
	fmt.Println("Calculation done")
}
