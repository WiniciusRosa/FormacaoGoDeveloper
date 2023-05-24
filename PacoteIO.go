package main

import (
	"io"
	"log"
	"os"
)

func main() {

	if _, err := io.WriteString(os.Stdout, "Olá Winicius, bem vindo novamente!"); err != nil {
		log.Fatal(err)
	}

}
