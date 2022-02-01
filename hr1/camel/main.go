package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	fmt.Println(camelcase(input))
}

func camelcase(s string) int32 {

	answer := 1
	for _, ch := range s {
		// min, max := 'A', 'Z'
		// if ch >= min && ch <= max {
		// 	answer++
		// }

		str := string(ch)
		if strings.ToUpper(str) == str {
			answer++
		}

	}
	return int32(answer)
}
