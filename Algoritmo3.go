package main

import "fmt"

func InverterString(p *string) {
	str := *p
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	*p = string(runes)
}

func main() {
	var texto string
	fmt.Print("Digite um texto: ")
	fmt.Scan(&texto)

	InverterString(&texto)
	fmt.Println("Texto invertido:", texto)
}
