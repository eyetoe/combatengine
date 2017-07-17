package main

type Location struct {
	Name        string
	Description string
	Exits       []Coord
	Events      []Event
}
