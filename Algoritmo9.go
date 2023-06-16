package main

import "fmt"

type Livro struct {
	Título string
	Autor  string
	Preço  float64
}

func AplicarDesconto(l *Livro) {
	desconto := 0.10 * l.Preço
	l.Preço -= desconto
}

func main() {
	livro := Livro{
		Título: "Livro 1",
		Autor:  "Autor 1",
		Preço:  100.0,
	}

	fmt.Println("Preço original:", livro.Preço)
	AplicarDesconto(&livro)
	fmt.Println("Preço com desconto:", livro.Preço)
}
