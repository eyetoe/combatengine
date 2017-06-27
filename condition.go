package main

type Condition struct {
	Name     string
	Duration int
	Affect   func(*Agent)
}

var Regeneration = Condition{
	Name:     "Natural Healing",
	Duration: 0,
	Affect: func(a *Agent) {
		a.HealthAdjust(1)
		return
	},
}

var OnFire = Condition{
	Name:     "OnFire",
	Duration: 3,
	Affect: func(a *Agent) {
		a.HealthAdjust(-2)
		return
	},
}
