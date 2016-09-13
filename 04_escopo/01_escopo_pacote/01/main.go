package main

import "fmt"

/*
	Escopo a nível de pacote. É acessível por todos os métodos dentro do pacote.

*/

var x int = 42

func main() {
	fmt.Println(x)
	foo()
	y := "Bla" // Essa variável é acessada somente por esse bloco de código.
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}
