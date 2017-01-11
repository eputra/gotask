package main

import "fmt"

type Pokemon struct {
	Name      string
	HP        float64
	Attack    float64
	Defense   float64
	Speed     float64
	CountTrng int
}

func main() {
	var pikachu Pokemon

	pikachu = Pokemon{}
	pikachu.Name = "Pikachu"
	pikachu.HP = 35
	pikachu.Attack = 55
	pikachu.Defense = 40
	pikachu.Speed = 90

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

func (p Pokemon) Ability() {
	fmt.Printf("HP\t: %f\n", p.HP)
	fmt.Printf("Defense\t: %f\n", p.Defense)
	fmt.Printf("Speed\t: %f\n", p.Speed)
}

func (p *Pokemon) Training() {
	p.HP += p.HP * 0.05
	p.Attack += p.Attack * 0.05
	p.Defense += p.Defense * 0.05
	p.Speed += p.Speed * 0.05
	p.CountTrng += 1
}

func (p Pokemon) StoneAttack() float64 {
	return p.Attack + 50
}
