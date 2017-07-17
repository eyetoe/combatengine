package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

var Console = MenuList{
	Name:        "",
	Description: "",
	Header: func() {
		Hero.StatusBar()
		Hero.Print()
	},
	Choices: []Choice{
		Choice{
			Item: "Test Choices",
			Action: func() {
				TestChoices()
			},
		},
		Choice{
			Item: SelfCondition.Name,
			Action: func() {
				SelfCondition.Choose()
			},
		},
		Choice{
			Item: "Debug",
			Action: func() {
				Debug()
			},
		},
		Choice{
			Item: "Take Action",
			Action: func() {
				Action.Choose()
			},
		},
	},
}

var Action = MenuList{
	Name: "Take Action",
	Header: func() {
		Hero.StatusBar()
		Hero.Print()
	},
	Choices: []Choice{
		Choice{
			Item: Red("Cast Spell"),
			Action: func() {
				Spell()
			},
		},
		/*
			Choice{
				Item: Blue("Console"),
				Action: func() {
					Console.Choose()
				},
			},
		*/
	},
}

var SelfCondition = MenuList{
	Name:        "Self Conditions",
	Description: "Apply Condition to Self",
	Header: func() {
		Hero.StatusBar()
		Hero.Print()
	},
	Choices: []Choice{
		Choice{
			Item: Regen.Color("Regen"),
			Action: func() {
				Hero.addCondition(Regen)
			},
		},
		Choice{
			Item: Burning.Color("Burning"),
			Action: func() {
				Hero.addCondition(Burning)
			},
		},
		Choice{
			Item: Cursed.Color("Cursed"),
			Action: func() {
				Hero.addCondition(Cursed)
			},
		},
		Choice{
			Item: Dazed.Color("Dazed"),
			Action: func() {
				Hero.addCondition(Dazed)
			},
		},
		Choice{
			Item: Feeble.Color("Feeble"),
			Action: func() {
				Hero.addCondition(Feeble)
			},
		},
		Choice{
			Item: Insight.Color("Insight"),
			Action: func() {
				Hero.addCondition(Insight)
			},
		},
		Choice{
			Item: Paralized.Color("Paralized"),
			Action: func() {
				Hero.addCondition(Paralized)
			},
		},
		/*
			Choice{
				Item: Blue("Console"),
				Action: func() {
					Console.Choose()
				},
			},
		*/
	},
}

func Spell() {
	fmt.Println("placeholder for spell stuff")
	Continue()
}
