package main

import (
	"bytes"
	"fmt"
	"strconv"

	. "github.com/eyetoe/foobarbaz/colors"
)

// Meter(c.Health.Val, c.MaxHealth.Val, c.FoeMaxHit, "Health", "█")

// Meter provides a health meter to represent health scaled to barWidth.
func BetterMeter(num, max, fmax int, l, t string) string {
	// calculate the actual percentage
	p := float32(num) / float32(max)

	barWidth := 80

	// prepare as percentage of barWidth columns
	percentBar := p * float32(barWidth)

	// display meter
	fmt.Printf("%s	%s	:", Yellow(l), White(strconv.Itoa(num)))

	var buffer bytes.Buffer

	// drawing from left to right
	for c := 1; c <= barWidth; c++ {

		if c <= int(percentBar) {
			fmt.Printf(Cyan(t))
		} else {
			if num <= fmax {
				buffer.WriteString(Red(t))
			} else {
				buffer.WriteString(Black(t))
			}
		}
		if c == barWidth {
			fmt.Println(":")
		}
	}
	return buffer.String()

}

/*

// Meter(c.Health.Val, c.MaxHealth.Val, c.FoeMaxHit, "Health", "█")

// Meter provides a health meter to represent health scaled to barWidth.
func BetterMeter(num, max, fmax int, l, t string) {
	// calculate the actual percentage
	p := float32(num) / float32(max)

	barWidth := 80

	// prepare as percentage of barWidth columns
	percentBar := p * float32(barWidth)

	// display meter
	fmt.Printf("%s	%s	:", Yellow(l), White(strconv.Itoa(num)))

	// drawing from left to right
	for c := 1; c <= barWidth; c++ {

		if c <= int(percentBar) {
			fmt.Printf(Cyan(t))
		} else {
			if num <= fmax {
				fmt.Printf(Red(t))
			} else {
				fmt.Printf(Black(t))
			}
		}
		if c == barWidth {
			fmt.Println(":")
		}
	}

}
*/
