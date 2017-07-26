package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

var Navigation = MenuList{
	Name:        "",
	Description: "",
	Choices: []Choice{
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
	},
	Tabs: []Choice{
		Choice{
			Hotkey: Red("E"),
			Item:   "xit",
			Action: func() {
				fmt.Println("You pressed E")
				Continue()
			},
		},
		Choice{
			Hotkey: Red("S"),
			Item:   "pells",
			Action: func() {
				fmt.Println("You pressed S")
				Continue()
			},
		},
		Choice{
			Hotkey: Red("W"),
			Item:   "eapon",
			Action: func() {
				fmt.Println("You pressed W")
				Continue()
			},
		},
	},
}

func main() {
	AlphaInput = make(chan string)
	IntInput = make(chan int)
	go KeyReader()
	Navigation.Choose()
}
