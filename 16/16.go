// Напишите программу, которая будет подсчитывать количество вхождений каждого слова в заданном тексте.
// Выведите на экран список слов и соответствующих им частот. Используйте мапу для хранения информации.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	s = "привет как дела чувак привет как дела"
	mapcount := make(map[string]int)
	slices := spliter(s)
	fmt.Println(slices)
	for _, r := range slices {
		mapcount[r]++
	}
	for k, v := range mapcount {
		fmt.Println(k, ":", v)
	}
}

func spliter(s string) []string {
	slices := strings.Split(s, " ")
	return slices
}
