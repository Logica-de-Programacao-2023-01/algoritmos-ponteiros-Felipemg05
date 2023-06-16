package main

import "fmt"

type Livro struct {
	Título string
	Autor  string
}

func AlterarTítulo(l *Livro) {
	if l.Autor == "Anônimo" {
		l.Título = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Título: "Livro 1",
		Autor:  "Anônimo",
	}

	fmt.Println("Título original:", livro.Título)
	AlterarTítulo(&livro)
	fmt.Println("Título atualizado:", livro.Título)
}
Neste exemplo, definimos a struct Livro com os campos Título e Autor. A função AlterarTítulo recebe um ponteiro para um struct Livro l. Dentro da função, verificamos se o autor do livro é "Anônimo". Se for, alteramos o título do livro para "Desconhecido" atribuindo o novo valor ao campo Título do struct.

Na função main, criamos um exemplo de livro com um título e autor definidos. Imprimimos o título original do livro e, em seguida, chamamos a função `AlterarT



