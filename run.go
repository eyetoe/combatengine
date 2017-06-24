package main

import (
	"fmt"
)

var Monster = Agent{
	Name:   "Monster",
	Str:    Stat{"Strength", 10, 10},
	Int:    Stat{"Intelligence", 10, 10},
	Dex:    Stat{"Dexterity", 10, 10},
	Health: Stat{"Health", 10, 10},
	Dead:   false,
}

var Hero = Agent{
	Name:   "Hero",
	Str:    Stat{"Strength", 10, 10},
	Int:    Stat{"Intelligence", 10, 10},
	Dex:    Stat{"Dexterity", 10, 10},
	Health: Stat{"Health", 10, 5},
	Dead:   false,
}

// yo dummy! add interfaces to all items, so they support conditions method which applies conditions

func main() {

	newCond := Regeneration
	newCond.Apply()
	//Hero.Conditions = append(Hero.Conditions, Regeneration)
	Hero.Conditions = append(Hero.Conditions, newCond)

	newCond2 := OnFire
	newCond2.Apply()
	Hero.Conditions = append(Hero.Conditions, newCond2)

	Hero.Print()
	fmt.Println()
	Monster.Print()
	fmt.Println()
	Hero.Attack(Monster)
	Monster.Attack(Hero)
	fmt.Println()

	for l := 0; l < 3; l++ {
		Hero.NewTurn()
		fmt.Println("Hero Health: ", Hero.Health.Val)
	}

	fmt.Println("Die Roll:", Roll(1, 20))
	Hero.Print()

}
