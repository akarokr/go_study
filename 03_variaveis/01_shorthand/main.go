package main

import "fmt"

/*
	%v - Mostra o valor no formato padrão.
	%T - Mostra o tipo da variável.
*/

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
