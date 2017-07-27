package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

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
		Item: "Apple",
		Action: func() {
			fmt.Println("You pressed Apple")
			Continue()
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
}

func main() {
	AlphaInput = make(chan string)
	IntInput = make(chan int)
	go KeyReader()
	Navigation.Choose()
}
