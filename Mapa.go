package main

import "fmt"

func main() {

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "lítio"
	fmt.Println(elemento["H"])

	//	x := make(map[int]int)
	//	x[1] = 20
	//	x[2] = 30
	//	fmt.Println(x[1], x[2])

	//	x := make(map[string]int)
	//	x["Chave"] = 10
	//	fmt.Println(x["Chave"])

	// coleção ordenada (Lista) pares chave-valor
	// var x map[string]int
	// x é um mapa de strings para ints

}
