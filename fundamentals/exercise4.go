package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	return nil
}
func Exercise4() {
	fmt.Print("Enter file path: ")
	var path string
	fmt.Scanln(&path)

	if err := readFile(path); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Read file successfully")
}
