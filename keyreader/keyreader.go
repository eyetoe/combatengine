package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

//Choice, is something you may choose from a MenuList
type Choice struct {
	Item string
	//Action func(...interface{})
	Action func()
	Hotkey string
}

type MenuList struct {
	Name        string
	Description string
	Choices     []Choice
	Tabs        []Choice
	Header      func()
}

// Choose()
func (m MenuList) Choose() {
	for {
		if m.Header != nil {
			ClearScreen()
			m.Header()
		}
		if m.Name != "" {
			fmt.Println(YellowU(m.Name), "\n")
		}
		if m.Description != "" {
			fmt.Println(Blue(m.Description))
		}

		if m.Tabs != nil {
			fmt.Printf("\r")
			for _, o := range m.Tabs {
				fmt.Printf("%s%s%s%s%s", "(", o.Hotkey, ")", o.Item, " ") // (S)pells
			}
			fmt.Printf("\r\n\n")
		}

		if m.Choices != nil {
			for i, c := range m.Choices {
				fmt.Printf("\r%s%s%s%s%s\n", "(", GreenU(i), ")", " ", c.Item)
			}
		}

		fmt.Printf(Blue("\r> "))

		select {
		case choice := <-AlphaInput:
			if choice == choice {
				fmt.Println("Alpha input detected")
			}
		case choice := <-IntInput:
			if choice == choice {
				fmt.Println("Integer input detected")
			}
		}

		/*
			choice, _, _ := GetChar()
			fmt.Println("choice was:", choice)

			if c, err := strconv.Atoi(choice); err == nil && c < len(m.Choices) {
				fmt.Printf("%T, %v", c, c)
				fmt.Println()
				m.Choices[c].Action()
				break
			} else {
				continue
			}
		*/
	}

}

func KeyReader() {
	for {
		choice, _, _ := GetChar()
		if choice == "\x03" {
			os.Exit(0)
		}
		if i, err := strconv.Atoi(choice); err == nil {
			IntInput <- i
		} else {
			AlphaInput <- choice
		}
	}
}
func AlphaKeyPrinter() {
	for i := range AlphaInput {
		fmt.Printf("Alpha: %q	%q\r\n", i, time.Since(startTime))
	}
}
func IntKeyPrinter() {
	for i := range IntInput {
		fmt.Printf("Int: %d		%q\r\n", i, time.Since(startTime))
	}
}

var AlphaInput chan string
var IntInput chan int
var startTime = time.Now()

func TestKeyReader() {

	AlphaInput = make(chan string)
	IntInput = make(chan int)
	go KeyReader()
	go AlphaKeyPrinter()
	go IntKeyPrinter()
	select {} // block forever
}
