package main

import "fmt"

func main() {
	s := "abc"
	step := 1
	sOut := ""

	for i := 0; i < len(s); i++ {

		sOut = sOut + string(byte(int(s[i])+step))
	}
	fmt.Println(sOut)
}

/*	p := ""
	for i := 0; i < len(s); i++ {

		if (int(byte(s[i])))+step > 255 || int(byte(s[i]))+step < 0 {
			p = string(byte((int(byte(s[i]))) + step%256))
		} else {
			p = string(byte((int(byte(s[i]))) + step))
		}

		sOut += p
	}
*/
