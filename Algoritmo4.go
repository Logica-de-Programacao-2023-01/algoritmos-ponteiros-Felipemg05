package main

import "fmt"

func SomaUltimosDigitos(p *int) {
	numero := *p
	digito1 := numero % 10
	digito2 := (numero / 10) % 10
	soma := digito1 + digito2
	*p = soma
}

func main() {
	var valor int
	fmt.Print("Digite um valor inteiro: ")
	fmt.Scan(&valor)

	SomaUltimosDigitos(&valor)
	fmt.Println("Valor atualizado:", valor)
}
