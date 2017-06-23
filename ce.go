package main

import (
	"fmt"
)

type Stat struct {
	Name    string
	BaseVal int
	Val     int
}

type Condition struct {
	Name         string
	BaseDuration int
	Duration     int
	Affect       func(Agent) Agent
}

type Agent struct {
	Name string
	Str  Stat
	Int  Stat
	Dex  Stat

	Health Stat
	Dead   bool
	Focus  int

	Conditions []Condition
}

var Regeneration = Condition{
	Name:         "Natural Health Recovery",
	BaseDuration: 0,
	Affect: func(a Agent) Agent {
		//if a.Health.Val < a.Health.BaseVal {
		a.Health.Val = a.Health.Val + 1
		//	}
		return a
	},
}

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

func (a Agent) Attack(d Agent) *Agent {
	a.Focus = a.Focus + a.MaxFocus()
	if a.Focus > a.MaxFocus() {
		a.Focus = a.MaxFocus()
	}
	fmt.Println(a.Name + " attacks " + d.Name)
	return &a
}

func (a *Agent) NewTurn() {
	fmt.Println("Starting new turn")
	for index, cond := range a.Conditions {
		fmt.Println(index, cond.Name)
		*a = cond.Affect(*a)
	}
	return
}

func (a Agent) MaxFocus() int {
	return a.Int.Val / 10
}

func (a Agent) Print() {
	//fmt.Printf("%+v", a, "\n")
	fmt.Printf("%+v", a)
	fmt.Println()
	fmt.Println(a.Str.Name+" Modifier for: "+a.Name+" =", a.Str.Modifier())
	fmt.Println(a.Int.Name+" Modifier for: "+a.Name+" =", a.Int.Modifier())
	fmt.Println(a.Dex.Name+" Modifier for: "+a.Name+" =", a.Dex.Modifier())
	fmt.Println("Max Focus for", a.Name, "=", a.MaxFocus())
	fmt.Println("Current Focus for", a.Name, "=", a.Focus)
}

func (s Stat) Modifier() int {
	return s.Val / 5
}

func main() {

	Hero.Conditions = append(Hero.Conditions, Regeneration)
	Hero.Print()
	Monster.Print()
	Hero.Attack(Monster)
	Monster.Attack(Hero)
	fmt.Println()
	fmt.Println("Hero Health: ", Hero.Health.Val)
	Hero.NewTurn()
	Hero.NewTurn()
	Hero.NewTurn()
	fmt.Println("Hero Health: ", Hero.Health.Val)

}
