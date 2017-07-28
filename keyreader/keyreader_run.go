package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

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
	Name:   "Grok",
	Str:    Stat{"Strength", 10, 10},
	Int:    Stat{"Intelligence", 15, 15},
	Dex:    Stat{"Dexterity", 15, 15},
	Health: Stat{"Health", 10, 10},
	Dead:   false,
	Nl:     "\r",
}

var MainTab = ChoiceSet{
	Choice{
		Hotkey: Red("A"),
		Item:   "ttack",
		Action: func() {
			fmt.Println("You pressed A")
			Continue()
		},
	},
	Choice{
		Hotkey: Red("B"),
		Item:   "ackpack",
		Action: func() {
			fmt.Println("You pressed C")
			Continue()
		},
	},
	Choice{
		Hotkey: Red("C"),
		Item:   "haracter",
		Action: func() {
			fmt.Println("You pressed C")
			Continue()
		},
	},
	Choice{
		Hotkey: Red("M"),
		Item:   "ap",
		Action: func() {
			fmt.Println("You pressed W")
			Continue()
		},
	},
	Choice{
		Hotkey: Red("S"),
		Item:   "pells",
		Action: func() {
			fmt.Println("You pressed W")
			Continue()
		},
	},
	Choice{
		Hotkey: Red("E"),
		Item:   "xit",
		Action: func() {
			fmt.Println("You pressed W")
			Continue()
		},
	},
}
var MainList = ChoiceSet{
	Choice{
		//Hotkey: Red("E"),
		Item: "Self Condition",
		Action: func() {
			SelfCondition.Choose()
		},
	},
	Choice{
		//Hotkey: Red("S"),
		Item: "Banana",
		Action: func() {
			fmt.Println("You pressed Banana")
			Continue()
		},
	},
	Choice{
		//Hotkey: Red("W"),
		Item: "Bacon",
		Action: func() {
			fmt.Println("You pressed Bacon")
			Continue()
		},
	},
}

var Navigation = MenuList{
	Name:        "",
	Description: "",
	// Menu Bar
	Tabs: MainTab,
	// List Choices
	Choices: MainList,
	Header:  func() { Hero.Template() },
}

func main() {

	AlphaInput = make(chan string)
	IntInput = make(chan int)
	go KeyReader()
	Navigation.Choose()
}
