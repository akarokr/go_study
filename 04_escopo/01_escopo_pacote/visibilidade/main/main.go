package main

import "fmt"
import "go_study/04_escopo/01_escopo_pacote/visibilidade/vis"

func main() {
	fmt.Println(vis.MeuNome)
	vis.PrintVar()
}
