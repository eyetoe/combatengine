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
	Color func(...interface{}) string
}

var Regeneration = Condition{
	Name:     "Natural Healing",
	Duration: 0,
	Color:    color.New(color.FgGreen).SprintFunc(),
	Affect: func(a *Agent) {
		a.HealthAdjust(1)
		return
	},
}

var OnFire = Condition{
	Name:     "OnFire",
	Duration: 3,
	Color:    color.New(color.FgRed).SprintFunc(),
	Affect: func(a *Agent) {
		a.HealthAdjust(-2)
		return
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
}

var Cursed = Condition{
	Name:     "Cursed",
	Duration: 2,
	Color:    color.New(color.FgYellow).SprintFunc(),
	Affect: func(a *Agent) {
		a.Dex.Val = a.Dex.BaseVal / 2
		return
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
}

var Insight = Condition{
	Name:     "Insight",
	Duration: 10,
	Color:    color.New(color.FgHiBlue).SprintFunc(),
	Affect: func(a *Agent) {
		a.Int.Val = a.Int.Val + a.Int.BaseVal/2
		return
	},
}
