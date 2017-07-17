package main

import (
	"fmt"
	"testing"
)

func TestWorld(t *testing.T) {

	w := NewWorld()

	fmt.Printf("%+v", w)

	/*
		if err != nil {

			t.Errorf("Creating world %q", w)

		}
	*/
	/*
		cases := []struct {
			in, want string
		}{
			{"Hello, world", "dlrow ,olleH"},
			{"Hello, 世界", "界世 ,olleH"},
			{"", ""},
		}
		for _, c := range cases {
			got := Reverse(c.in)
			if got != c.want {
				t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	*/
}
