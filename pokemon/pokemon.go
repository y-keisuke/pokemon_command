package pokemon

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pokemon struct {
	Name string `json:"name"`
	HP int `json:"hp"`
	Attack int `json:"attack"`
	Defense int `json:"defense"`
	SpAttack int `json:"sp_attack"`
	SpDefense int `json:"sp_defense"`
	Speed int `json:"speed"`
}

func JsonToPokemon(pokemonJson []byte) *Pokemon {
	pokemon := new(Pokemon)
	err := json.Unmarshal(pokemonJson, pokemon)
	if err != nil {
		log.Fatal(err)
	}
	return pokemon
}

func PrintPokemon(pokemon *Pokemon) {
	fmt.Println("なまえ : ", pokemon.Name)
	fmt.Println("HP　    : ", pokemon.HP)
	fmt.Println("こうげき : ", pokemon.Attack)
	fmt.Println("ぼうぎょ : ", pokemon.Defense)
	fmt.Println("とくこう : ", pokemon.SpAttack)
	fmt.Println("とくぼう : ", pokemon.SpDefense)
	fmt.Println("すばやさ : ", pokemon.Speed)
}