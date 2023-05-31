package main

import "fmt"

func main() {
	s := []string{"c", "b", "a", "c", "b", "d"}
	fmt.Println(MostPopularWord(s))
}

func MostPopularWord(words []string) string {
	wordsCount := make(map[string]int, 0)
	mostPopWord := ""
	max := 0

	for _, word := range words {
		wordsCount[word]++
		if wordsCount[word] > max {
			max = wordsCount[word]
			mostPopWord = word
		}
	}

	return mostPopWord
}
