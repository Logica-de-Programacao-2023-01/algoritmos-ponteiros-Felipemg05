package main

import (
	"fmt"
	"math"
)

func AtualizarMedia(p *float64) {
	valor := *p
	media := (valor + math.Pi) / 2
	*p = media
}

func main() {
	var numero float64
	fmt.Print("Digite um valor: ")
	fmt.Scan(&numero)

	AtualizarMedia(&numero)
	fmt.Println("Valor atualizado:", numero)
}
