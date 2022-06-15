package main

import (
	"fmt"
	"math/rand"
	"time"
)

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// Seed заполняет вселенную случайными живыми клетками.
func (u Universe) Seed() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < (width * height / 4); i++ {
		u[rand.Intn(height)][rand.Intn(width)] = true
	}
}

func normilizeCoord(x, size int) int {
	if x < 0 {
		return normilizeCoord(x+size, size)
	}
	return x % size
}

func (u Universe) Alive(x, y int) bool {
	// x = (x + width) % width
	// y = (y + height) % height

	x = normilizeCoord(x, width)
	y = normilizeCoord(y, height)
	return u[y][x]
}

// Neighbors подсчитывает прилегающие живые клетки.
func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}

	return string(buf)
}

func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
	fmt.Println("---------------------------------------------------------------------------")
}

func (u Universe) isAlive() bool {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if u[y][x] {
				return true
			}
		}
	}

	return false
}

func Step(a, b Universe) bool {
	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			b[x][y] = a.Next(x, y)
		}
	}

	return b.isAlive()
}
