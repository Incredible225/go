package main

import "fmt"

func isCel(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("Введите целое число:")
	fmt.Scanln(&n)

	for i := n + 1; ; i++ {
		if isCel(i) == true {
			fmt.Printf("Наименьшее простое число больше %d: %d\n", n, i)
			break
		}
	}

}
