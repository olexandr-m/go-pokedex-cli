package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		pokeapiClient: pokeClient,
		commands:      getCommands(),
	}
	startRepl(cfg)
}
