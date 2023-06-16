package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func AdicionarValor(c *Conta, valor float64) {
	c.Saldo += valor
}

func main() {
	conta := Conta{
		Saldo:   100.0,
		Titular: "João",
	}

	fmt.Println("Saldo inicial:", conta.Saldo)
	AdicionarValor(&conta, 50.0)
	fmt.Println("Saldo atualizado:", conta.Saldo)
}
