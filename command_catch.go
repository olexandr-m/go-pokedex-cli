package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
	if len(cfg.args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := cfg.args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	res := rand.Intn(pokemon.BaseExperience)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
