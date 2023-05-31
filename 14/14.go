package main

import "fmt"

type Person struct {
	Age uint8
}

// BEGIN (write your solution here)
type PersonList []Person

func (pl PersonList) GetAgePopularity() map[uint8]int {
	m := make(map[uint8]int)

	for _, r := range pl {
		m[r.Age]++
	}
	return m
}
func main() {
	pl := PersonList{
		{Age: 18},
		{Age: 44},
		{Age: 18},
	}

	fmt.Println(pl.GetAgePopularity())
}
