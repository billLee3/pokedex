package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s... \n", pokemonName)

	const threshold = 50
	randomNumber := rand.Intn(pokemon.BaseExperience)
	if randomNumber <= threshold {
		fmt.Printf("%s was caught \n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		// pokedex := &cfg.pokeapiClient.Pokedex
		// pokedex.Pokedex[pokemon.Name] = pokemon
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped\n", pokemon.Name)
	}

	return nil
}
