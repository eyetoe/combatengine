package main

type Condition struct {
	Name         string
	BaseDuration int
	Duration     int
	Affect       func(Agent) Agent
}

func (c *Condition) Apply() {
	c.Duration = c.BaseDuration
	return
}

var Regeneration = Condition{
	Name:         "Natural Health Recovery",
	BaseDuration: 0,
	Affect: func(a Agent) Agent {
		if a.Health.Val < a.Health.BaseVal {
			a.Health.Val = a.Health.Val + 1
		}
		return a
	},
}

var OnFire = Condition{
	Name:         "On Fire!",
	BaseDuration: 3,
	Affect: func(a Agent) Agent {
		if a.Health.Val > 0 {
			a.Health.Val = a.Health.Val - 1
		}
		return a
	},
}
