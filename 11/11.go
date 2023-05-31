package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	k := "хелллоу world"

	fmt.Println(latinLetters((k)))
}

// BEGIN (write your solution here) (write your solution here)
func latinLetters(s string) string {

	m := ""
	sb := make([]string, 0)
	for _, r := range s {
		if unicode.Is(unicode.Latin, rune(r)) {

			sb = append(sb, string(r))
			m = strings.Join(sb, "")

		}

	}
	return m

	/*	sb := &strings.Builder{}

		for _, r := range s {
			if unicode.Is(unicode.Latin, r)  {
				sb.WriteRune(r)
			}
		}

		return sb.String()
	*/
}
