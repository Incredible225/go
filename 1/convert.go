package main

import (
	"fmt"
)

var lowNames = []string{"ноль", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять", "десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}

var tensNames = []string{"двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}

var handNames = []string{"сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}

var bigNames = []string{"одна тысяча", "две тысячи", "три тысячи", "четыре тысячи", "пять тысячь", "шесть тысячь", "семь тысячь", "восемь тысячь", "девять тысячь"}

func main() {
	var a int
	fmt.Scanf("%d", &a)
	//fmt.Println(a % 100)
	if a < 20 {
		switch a {
		case 1:
			fmt.Println(lowNames[a], "Доллар")
		case 2, 3, 4:
			fmt.Println(lowNames[a], "Доллара")
		case 0, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
			fmt.Println(lowNames[a], "Долларов")
		}
	} else if a > 19 && a < 100 {

		switch a % 10 {
		case 0:
			fmt.Println(tensNames[(a/10)-2], "Долларов")
		case 1:
			fmt.Println(tensNames[(a/10)-2], lowNames[a%10], "Доллар")
		case 2, 3, 4:
			fmt.Println(tensNames[(a/10)-2], lowNames[a%10], "Доллара")
		case 5, 6, 7, 8, 9:
			fmt.Println(tensNames[a/10-2], lowNames[a%10], "Долларов")
			fmt.Println(a % 100)
		}
	} else if a > 99 && a < 1000 {
		if (a%100) < 20 && (a%100) > 9 {
			fmt.Println(handNames[a/100-1], lowNames[a%100], "Долларов")
		} else {
			switch a % 10 {
			case 0:
				fmt.Println(handNames[a/100-1], tensNames[(a%100/10)-2], "Долларов")
			case 1:
				fmt.Println(handNames[a/100-1], tensNames[(a%100/10)-2], lowNames[a%10], "Доллар")
			case 2, 3, 4:
				fmt.Println(handNames[a/100-1], tensNames[(a%100/10)-2], lowNames[a%10], "Доллара")
			case 5, 6, 7, 8, 9:
				fmt.Println(handNames[a/100-1], tensNames[(a%100/10)-2], lowNames[a%10], "Долларов")
			}
		}
	} else {
		if (a%100) < 20 && (a%100) > 9 {
			fmt.Println(bigNames[a/1000-1], handNames[(a/100)%10-1], lowNames[a%100], "Долларов")
		} else {
			switch a % 10 {
			case 1:
				fmt.Println(bigNames[a/1000-1], handNames[(a/100)%10-1], tensNames[(a%100/10)-2], lowNames[a%10], "Доллар")
			case 2, 3, 4:
				fmt.Println(bigNames[a/1000-1], handNames[(a/100)%10-1], tensNames[(a%100/10)-2], lowNames[a%10], "Доллара")
			case 5, 6, 7, 8, 9:
				fmt.Println(bigNames[a/1000-1], handNames[(a/100)%10-1], tensNames[(a%100/10)-2], lowNames[a%10], "Долларов")
			}
		}
	}

	/*
		switch a % 10 {
		case 1:
			fmt.Println(a, "Доллар")
		case 2, 3, 4:
			fmt.Println(a, "Доллара")
		case 5, 6, 7, 8, 9, 0:
			fmt.Println(a, "Долларов")
		}
	*/
}
