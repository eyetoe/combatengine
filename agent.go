package main

import (
	"fmt"
	"os"
)

type Agent struct {
	Name string
	Str  Stat
	Int  Stat
	Dex  Stat

	Health Stat
	Dead   bool
	Focus  int

	Conditions map[string]Condition
}

// map maker method
func (d *Agent) addCondition(c Condition) {
	//if len(d.Conditions) == 0 { // also works
	if d.Conditions == nil {
		d.Conditions = make(map[string]Condition)
		d.Conditions[c.Name] = c
	} else {
		d.Conditions[c.Name] = c
	}
	return
}

func (d *Agent) removeCondition(c Condition) {
	delete(d.Conditions, c.Name)
	return
}

func (a Agent) Attack(d Agent) *Agent {
	a.Focus = a.Focus + a.MaxFocus()
	if a.Focus > a.MaxFocus() {
		a.Focus = a.MaxFocus()
	}
	fmt.Println(a.Name + " attacks " + d.Name)
	return &a
}

func (a *Agent) Tick(c Condition) {
	if c.Duration == 1 {
		a.removeCondition(c)
	} else if c.Duration != 0 {
		var n Condition = c
		n.Duration -= 1
		a.Conditions[n.Name] = n
	}
}

func (a *Agent) HealthAdjust(h int) {
	switch {
	case a.Health.Val+h > a.Health.BaseVal:
		a.Health.Val = a.Health.BaseVal
	case a.Health.Val+h < 1:
		a.Health.Val = 0
		a.Dead = true
	default:
		a.Health.Val = a.Health.Val + h
	}
	return

}

func (a Agent) MaxFocus() int {
	return a.Int.Val / 10
}

func (a Agent) ShowConditions() {
	fmt.Println("<---: Active Conditions:")
	for l, c := range a.Conditions {
		fmt.Println("<---:", l, "	(Duration:", c.Duration, ")")
	}

}
func (a *Agent) Rebase() {
	a.Str.Val = a.Str.BaseVal
	a.Int.Val = a.Int.BaseVal
	a.Dex.Val = a.Dex.BaseVal
}

func (a *Agent) Pulse() {
	if a.Health.Val <= 0 {
		a.Dead = true
		fmt.Println("Alas,", a.Name, "has fallen!")
		os.Exit(0)
		return
	} else {
		fmt.Println(a.Name, "persists!")
		return
	}
}

func (a *Agent) NewTurn() {
	fmt.Println("<-------------------------------------------: Starting new turn")
	a.Rebase()
	for l, c := range a.Conditions {
		fmt.Println("  <: ", l, " ticked...")
		c.Affect(a)
		a.Tick(c)
	}
	a.Print()
	a.Pulse()
	return
}

func (a Agent) Print() {
	//fmt.Printf("%+v", a)
	fmt.Println("Name:", a.Name)
	fmt.Println("<:", a.Str.Val, a.Str.Name, "		+", a.Str.Modifier())
	fmt.Println("<:", a.Int.Val, a.Int.Name, "	+", a.Int.Modifier())
	fmt.Println("<:", a.Dex.Val, a.Dex.Name, "	+", a.Dex.Modifier())
	fmt.Println("<:", a.Health.Val, a.Health.Name)
	a.ShowConditions()
	//fmt.Println("Focus Points: ", a.Focus)
	//	fmt.Println(a.Conditions)
	//	fmt.Println()
	//	fmt.Println(a.Str.Name+" Modifier for: "+a.Name+" =", a.Str.Modifier())
	//	fmt.Println(a.Int.Name+" Modifier for: "+a.Name+" =", a.Int.Modifier())
	//	fmt.Println(a.Dex.Name+" Modifier for: "+a.Name+" =", a.Dex.Modifier())
	//	fmt.Println("Max Focus for", a.Name, "=", a.MaxFocus())
}
