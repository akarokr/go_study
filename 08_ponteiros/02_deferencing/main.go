package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0xc82000a340

	var b *int = &a

	fmt.Println(b)  // 0xc82000a340
	fmt.Println(*b) // 43
}
