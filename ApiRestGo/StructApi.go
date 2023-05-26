package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type response struct {
	Nome    string    `json: "name"`
	pokemon []pokemon `json: "pokemon_entries"`
}

type pokemon struct {
	numero  int            `json: "entry_number"`
	species pokemonSpecies `json: "pokemon_species"`
}

type pokemonSpecies struct {
	Nome string `json: "name"`
}

func main() {

	response, err := http.Get("http://pokeapi.com/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.name)
	fmt.Println(responseObject.Pokemon)

	for _, Pokemon := range responseObject.pokemon {

		fmt.Println(Pokemon.especie.name)

	}
}
