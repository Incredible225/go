package main

import "fmt"

func main() {
	name := "youechka"
	age := 21
	money := 10.0000025
	fmt.Println(generateSelfStory(name, age, money))

}

// BEGIN (write your solution here)
func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%0.2f in my wallet right now.", name, age, money)
}
