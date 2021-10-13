// Golang program to count the number of
// repeating words in given Golang String
package files

import (
	"strings"
)

func Repetition(st string) map[string]int {

	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}
