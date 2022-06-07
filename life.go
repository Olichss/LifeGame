package main

import "fmt"

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	a := make([][]bool, height)
	for i := range a {
		a[i] = make([]bool, width)
	}
	return Universe(a)
}

func (u Universe) Show() {
	for _, rows := range u {
		for _, cell := range rows {
			if cell {
				fmt.Print("+")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}
