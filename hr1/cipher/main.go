package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int32
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	// fmt.Printf("length: %d\n", length)
	// fmt.Printf("input: %s\n", input)
	// fmt.Printf("delta: %d\n", delta)

	fmt.Println(caesarCipher(input, delta))

	// fmt.Println(string(rotate('z', 2, alphabet)))
}

func rotate(s rune, delta int32, key []rune) rune {

	// --------------------------------------------
	idx := int32(strings.IndexRune(string(key), s))
	// idx := -1
	// for i, r := range key {
	// 	if r == s {
	// 		idx = i
	// 		break
	// 	}
	// }
	// --------------------------------------------
	if idx < 0 {
		panic("idx < 0")
	}

	// --------------------------------------------
	idx = (idx + delta) % int32(len(key))
	// for i := 0; i < delta; i++ {
	// 	idx++
	// 	if idx >= len(key) {
	// 		idx = 0
	// 	}
	// }
	// --------------------------------------------

	return key[idx]
}

func caesarCipher(s string, k int32) string {
	// Write your code here
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := strings.ToUpper(alphabetLower)

	ret := ""
	for _, ch := range s {
		switch {
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret = ret + string(rotate(ch, k, []rune(alphabetLower)))
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret = ret + string(rotate(ch, k, []rune(alphabetUpper)))
		default:
			ret = ret + string(ch)
		}
	}
	return ret
}
