package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

var exampleInput = `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

var noPositions = 9

func main() {
	//part2
	field := NewField(readInput("input.txt"))
	fmt.Println(field.Search())

}



type Point struct { x,y int }

type Field struct {
	fld [][]byte
}
func NewField(s string) Field{
	lines := strings.Split(s, "\n")
	res := make([][]byte, len(lines))
	for i, line := range lines { res[i] = []byte(line) }
	return Field{res}
}
func (f Field) String() string {
	res := ""
	for _, line := range f.fld {
		res += string(line) + "\n"
	}
	return res
}
func (f Field) StartLocation() Point{
	for y, line := range f.fld {
		for x, c := range line {		
			if c == '0' {	return Point{x,y} }
		}
	}
	return Point{}
}
func (f Field) NextStates(s State) []State {
	res := []State{}

	for _, delta := range []Point{ {-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		x := s.currPos.x + delta.x
		y := s.currPos.y + delta.y
		
		
		if y < 0 || y >= len(f.fld) || x < 0 || x >= len(f.fld[y])  {
			continue
		}

		c := string(f.fld[y][x])
		if c == "#" { continue }

		found := s.found
		if c != "." && !strings.Contains(found, c){
			//did we find something ?

			if c != "0"{
				found += c
			} else if len(found) == noPositions-1 {
				found += c
			}
		}

		res = append(res, State{
			distance : s.distance+1,
			currPos  : Point{x,y},
			found    : found,
		})
	}

	return res
}
func (f Field) Search() int {

	seen  := map[string]bool{}

	todos := make(chan State, 100000) //channel as queue
	todos <- State{
		found    : "_",
		distance : 0,
		currPos  : f.StartLocation(),
	}

	for curr := range todos {

		if seen[curr.Hash()] { 
			if len(todos) == 0 { close(todos) }
			continue 
		}
		seen[curr.Hash()] = true

		if curr.IsFinal(noPositions) { 
			close(todos)
			fmt.Println(curr)
			return curr.distance
		}

		for _, state := range f.NextStates(curr){ todos <- state }
		if len(todos) == 0 { close(todos) }
	}

	return -1
}

type State struct{
	distance int
	currPos Point
	found string
}
func (s State) IsFinal(positions int) bool{
	return len(s.found) == positions && s.found[positions-1] == '0'
}
func (s State) Hash() string{
	return fmt.Sprintf("%s At: %d,%d",
		s.found, 
		s.currPos.x, 
		s.currPos.y,
	)
}
func (s State) IsEmpty() bool {
	return len(s.found) <= 0
}

type stack []State
func (s *stack) pop() State{
	l := len(*s)
	if l == 0 {
		return State{}
	} else { 
		c := (*s)[l-1]
		*s = (*s)[:l-1] 
		return c
	}
}
func (s *stack) push(c State) { *s = append(*s, c) }

func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}