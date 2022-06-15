package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := NewUniverse(), NewUniverse()
	a.Seed()

	for {
		if !Step(a, b) {
			fmt.Println("Жизнь умерла")
		}
		a.Show()
		time.Sleep(time.Second / 30)
		a, b = b, a
	}

}
