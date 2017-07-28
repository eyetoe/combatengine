package main

type Stat struct {
	Name    string
	BaseVal int
	Val     int
}

func (s Stat) Modifier() int {
	return s.Val / 5
}
