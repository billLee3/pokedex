package pokeapi

type Pokedex struct {
	Pokedex map[string]Pokemon
}

func NewPokedex() Pokedex {
	// p := Cache{
	// 	cache: make(map[string]cacheEntry),
	// 	mux:   &sync.Mutex{},
	// }
	return Pokedex{
		Pokedex: make(map[string]Pokemon),
	}
}
