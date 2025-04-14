package main

import (
	"fmt"
)

func checkEqual(list1 []int, list2 []int) bool {
	for ind := 0; ind < len(list1); ind++ {
		if list1[ind] != list2[ind] {
			return false
		}
	}
	return true
}

func checkSublist(list1 []int, list2 []int) bool {
	for ind := 0; ind <= len(list2)-len(list1); ind++ {
		matched := true
		for i, j := 0, ind; i < len(list1) && j < len(list2); i, j = i+1, j+1 {
			if list1[i] != list2[j] {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}
	return false
}

func Sublist() {
	list1 := []int{1, 1, 2, 1}
	list2 := []int{1, 1, 2}
	len1 := len(list1)
	len2 := len(list2)
	if len1 == len2 && checkEqual(list1, list2) {
		fmt.Println("A and B are equal")
		return
	}
	if len1 < len2 && checkSublist(list1, list2) {
		fmt.Println("A is sublist of B")
		return
	}
	if len1 > len2 && checkSublist(list2, list1) {
		fmt.Println("A is superlist of B")
		return
	}
	fmt.Println("A and B are unequal")
}
