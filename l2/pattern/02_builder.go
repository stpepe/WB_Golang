package pettern

import (

)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Character struct{
	name string
}

func CreateCharacter(name string)*Character {
	return &Character{
		name: name,
	}
}

func (c *Character) Move(){
}

func (c *Character) Jump(){
}

type Sword struct{
	damage int
	speed float32
}

func CreateSword(damage int, speed float32) *Sword {
	return &Sword{
		damage: damage,
		speed: speed,
	}
}

func (s *Sword) StabAttack() {
}

type Bow struct{
	arrows int
	damage int
	speed float32
}

func CreateBow(arrows, damage int, speed float32) *Bow {
	return &Bow{
		arrows: arrows,
		damage: damage,
		speed: speed,
	}
}

func (b *Bow) ArcheryShot(){
}


type Warrior struct{
	*Character
	*Sword
}

type Archer struct{
	*Character
	*Bow
}

type Builder struct{
}

func (builder *Builder) CreateArcher(name string) *Archer {
	return &Archer{
		Character: CreateCharacter(name),
		Bow: CreateBow(64, 10, 0.5),
	}
}

func (builder *Builder) CreateWarrior(name string) *Warrior {
	return &Warrior{
		Character: CreateCharacter(name),
		Sword: CreateSword(15, 1),
	}
}

func main()  {
	
}