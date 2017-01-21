package gtslice

import (
	"fmt"

	"github.com/eputra/gotask/gtstruct"
)

type Bag struct {
	Pokemon []gtstruct.Pokemon
}

func (b *Bag) AddPokemon(pokemon gtstruct.Pokemon) {
	sama := 0

	for _, pokeBag := range b.Pokemon {
		if pokeBag.Name == pokemon.Name {
			sama = 1
		}
	}

	if sama == 0 {
		b.Pokemon = append(b.Pokemon, pokemon)
	}
}

func (b Bag) Contents() {
	fmt.Println(len(b.Pokemon))
	fmt.Println(b.Pokemon)
}
