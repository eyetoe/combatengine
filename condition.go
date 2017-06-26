package main

type Condition struct {
	Name         string
	BaseDuration int
	Duration     int
	//Affect       func(Agent) Agent
	Affect func(*Agent)
}

/*
func (c *Condition) Apply() {
	c.Duration = c.BaseDuration
	return
}
*/

var Regeneration = Condition{
	Name:         "Natural Health Recovery",
	BaseDuration: 0,
	Affect: func(a *Agent) {
		a.HealthAdjust(1)
		return
	},
}

var OnFire = Condition{
	Name:         "OnFire",
	BaseDuration: 3,
	Affect: func(a *Agent) {
		a.HealthAdjust(-2)
		return
	},
}
