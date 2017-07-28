package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

////////////////////////////////////////////////////////////////////////////////
type Agent struct {
	Name string
	Str  Stat
	Int  Stat
	Dex  Stat

	Health    Stat
	Dead      bool
	Focus     int
	Player    bool
	FoeMaxHit int
	/* Alternative names for Focus
	Do stuff points DSP
	Moxie
	Energy
	Exertion
	Red Bull Points
	Concentration
	Reserves
	Multi-tasking points
	Threads
	Tasks
	Determination
	Will
	Drive
	Stamina
	Endurance
	Potential
	Motivation

	*/

	Conditions map[string]Condition
	Nl         string
}

////////////////////////////////////////////////////////////////////////////////
//func (a *Agent) HealthMeter(num, max, fmax int, l, t string) string {
func (a *Agent) HealthMeter() string {
	// Meter(a.Health.Val, a.Health.BaseVal, a.FoeMaxHit, "Health", "█")
	token := "█"

	// calculate the actual percentage
	p := float32(a.Health.Val) / float32(a.Health.BaseVal)

	barWidth := 80

	// prepare as percentage of barWidth columns
	percentBar := p * float32(barWidth)

	// display meter
	//fmt.Printf("%s	%s	:", Yellow("Health"), White(strconv.Itoa(a.Health.Val)))

	var buffer bytes.Buffer

	// drawing from left to right
	for c := 1; c <= barWidth; c++ {

		if c <= int(percentBar) {
			//fmt.Printf(Cyan(token))
			buffer.WriteString(Cyan(token))
		} else {
			if a.Health.Val <= a.FoeMaxHit {
				buffer.WriteString(Red(token))
			} else {
				buffer.WriteString(Black(token))
			}
		}
		if c == barWidth {
			//buffer.WriteString(":")
		}
	}
	return buffer.String()
}

////////////////////////////////////////////////////////////////////////////////
func (a *Agent) Template() {
	//tmpl, err := template.ParseFiles("templates/agent.template")
	tmpl, err := template.ParseGlob("templates/*.template")
	if err != nil {
		panic(err)
	}
	// repeat this block for more template filess
	err = tmpl.ExecuteTemplate(os.Stdout, "agent.template", a)
	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////
func (a *Agent) Initiative() int {
	return Roll(1, 100) + (a.Int.Val+a.Dex.Val)/2
}

////////////////////////////////////////////////////////////////////////////////
// map maker method
func (d *Agent) addCondition(c Condition) {
	//if len(d.Conditions) == 0 { // also works
	if d.Conditions == nil {
		d.Conditions = make(map[string]Condition)
		d.Conditions[c.Name] = c
	} else {
		d.Conditions[c.Name] = c
	}
	fmt.Println(c.Color(c.Name), ":", c.Describe["self"])
	fmt.Println(c.Color(c.Name), ":", c.Describe["target"])
	return
}

////////////////////////////////////////////////////////////////////////////////
func (d *Agent) removeCondition(c Condition) {
	delete(d.Conditions, c.Name)
	return
}

/*
func (a Agent) Attack(d Agent) *Agent {
	a.Focus = a.Focus + a.MaxFocus()
	if a.Focus > a.MaxFocus() {
		a.Focus = a.MaxFocus()
	}
	fmt.Println(a.Name + " attacks " + d.Name)
	return &a
}
*/

////////////////////////////////////////////////////////////////////////////////
func (a *Agent) Tick(c Condition) {
	if c.Duration == 1 {
		a.removeCondition(c)
	} else if c.Duration != 0 {
		var n Condition = c
		n.Duration -= 1
		a.Conditions[n.Name] = n
	}
}

////////////////////////////////////////////////////////////////////////////////
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

/*
func (a Agent) ShowConditions() {
	fmt.Println(CyanBU("\nActive conditions..."))
	for l, c := range a.Conditions {
		fmt.Println("	", c.Color(l), Blue(c.Duration))
	}

}
*/

////////////////////////////////////////////////////////////////////////////////
func (a Agent) ShowConditions() {
	fmt.Printf(CyanBU("\nConditions:"))
	fmt.Printf(" ")
	for l, c := range a.Conditions {
		fmt.Printf("%s", c.Color(l))
		fmt.Printf("%s", ":")
		fmt.Printf("%s", Blue(strconv.Itoa(c.Duration)))
		fmt.Printf("%s", " ")
	}

}

////////////////////////////////////////////////////////////////////////////////
func (a *Agent) Rebase() {
	a.Str.Val = a.Str.BaseVal
	a.Int.Val = a.Int.BaseVal
	a.Dex.Val = a.Dex.BaseVal
	// reset focus
	a.Focus = a.Focus + a.MaxFocus()
	if a.Focus > a.MaxFocus() {
		a.Focus = a.MaxFocus()
	}
}

////////////////////////////////////////////////////////////////////////////////
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
	fmt.Println("<--------------------------------------: Starting new turn")
	a.Rebase()
	fmt.Println(CyanBU("Appling conditions..."))
	for l, c := range a.Conditions {
		fmt.Println("	", c.Color(l), Black("ticked..."))
		c.Affect(a)
		a.Tick(c)
	}
	a.Print()
	a.Pulse()
	return
}

