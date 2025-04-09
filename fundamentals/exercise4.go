package main

import (
	"bufio"
	"fmt"
	"os"
)

func Exercise4() {
	fmt.Print("Enter file path: ")
	var path string
	fmt.Scanln(&path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
