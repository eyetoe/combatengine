package main

type Coord struct {
	x        int
	y        int
	Location Location
}

type World [][]Coord

func NewWorld() World {
	var height = 20
	var width = 20
	var new World
	for c := 1; c <= height; c++ {
		var row []Coord
		for r := 1; r <= width; r++ {
			row = append(row, Coord{x: c, y: r})
		}
		new = append(new, row)
	}
	return new
}
