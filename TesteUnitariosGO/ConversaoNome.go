package main

import "testing"

func ShouldSumCorrect(t *testing.T) {

	teste := soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}

}

func ShouldSumIncorrect(t *testing.T) {

	teste := soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}

}

func ShouldMultCorrect(t *testing.T) {

	teste := Multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
	
	
	
   func ShouldMultIncorrect(t *testing.T) {

	teste := Multiplica(10, 10)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}

}

}
