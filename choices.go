package main

import (
	"fmt"
	"strconv"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

//Choice, is something you may choose from a MenuList
type Choice struct {
	Item string
	//Action func(...interface{})
	Action func()
}

type MenuList struct {
	Name        string
	Description string
	Choices     []Choice
	Header      func()
}

func (m MenuList) Choose() {
	for {
		if m.Header != nil {
			ClearScreen()
			m.Header()
		}
		fmt.Println(YellowU(m.Name), "\n")
		fmt.Println(Blue(m.Description))
		for i, c := range m.Choices {
			fmt.Println(CyanU(i), c.Item)
		}
		choice, _, _ := GetChar()
		fmt.Println("choice was:", choice)

		if c, err := strconv.Atoi(choice); err == nil && c < len(m.Choices) {
			fmt.Printf("%T, %v", c, c)
			fmt.Println()
			m.Choices[c].Action()
			break
		} else {
			continue
		}
	}

}

func TestChoices() {

	var Dude = Agent{
		Name:   "Monster",
		Str:    Stat{"Strength", 10, 10},
		Int:    Stat{"Intelligence", 10, 10},
		Dex:    Stat{"Dexterity", 10, 10},
		Health: Stat{"Health", 10, 10},
		Dead:   false,
	}

	var Question = MenuList{
		Name:        "Favorite Color",
		Description: "What is your favorite color?",
		Choices: []Choice{
			Choice{
				Item: "red",
				Action: func() {
					fmt.Println("You chose red!")
				},
			},
			Choice{
				Item: "blue",
				Action: func() {
					fmt.Println("You chose blue!")
				},
			},
			Choice{
				Item:   "again-again",
				Action: TestChoices,
			},
			Choice{
				Item: "Something arbitrary perhaps?",
				Action: func() {
					Hero.Print()
				},
			},
		},
	}

	var Question2 = MenuList{
		Name:        "Volcanos",
		Description: "What comes out of volcanos?",
		Choices: []Choice{
			Choice{
				Item:   "Magma",
				Action: func() { fmt.Println("Magma is what stays in the ground :(") },
			},
			Choice{
				Item:   "Lava",
				Action: func() { fmt.Println("Yes, lava is what comes out of the volcano!") },
			},
			Choice{
				Item:   "Potato juice",
				Action: func() { fmt.Println("What have you been smoking?") },
			},
			Choice{
				Item:   "Lemon juice",
				Action: func() { fmt.Println("So tart.. but alas not from a volcano.") },
			},
			Choice{
				Item:   "Frog's ball juice",
				Action: func() { fmt.Println("Frog daddy juice is important to frogs, but not in Pompeii!") },
			},
			Choice{
				Item:   "Endangerd feces",
				Action: func() { fmt.Println("A plort for all my men!  But not for my volcano!") },
			},
		},
	}

	fmt.Printf("%+v", Question)
	Continue()
	fmt.Printf("%+v", Dude)
	Continue()
	Question.Choose()
	Continue()
	Question2.Choose()
	Continue()
}

//	Describe: map[string]string{
//		"self":   "Your body regenerates.",
//		"target": "It regenerates.",
//	},

/*
		// ask the first question
		fmt.Printf("%svent, ", GreenU("E"))
		fmt.Printf("%sest, ", GreenU("R"))
		fmt.Printf("%srain ", GreenU("T"))
		fmt.Printf("%so, ", GreenU("G"))
		fmt.Printf("%soad, ", GreenU("L"))
		fmt.Printf("%spawn, ", GreenU("S"))
		fmt.Printf("%snventory <: ", GreenU("I"))

		choice, _, _ := GetChar()

		// choices here
		switch choice {
		case "e", "E":
			//var Foe Agent
			Foe = Spawn(Char)
			Fight(&Char, &Foe)
			continue
		case "r", "R":
			fmt.Printf(Blue("\nYou rest and recuperate.\n"))
			if Char.Health.Val < Char.MaxHealth.Val/2 {
				Char.Health.Val = Char.MaxHealth.Val / 2
				fmt.Printf(Green("You heal.\n"))
				Char.Save()
			}
			Continue()
			continue
		case "t", "T":
			ExpMgr(&Char)
			continue
		case "g", "G":
			Go(&Char)
			break
		case "l", "L":
			Banner()
			PickChar()
			Char.Load()
			Resurrect(&Char)
			continue
		case "s", "S":
			var Foe Agent
			Foe = *SpawnChooser(&Char)
			Fight(&Char, &Foe)
			continue
		case "i", "I":
			Inventory(&Char)
			break
		}
	}
*/
