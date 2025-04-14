package main

import (
	"fmt"
)

func ETL() {
	grp := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}
	res := make(map[string]int)
	for point, alphabets := range grp {
		for _, alphabet := range alphabets {
			res[alphabet] = point
		}
	}
	fmt.Println(res)
}
