package main

import "fmt"

func VerificarParImpar(p *int) {
	if *p%2 == 0 {
		*p = 0
	} else {
		*p = 1
	}
}

func main() {
	var valor int
	fmt.Print("Digite um valor inteiro: ")
	fmt.Scan(&valor)

	VerificarParImpar(&valor)
	fmt.Println("Valor atualizado:", valor)
}
