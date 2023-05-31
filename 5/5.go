package main

import (
	"fmt"
	"strings"
)

func Greetings(name string, age int) string {

	name = strings.Trim(name, " ")
	name = strings.ToLower(name)
	name = strings.Title(name)
	return fmt.Sprintf("Привет, %s! %d", name, age)

}
func main() {
	fmt.Println(Greetings("аЛ", 21))
}
