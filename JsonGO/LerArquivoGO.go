package main

import (
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.Open("usuarios.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Arquivo aberto com Sucesso")
	defer jsonFile.Close()

}
