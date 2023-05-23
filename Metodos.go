package main

import "fmt"

type retangulo struct { // Este método posssui um _tipo retangulo
	comprimento, altura int
}

func (r *retangulo) area() int {
	return r.comprimento * r.altura

}

func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura

}

//Método podem ser defenidos por qualquer tipo de receptor
//Ponteiro ou Valor. Aqui temos um Exmplo de Receptor de Valor

func main() {

	r := retangulo{comprimento: 10, altura: 5}

	fmt.Println("Area: ", r.area())
	fmt.Println("Perímetro: ", r.perimetro())
}

// método é uma função associada a um tipo particular
// Isto é, em POO(Prog, Orientada, a Objetos), objetos é um valor (Váriavel) e o mnétodo
// é uma função associada a esse objeto