func (a Agent) Print() {
	//fmt.Printf("%+v", a)
	fmt.Println()
	fmt.Println("<: Name:", a.Name)
	fmt.Println("	", a.Str.Val, "	", a.Str.Name, "	+", a.Str.Modifier())
	fmt.Println("	", a.Int.Val, "	", a.Int.Name, "	+", a.Int.Modifier())
	fmt.Println("	", a.Dex.Val, "	", a.Dex.Name, "	+", a.Dex.Modifier())
	fmt.Println("	", a.Health.Val, "	", a.Health.Name)
	fmt.Println("	", a.Focus, "	", "Focus Points")
	a.ShowConditions()
}

// Render character status bar
func (c Agent) StatusBar() {
	// For sanity layout the StatusBar vertically while printing horizonal
	fmt.Println()
	fmt.Printf("%s", Fbb("FooBarBaz:"))
	fmt.Printf("%s", Yellow(" "))
	fmt.Printf("%s", BlueU(c.Name))
	fmt.Printf("%s", Yellow(" S:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Str.Val)))
	fmt.Printf("%s", Yellow(" I:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Int.Val)))
	fmt.Printf("%s", Yellow(" D:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Dex.Val)))
	fmt.Printf("%s", Yellow(" HP:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Health.BaseVal)))
	fmt.Printf("%s", Yellow("|"))

	if c.Health.Val < c.Health.BaseVal {
		fmt.Printf("%s", Red(strconv.Itoa(c.Health.Val)))
	} else {
		fmt.Printf("%s", Green(strconv.Itoa(c.Health.Val)))
	}
	c.ShowConditions()

	/*
		fmt.Printf("%s", Yellow(" XP:"))
		if StatCost(c) <= c.Exp {
			fmt.Printf("%s", Green(strconv.Itoa(c.Exp)))
		} else {
			fmt.Printf("%s", strconv.Itoa(c.Exp))
		}

		fmt.Printf("%s", Yellow(" W:"))
		fmt.Printf("%s", ItemC(c.Weap.Name))
		fmt.Printf("%s", Yellow(" A:"))
		fmt.Printf("%s", ItemC(c.Armor.Name))
		fmt.Printf("%s", Yellow(" T:"))
		fmt.Printf("%s", ItemC(c.Trink.Name))
	*/
	fmt.Println()
	Meter(c.Health.Val, c.Health.BaseVal, c.FoeMaxHit, "Health", "█", "hero")
	// interesting palate of ascii for perusing
	//░▒█░   ░ ████▓▒░░ ████▓▒░░▓█  ▀█▓ ▓█   ▓██▒░██▓ ▒██▒░▓█  ▀█▓ ▓█   ▓██▒███████▒
}
func FoeBar(c Agent, f Agent) {
	// this may be useful it clears the screen
	//fmt.Print("[H[J")

	// Color Foe's name and show victory chance percentage
	/*
		odds := Odds(&c, &f)
		var conName string
		switch {
		case odds >= 80:
			conName = fmt.Sprintf(GreenU("%s:%d%%"), f.Name, odds)
		case odds >= 60:
			conName = fmt.Sprintf(CyanU("%s:%d%%"), f.Name, odds)
		case odds >= 40:
			conName = fmt.Sprintf(BlueU("%s:%d%%"), f.Name, odds)
		case odds >= 20:
			conName = fmt.Sprintf(YellowU("%s:%d%%"), f.Name, odds)
		case odds >= 0:
			conName = fmt.Sprintf(RedU("%s:%d%%"), f.Name, odds)
		}
	*/

	var conName string
	odds := 0
	conName = fmt.Sprintf(RedU("%s:%d%%"), f.Name, odds)
	// For sanity layout the StatusBar vertically here although printing horizonal
	fmt.Printf("%s", Black("    :"))
	fmt.Printf("%s", RedU(conName))
	fmt.Printf("%s", Yellow(" S:"))
	fmt.Printf("%s", Green(strconv.Itoa(f.Str.Val)))
	fmt.Printf("%s", Yellow(" I:"))
	fmt.Printf("%s", Green(strconv.Itoa(f.Int.Val)))
	fmt.Printf("%s", Yellow(" D:"))
	fmt.Printf("%s", Green(strconv.Itoa(f.Dex.Val)))
	fmt.Printf("%s", Yellow(" HP:"))
	fmt.Printf("%s", Green(strconv.Itoa(f.Health.BaseVal)))
	fmt.Printf("%s", Yellow("|"))
	fmt.Printf("%s", Red(strconv.Itoa(f.Health.Val)))
	/*
		fmt.Printf("%s", Yellow(" W:"))
		fmt.Printf("%s", White(f.Weap.Name))
		fmt.Printf("%s", Yellow(" A:"))
		fmt.Printf("%s", White(f.Armor.Name))
		fmt.Printf("%s", Yellow(" T:"))
		fmt.Printf("%s", White(f.Trink.Name))
	*/
	fmt.Println()
	Meter(f.Health.Val, f.Health.BaseVal, f.FoeMaxHit, "Health", "█", "foe")
}
