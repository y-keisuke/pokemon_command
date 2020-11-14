package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pokemon_command/input"
	"pokemon_command/pokemon"
)

func main() {
	url := input.CreateUrl()

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	var errCheck map[string]string
	json.Unmarshal(byteArray, &errCheck)
	if val, ok := errCheck["err"]; ok {
		fmt.Println(val)
		return
	}

	pokemonStruct := pokemon.JsonToPokemon(byteArray)

	pokemon.PrintPokemon(pokemonStruct)
}