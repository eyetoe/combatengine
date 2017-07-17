package main

import "fmt"

var Phases = []string{
	"Initiative",
	"Action",
	"Parry",
	"Finisher",
}

func PrintPhases() {
	for i := 0; i < len(Phases); i++ {
		fmt.Println(Phases[i])
	}
}
