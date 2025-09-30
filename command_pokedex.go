package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	for _, p := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", p.Name)
	}

	return nil
}
