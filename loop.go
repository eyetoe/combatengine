package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/util"
)

var Monster = Agent{
	Name:   "Monster",
	Str:    Stat{"Strength", 10, 10},
	Int:    Stat{"Intelligence", 10, 10},
	Dex:    Stat{"Dexterity", 10, 10},
	Health: Stat{"Health", 10, 10},
	AI:     true,
	Dead:   false,
}

var Hero = Agent{
	Name:   "Hero",
	Str:    Stat{"Strength", 10, 10},
	Int:    Stat{"Intelligence", 15, 15},
	Dex:    Stat{"Dexterity", 15, 15},
	Health: Stat{"Health", 10, 10},
	Dead:   false,
}

func main() {
	//	Debug()

	for {
		Console.Choose()
	}
}

func Battle(a, b *Agent) {

	var first *Agent
	var second *Agent
	// stealth here?

	// Who goes first?
	if a.Initiative() > b.Initiative() {
		fmt.Println(a.Name, "has the initiative.")
		first = a
		second = b
	} else if b.Initiative() > a.Initiative() {
		fmt.Println(b.Name, "has the initiative.")
		first = b
		second = a
	} else {
		Battle(a, b)
	}

	for {
		Attack(first, second)
		second.NewTurn()
		Attack(second, first)
		first.NewTurn()

	}
}

func Attack(a, d *Agent) {
	var attackRoll = Roll(1, 100)
	var defenseRoll = Roll(1, 100)
	if a.AI == false && a.Focus > 0 {
		if Confirm("Use focus point to attack?") {
			fmt.Println(a.Name, "attacks", d.Name)
			fmt.Println(a.Name, "rolled:", attackRoll)
			fmt.Println(d.Name, "rolled:", defenseRoll)
			a.Focus--
		}
	} else if a.AI == true && a.Focus > 0 {
		fmt.Println(a.Name, "is attacking!")
		fmt.Println(a.Name, "attacks", d.Name)
		fmt.Println(a.Name, "rolled:", attackRoll)
		if Confirm("Use focus point to defend?") {
			fmt.Println(d.Name, "rolled:", defenseRoll)
			d.Focus--
		}
	}

}
