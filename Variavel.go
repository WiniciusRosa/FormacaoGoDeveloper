package main

import "fmt"

func main() {
	var nome string = "Winicius"  // No GO não preciso passar o tipo da Variavel, String é desnessesario
	var idade int = 22
	var versao float32 = 2.2

	fmt.Println("Olá Sr(a).", nome, "sua idade é", idade)
	fmt.Println("Esse programa está na versão", versao)
}
