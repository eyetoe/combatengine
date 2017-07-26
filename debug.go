package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/util"
)

// yo dummy! add interfaces to all items, so they support conditions method which applies conditions

func Debug() {

	//var s = Screen{Name: "screen"}
	var s = Screen{120, 50}

	s.Resize()
	s.Layout()
	Continue()
	PrintPhases()
	Continue()
	ClearScreen()
	fmt.Println("testing world creation")
	ThisWorld := NewWorld()
	//fmt.Printf("%+v", ThisWorld)
	fmt.Printf("%+v", ThisWorld[5][8])

	Continue()
	ClearScreen()

	var Monster = Agent{
		Name:   "Monster",
		Str:    Stat{"Strength", 10, 10},
		Int:    Stat{"Intelligence", 10, 10},
		Dex:    Stat{"Dexterity", 10, 10},
		Health: Stat{"Health", 10, 10},
		Player: true,
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

	// Initiative Testing
	fmt.Println("\nInitiative roll")
	var heroScore int
	var monsterScore int
	for i := 0; i < 10000; i++ {
		hinit := Hero.Initiative()
		minit := Monster.Initiative()

		if hinit >= minit {
			heroScore++
		} else {
			monsterScore++

		}
	}

	fmt.Println("Hero score:", heroScore)
	fmt.Println("Monster score:", monsterScore)
	Continue()

	// Hero and Condition testing
	Hero.Print()
	fmt.Println()
	fmt.Println(":---> Adding conditions:")
	Hero.addCondition(Regen)
	Hero.addCondition(Burning)
	Hero.addCondition(Cursed)
	Hero.addCondition(Dazed)
	Hero.addCondition(Feeble)
	Hero.addCondition(Insight)
	Hero.addCondition(Paralized)
	Hero.addCondition(Petrified)
	Continue()

	Hero.NewTurn()
	Monster.NewTurn()
	Battle(&Hero, &Monster)
	//for l := 0; l < 20; l++ {
	for {
		Hero.NewTurn()
		Continue()
	}

	//fmt.Println("Die Roll:", Roll(1, 20))

}
