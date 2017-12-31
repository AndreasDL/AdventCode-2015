package main

import (
	"strings"
	"strconv"
	"io/ioutil"
	"fmt"
)

type Grid struct {
	lights [][]int
}

func (g *Grid) turnOn (startY, stopY, startX, stopX int){
	for y := startY ; y <= stopY ;  y++ {
		for x := startX ; x <= stopX ; x++ {
			g.lights[y][x]++
		}
	}
}

func (g *Grid) turnOff (startY, stopY, startX, stopX int){
	for y := startY ; y <= stopY ;  y++ {
		for x := startX ; x <= stopX ; x++ {
			if g.lights[y][x] > 0 { g.lights[y][x]-- }
		}
	}
}

func (g *Grid) toggle (startY, stopY, startX, stopX int){
	for y := startY ; y <= stopY ;  y++ {
		for x := startX ; x <= stopX ; x++ {
			g.lights[y][x]+=2
		}
	}
}

func (g *Grid) handleLine(s string){
	parts := strings.Fields(s)

	if parts[0] == "toggle"{
		startY, startX := parsePart(parts[1])
		stopY , stopX  := parsePart(parts[3])
		g.toggle(
			startY, stopY,
			startX, stopX,
		)
	} else if parts[1] == "on" {
		startY, startX := parsePart(parts[2])
		stopY , stopX  := parsePart(parts[4])
		g.turnOn(
			startY, stopY,
			startX, stopX,
		)
	} else {
		startY, startX := parsePart(parts[2])
		stopY , stopX  := parsePart(parts[4])
		g.turnOff(
			startY, stopY,
			startX, stopX,
		)
	}
}

func (g Grid) totalBrightness() int{
	ctr := 0
	for _, line := range g.lights {
		for _, l := range line {
			ctr += l
		}
	}
	return ctr
}

func newGrid() Grid{
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	return Grid{ lights: grid}
}

func parsePart(s string) (int,int){
	parts := strings.Split(s, ",")
	y, _ := strconv.Atoi(parts[0])
	x, _ := strconv.Atoi(parts[1])
	return y,x
}

func readInput(fname string) string{
	res , _ := ioutil.ReadFile(fname)
	return string(res)
}

func main() {
	input := readInput("input.txt")
	grid := newGrid()

	for _, line := range strings.Split(input, "\n"){
		grid.handleLine(line)
	}

	fmt.Println(grid.totalBrightness())
}