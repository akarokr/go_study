package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x) //endereço de x
	fmt.Println(&x)        // endereço de x na func
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // endereço no main
	fmt.Println(&x)        // endereço no main
	zero(x)
	fmt.Println(x) // x continua 5
}
