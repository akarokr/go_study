package main

/*
	- Variável em caixa alta: acessível fora do pacote
	- Variável em camelCase: acessível somente dentro do pacote
*/

import (
	"fmt"
	"go_study/02_package/nomes"
)

func main() {
	fmt.Println(nomes.MeuNome)
}
