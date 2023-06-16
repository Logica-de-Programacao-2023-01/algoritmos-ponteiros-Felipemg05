package main

import "fmt"

func AtualizarValor(p *int, n int) {
	soma := 0
	for i := 1; i <= n; i++ {
		soma += i
	}
	*p = soma
}

func main() {
	var valor int
	AtualizarValor(&valor, 5)
	fmt.Println("Valor atualizado:", valor)
}
