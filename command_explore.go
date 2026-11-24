package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := cfg.args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
