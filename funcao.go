// Função também é chamada de procedimento ou sub-rotina
//parte código
//ela pega um dado de entrada e dá uma dado de saída

package main

import "fmt"

func media(lista []float64) float64 {
	total := 0.0

	for _, valor := range lista {
		total += valor

	}
	return total / float64(len(lista)) // interronpe a função para mandar para baixo
}

func main() { // Programa de Calculo de média

	lista := []float64{3.60, 8, 10, 9.80} // Lista de notas
	fmt.Println(media(lista))

	// ------------------------------------------------------------------------------//
	// lista := []float64{3.60, 8, 10, 9.80} // Lista de notas
	// total := 0.0
	//
	//	for _, valor := range lista {
	//		total += valor
	//	}
	//
	// fmt.Println(total / float64(len(lista))) //Imprimir média da sala
}
