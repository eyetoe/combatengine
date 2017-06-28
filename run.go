package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/util"
)

// yo dummy! add interfaces to all items, so they support conditions method which applies conditions

func main() {
	ClearScreen()

	/*
		var Monster = Agent{
			Name:   "Monster",
			Str:    Stat{"Strength", 10, 10},
			Int:    Stat{"Intelligence", 10, 10},
			Dex:    Stat{"Dexterity", 10, 10},
			Health: Stat{"Health", 10, 10},
			Dead:   false,
		}
	*/
	var Hero = Agent{
		Name:   "Hero",
		Str:    Stat{"Strength", 10, 10},
		Int:    Stat{"Intelligence", 10, 10},
		Dex:    Stat{"Dexterity", 10, 10},
		Health: Stat{"Health", 10, 10},
		Dead:   false,
	}

	Hero.Print()
	fmt.Println()
	fmt.Println(":---> Adding conditions:")
	Hero.addCondition(Regeneration)
	Hero.addCondition(OnFire)
	Continue()

	for l := 0; l < 20; l++ {
		Hero.NewTurn()
		Continue()
	}

	//fmt.Println("Die Roll:", Roll(1, 20))

}
