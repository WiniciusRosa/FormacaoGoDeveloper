package main

import (
	"io/ioutil"
	"log"
)

func main() {

	message := []byte("Teste Go Winicius")
	err := ioutil.WriteFile("Testando", message, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
