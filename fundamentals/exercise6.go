package main

import (
	"bufio"
	"fmt"
	"os"
)

func Exercise6() {
	fmt.Println("Enter file path")
	var path string
	fmt.Scanln(&path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
