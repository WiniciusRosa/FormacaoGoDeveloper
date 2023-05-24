// Pacote é Caixas de Funções
// Função contains
// Contains: procura dentro de outra string menor
// exemplo : Terra - rr (true)
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("terra", "rr"))
	fmt.Println(strings.Contains("computador", "dor"))
	fmt.Println(strings.Contains("pc", "computador"))
	fmt.Println(strings.Contains("Celular", "vicio"))

}
