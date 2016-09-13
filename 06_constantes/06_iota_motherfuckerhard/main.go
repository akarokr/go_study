package main

import "fmt"

const (
	_  = iota             // 0 essa variavel serve apenas para "iniciar" a contagem do iota
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
)

func main() {
	fmt.Println("binary\t\tdecima")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
}
