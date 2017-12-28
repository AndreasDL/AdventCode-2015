package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

type direction int
const (
	U direction = iota
	R
	D
	L
)
var dxdy = map[direction]struct{dx, dy int}{
	U: {  0,  1},
	R: {  1,  0},
	D: {  0, -1},
	L: { -1,  0},
}

var example = `..#
#..
...`

type health int
const (
	C health = iota
	W
	I
	F
)


type Field struct {
	nodes map[int]map[int]health //nodes[y][x] = isInfected()
	x,y int
	dir direction
	ctr int
}


func main() {
		
	//example
	field := NewField()
	field.loadMap(example)
	for i := 0 ; i < 10000000 ; i++ { field.step() }
	fmt.Println(field.ctr)


	//part2
	field = NewField()
	field.loadMap( readInput("input.txt") )
	for i := 0 ; i < 10000000 ; i++ { field.step() }
	fmt.Println(field.ctr)
}

func NewField() *Field{
	return &Field{
		nodes: map[int]map[int]health{},
		x: 0,
		y: 0,
		dir: U,
	}
}
func (f *Field) step(){
	switch f.nodes[f.y][f.x] {
		case C: f.dir = (f.dir+3)%4
		case I: f.dir = (f.dir+1)%4
		case F: f.dir = (f.dir+2)%4
	}

	//lookup fix !
	if _, ex := f.nodes[f.y] ; !ex { f.nodes[f.y] = map[int]health{} }
		
	//new state
	f.nodes[f.y][f.x] = (f.nodes[f.y][f.x]+1)%4
	if f.nodes[f.y][f.x] == I { 
		f.ctr++ 
	}

	f.x += dxdy[f.dir].dx
	f.y += dxdy[f.dir].dy
}
func (f *Field) loadMap(s string){
	f.nodes = map[int]map[int]health{}
	
	lines := strings.Split(s, "\n")
	offset := int(len(lines)/2)

	for j, y := 0, offset ; y >= -offset ; j,y = j+1, y-1 {
		line := lines[j]
		if _, ex := f.nodes[y] ; !ex { f.nodes[y] = map[int]health{} }
		
		for i,x := 0, -offset ; x <= offset ; i,x = i+1, x+1 {
			switch line[i]{
				case '.': f.nodes[y][x] = C
				case '#': f.nodes[y][x] = I
				case 'W': f.nodes[y][x] = W
				case 'F': f.nodes[y][x] = F
			}
		}
	}
}
func (f *Field) print(){
	r := int(len(f.nodes)/2) + 1
	for i := r ; i >= -r ; i-- {
		for j := -r ; j <= r ; j++ {

			if h, ex := f.nodes[i][j] ; ex {
				switch h {
					case C: fmt.Print(".")
					case W: fmt.Print("W")
					case F: fmt.Print("F")
					case I: fmt.Print("#")
				}
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
	fmt.Println()
}

func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}