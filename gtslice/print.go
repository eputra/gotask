package gtslice

import "github.com/eputra/gotask/gtstruct"

func Print() {
	var bag Bag

	pikachu := gtstruct.Pokemon{
		Name:    "Pikachu",
		HP:      35,
		Attack:  55,
		Defense: 40,
		Speed:   90,
	}

	bulbasaur := gtstruct.Pokemon{
		Name:    "Bulbasaur",
		HP:      45,
		Attack:  49,
		Defense: 49,
		Speed:   45,
	}

	charmander := gtstruct.Pokemon{
		Name:    "Charmander",
		HP:      39,
		Attack:  52,
		Defense: 43,
		Speed:   65,
	}

	bag.AddPokemon(pikachu)
	bag.AddPokemon(bulbasaur)
	bag.AddPokemon(charmander)
	bag.Contents()
}
