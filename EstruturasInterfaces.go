package main

import "fmt"

type pessoa struct {
	Nome  string
	idade int
}

func main() {

	fmt.Println(pessoa{"Winicius", 22})
	fmt.Println(pessoa{"Thiago", 27})
	fmt.Println(pessoa{"Fernando", 33})
}

// São coleção de "Campos"
// Agrupar dados
// Formar Registros
// Type (  ) struct
