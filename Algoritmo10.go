package main

import (
	"fmt"
	"math"
)

func isPrimo(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func PreencherPrimos(slice *[]int, tamanho int) {
	if tamanho < 0 {
		return
	}

	count := 0
	num := 2

	for count < tamanho {
		if isPrimo(num) {
			*slice = append(*slice, num)
			count++
		}

		num++
	}
}

func main() {
	var primos []int
	tamanho := 10

	PreencherPrimos(&primos, tamanho)

	fmt.Println(primos)
}
