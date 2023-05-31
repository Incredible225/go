package main

import "fmt"

func main() {
	var j string
	var s string

	fmt.Scanf("%s", &j)
	fmt.Scanf("%s", &s)

	seen := map[rune]struct{}{}
	for _, letter := range s {
		seen[letter] = struct{}{}
	}

	result := 0
	for _, stone := range j {
		if _, ok := seen[stone]; ok {
			result++
		}
	}

	fmt.Println(result)
}
