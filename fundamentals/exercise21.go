package main

import (
	"fmt"
)

func checkEqual(arr1 []int, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func checkSublist(arr1 []int, arr2 []int) bool {
	for ind := 0; ind <= len(arr2)-len(arr1); ind++ {
		flag := true
		for i, j := 0, ind; i < len(arr1) && j < len(arr2); i, j = i+1, j+1 {
			if arr1[i] != arr2[j] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

func Exercise21() {
	arr1 := []int{1, 1, 2, 1}
	arr2 := []int{1, 1, 2}
	l1 := len(arr1)
	l2 := len(arr2)
	if l1 == l2 {
		if checkEqual(arr1, arr2) {
			fmt.Println("A and B are equal")
			return
		}
	} else if l1 < l2 {
		if checkSublist(arr1, arr2) {
			fmt.Println("A is sublist of B")
			return
		}
	} else {
		if checkSublist(arr2, arr1) {
			fmt.Println("A is superlist of B")
			return
		}
	}
	fmt.Println("A and B are unequal")
}
