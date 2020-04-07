package main

import (
	"fmt"
	"time"
	"strings"
	"math/rand"
)

const (
	SizeH int = 10
	SizeW int = 20
	Tick string = "250ms"
)

type State int
const (
	Dead State = iota
	Alive
)

type World struct {
	cells [SizeH][SizeW]Cell
}
type Cell struct {
	state State
}

func main() {
	w := newWorld()
	s, _ := time.ParseDuration(Tick)

	seedWorld(&w)

	for {
		print(w)
		w = generation(w)
		time.Sleep(s)
	}
}

func newWorld() World {
	return World{[SizeH][SizeW]Cell{}}
}

func seedWorld(w *World) {
	for i, row := range w.cells {
		for j, _ := range row {
			if rand.Float64() > 0.5 {
				w.cells[i][j] = Cell{Alive}
			}
		}
	}
}

func generation(w World) World {
	newWorld := newWorld()
	for i, row := range w.cells {
		for j, _ := range row {
			newWorld.cells[i][j] = evaluate(w, i, j)
		}
	}
	return newWorld
}

func evaluate(prevWorld World, row int, col int) Cell {
	prevCell := prevWorld.cells[row][col]
	newCell := prevCell
	neighbours :=  getNeighbours(prevWorld, row, col)
	num_alive := 0

	for _, cell := range neighbours {
		if cell.state == Alive { num_alive++ }
	}

	// dead cells with 3 parents becomes alive
	if prevCell.state == Dead && num_alive == 3 {
		newCell = Cell{Alive}

	// alive cells without 2 or 3 parents, dies
	} else if prevCell.state == Alive && (num_alive < 2 || num_alive > 3) {
		newCell = Cell{Dead}
	}

	return newCell
}

func getNeighbours(world World, row int, col int) []Cell {
	left := col-1
	right := col+1
	above := row-1
	below := row+1

	if(above < 0) {above = SizeH-1}
	if(below == SizeH) {below = 0}
	if(left < 0) {left = SizeW-1}
	if(right == SizeW) {right = 0}

	cells := []Cell{
		world.cells[above][left],
		world.cells[above][col],
		world.cells[above][right],
		world.cells[row][left],
		world.cells[row][right],
		world.cells[below][left],
		world.cells[below][col],
		world.cells[below][right],
	}
	return cells
}

func print(w World) {
	fmt.Println("+" + strings.Repeat("-", SizeW) + "+")
	for _, row := range w.cells {
		fmt.Printf("|")
		for _, cell := range row {
			if(cell.state == Dead) {
				fmt.Printf(" ")
			} else {
				fmt.Printf("o")
			}
		}
		fmt.Printf("|\n")
	}
	fmt.Println("+" + strings.Repeat("-", SizeW) + "+")
	fmt.Println()
}

