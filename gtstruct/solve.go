package gtstruct

import "fmt"

type Pokemon struct {
	Name      string
	HP        float64
	Attack    float64
	Defense   float64
	Speed     float64
	CountTrng int
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
