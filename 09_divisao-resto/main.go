package main

import "fmt"

func odd_even(x int, y int) {
	m := x % y
	if m == 1 {
		fmt.Println("impar")
	} else {
		fmt.Println("par")
	}
}

func main() {
	//	x := 13 % 3
	//	fmt.Println("O resto é ", x)

	odd_even(13, 3)
}
