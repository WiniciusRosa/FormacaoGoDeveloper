package main

import "testing"

func ShouldSumCorrect(t *testing.T) {
   //arrange
	teste := soma(3, 2, 1)
   //act
	resultado := 6
   //assert
	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}

}
