package main

import "fmt"

type Event interface {
	Play(a *Agent)
}

// Encounter Event
type Encounter struct {
	Name string
	Foe  *Agent
}

func (e Encounter) Play(a *Agent) {
	fmt.Println("Playing an Encounter Event (interface) named:", e.Name)
}

// Merchant Event
type Merchant struct {
	Name   string
	Hero   *Agent
	Vendor *Agent
}

func (m Merchant) Play() {
	fmt.Println("Playing a Merchant Event (interface) named:", m.Name)
}
