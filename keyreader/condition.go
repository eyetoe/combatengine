package main

import (
	//	. "github.com/eyetoe/foobarbaz/colors"
	color "github.com/fatih/color"
)

type Condition struct {
	Name     string
	Duration int
	Affect   func(*Agent)
	//Color    interface{}
	Color    func(...interface{}) string
	Describe map[string]string
}

var Regen = Condition{
	Name:     "Regen",
	Duration: 0,
	Color:    color.New(color.FgGreen).SprintFunc(),
	Affect: func(a *Agent) {
		a.HealthAdjust(1)
		return
	},
	Describe: map[string]string{
		"self":   "Your body regenerates.",
		"target": "It regenerates.",
	},
}

var Burning = Condition{
	Name:     "Burning",
	Duration: 3,
	Color:    color.New(color.FgRed).SprintFunc(),
	Affect: func(a *Agent) {
		a.HealthAdjust(-2)
		return
	},
	Describe: map[string]string{
		"self":   "You are on fire!",
		"target": "It is on fire!",
	},
}

var Feeble = Condition{
	Name:     "Feeble",
	Duration: 2,
	Color:    color.New(color.FgHiYellow).SprintFunc(),
	Affect: func(a *Agent) {
		a.Str.Val = a.Str.BaseVal / 2
		return
	},
	Describe: map[string]string{
		"self":   "You have been enfeebled!",
		"target": "It has been enfeebled!",
	},
}

var Cursed = Condition{
	Name:     "Cursed",
	Duration: 2,
	Color:    color.New(color.FgHiMagenta).SprintFunc(),
	Affect: func(a *Agent) {
		a.Dex.Val = a.Dex.BaseVal / 2
		return
	},
	Describe: map[string]string{
		"self":   "You have been cursed!",
		"target": "It has been cursed!",
	},
}

var Dazed = Condition{
	Name:     "Dazed",
	Duration: 2,
	Color:    color.New(color.FgYellow).SprintFunc(),
	Affect: func(a *Agent) {
		a.Int.Val = a.Int.BaseVal / 2
		return
	},
	Describe: map[string]string{
		"self":   "You are dazed and confused!",
		"target": "It is dazed and confused!",
	},
}

var Insight = Condition{
	Name:     "Insight",
	Duration: 10,
	Color:    color.New(color.FgHiBlue).SprintFunc(),
	Affect: func(a *Agent) {
		a.Int.Val = a.Int.Val + a.Int.BaseVal/2
		return
	},
	Describe: map[string]string{
		"self":   "You are filled with insight and understanding.",
		"target": "It is filled with insight and understanding.",
	},
}

var Paralized = Condition{
	Name:     "Paralized",
	Duration: 1,
	Color:    color.New(color.FgHiMagenta).SprintFunc(),
	Affect: func(a *Agent) {
		a.Focus = a.Focus / 2
		return
	},
	Describe: map[string]string{
		"self":   "You have been paralized. You can't move!",
		"target": "It has been paralized. It can't move!",
	},
}
var Petrified = Condition{
	Name:     "Petrified",
	Duration: 1,
	Color:    color.New(color.FgHiBlack).SprintFunc(),
	Affect: func(a *Agent) {
		a.Focus = 0
		return
	},
	Describe: map[string]string{
		"self":   "You've been turned to stone. You're petrified!",
		"target": "It has been turned to stone. It's petrified!",
	},
}
