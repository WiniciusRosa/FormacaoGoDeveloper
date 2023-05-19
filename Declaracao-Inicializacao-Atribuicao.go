
// declarar meu pacote principal
package main

// Importa a Função fmt
import "fmt"

//Declaração da variável Const do ponto de ebulição da água em F
 Const ebulicaoF float64 = 212.0

// Função principal
   func main() {
 
	var tempF float64 = ebulicaoF // Variável do valor temperatura em graus F
	//	Var tempC float64 = (tempF - 32) * 5 / 9  // Conversão de F para C

	//Apareça na tela o resultado
	fmt.Println("A temperatua de ebulição da água em ºF é de ", tempF)
        fmt.Println("A temperatua de ebulição da água em ºC é de ", tempC)

	// Espero que apareça a temperatura
}
