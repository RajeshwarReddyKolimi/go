package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Exercise3() {
	str := `"That's the password: 'PASSWORD 123'!", cried the Special Agent. So I fled.`
	re := regexp.MustCompile(`\w+('\w+)?`)
	lStr := strings.ToLower(str)
	ar := re.FindAllString(lStr, -1)
	freq := make(map[string]int)
	for s := range ar {
		freq[ar[s]]++
	}
	fmt.Println(freq)
}
