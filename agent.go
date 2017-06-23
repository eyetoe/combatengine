package main

import "fmt"

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
	//fmt.Printf("%+v", a)
	fmt.Println(a.Name)
	fmt.Println(a.Str)
	fmt.Println(a.Int)
	fmt.Println(a.Dex)
	fmt.Println(a.Health)
	fmt.Println("Dead?: ", a.Dead)
	fmt.Println("Focus Points: ", a.Focus)
	fmt.Println(a.Conditions)
	fmt.Println()
	fmt.Println(a.Str.Name+" Modifier for: "+a.Name+" =", a.Str.Modifier())
	fmt.Println(a.Int.Name+" Modifier for: "+a.Name+" =", a.Int.Modifier())
	fmt.Println(a.Dex.Name+" Modifier for: "+a.Name+" =", a.Dex.Modifier())
	fmt.Println("Max Focus for", a.Name, "=", a.MaxFocus())
}
