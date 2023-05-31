package main

import (
	"fmt"
)

func main() {
	s := "sssss"
	s = s + "e"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}

}
