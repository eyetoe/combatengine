package main

import (
	"fmt"
	"strconv"
)

type Screen struct {
	Width  int
	Height int
}

//func ( Screen) Resize() {
func (s Screen) Resize() {
	//fmt.Printf("[8;60;120t")
	fmt.Printf("[8;" + strconv.Itoa(s.Height) + ";" + strconv.Itoa(s.Width) + "t")
	return
}

func (s Screen) Layout() {
	for h := 0; h < s.Height-2; h++ {
		for w := 0; w < s.Width; w++ {
			fmt.Printf("#")
		}
	}
	return
}
