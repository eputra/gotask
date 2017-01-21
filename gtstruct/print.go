package gtstruct

import "fmt"

func Print() {
	pikachu := Pokemon{
		Name:    "Pikachu",
		HP:      35,
		Attack:  55,
		Defense: 40,
		Speed:   90,
	}

	fmt.Printf("%s Default\n", pikachu.Name)
	pikachu.Ability()
	fmt.Printf("Attack\t: %f\n", pikachu.Attack)

	for i := 0; i < 5; i++ {
		fmt.Println()
		if i%2 == 0 {
			pikachu.Training()
			fmt.Printf("%s %d Training\n", pikachu.Name, pikachu.CountTrng)
			pikachu.Ability()
			fmt.Printf("Attack\t: %f\n", pikachu.Attack)
		} else {
			fmt.Printf("%s Stone Attack\n", pikachu.Name)
			pikachu.Ability()
			fmt.Printf("Attack\t: %f\n", pikachu.StoneAttack())
		}
	}
}
