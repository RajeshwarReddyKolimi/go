package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Exercise22() {
	str := `"That's the password: 'PASSWORD 123'!", cried the Special Agent. So I fled.`
	re := regexp.MustCompile(`\w+('\w+)?`)
	ar := re.FindAllString(strings.ToLower(str), -1)
	freq := make(map[string]int)
	for _, v := range ar {
		freq[v]++
	}
	fmt.Println(freq)
}
