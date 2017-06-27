package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/util"
)

// yo dummy! add interfaces to all items, so they support conditions method which applies conditions

func main() {
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

	/*
		newCond := Regeneration
		newCond.Apply()
		//Hero.Conditions = append(Hero.Conditions, Regeneration)
		Hero.Conditions = append(Hero.Conditions, newCond)

		newCond2 := OnFire
		newCond2.Apply()
		Hero.Conditions = append(Hero.Conditions, newCond2)
	*/

	Hero.addCondition(Regeneration)
	Hero.addCondition(OnFire)
	//	for l, c := range Hero.Conditions {
	//		fmt.Println("Condition:", l, "\n	Turns Left:", c.Duration, " :HP:", Hero.Health.Val)
	//	}
	/*
		Hero.Print()
		fmt.Println()
		Monster.Print()
		fmt.Println()
		Hero.Attack(Monster)
		Monster.Attack(Hero)
		fmt.Println()
	*/

	for l := 0; l < 20; l++ {
		Hero.NewTurn()
		Hero.Print()
		//fmt.Println("Hero Health: ", Hero.Health.Val)
		for l, c := range Hero.Conditions {
			//	fmt.Println("Condition:", l, "\n	Turns Left:", c.Duration)
			fmt.Println("Condition:", l, "\n	Turns Left:", c.Duration, " :HP:", Hero.Health.Val)
		}
		fmt.Println()
		Continue()
	}

	//fmt.Println("Die Roll:", Roll(1, 20))
	//Hero.Print()

}
