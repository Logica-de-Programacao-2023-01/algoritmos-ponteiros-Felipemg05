package main

import "fmt"

func ModificarValor(p *int) {
	*p = 42
}

func main() {
	var numero int
	fmt.Println("Valor inicial:", numero)

	ModificarValor(&numero)
	fmt.Println("Valor atualizado:", numero)
}
