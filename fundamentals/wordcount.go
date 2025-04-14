package main

import (
	"fmt"
	"regexp"
	"strings"
)

func WordCount() {
	str := `"That's the password: 'PASSWORD 123'!", cried the Special Agent. So I fled.`
	re := regexp.MustCompile(`\w+('\w+)?`)
	words := re.FindAllString(strings.ToLower(str), -1)
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	fmt.Println(freq)
}
