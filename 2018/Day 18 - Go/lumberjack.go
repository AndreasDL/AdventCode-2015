package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func min(a,b int) int{
	if a < b { return a}
	return b
}
func max(a,b int) int{
	if a > b { return a}
	return b
}
func readFile(fname string) string{
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}


type Field struct{
	positions *[][]byte
}
func (f Field) height() int{
	return len(*f.positions)
}
func (f Field) width() int {
	return len((*f.positions)[0])
}
func parseField(input string) *Field{

	lines := strings.Split(input, "\n")
	positions := make([][]byte, len(lines))

	for i, line := range lines {
		positions[i] = []byte(line)
	}

	return &Field{ &positions}
}
func (f Field) print(){
	for _, line := range *f.positions {
		fmt.Println(string(line))
	}
	fmt.Println()
}
func (f Field) adjacent(y,x int) (int, int, int){

	cnts := make(map[byte]int, 3)
	for _y := max(0, y-1) ; _y < min(f.height(), y+2) ; _y++ {
		for _x := max(0, x-1) ; _x < min(f.width(), x+2) ; _x++ {

			cnts[ (*f.positions)[_y][_x] ]++
		}
	}

	cnts[ (*f.positions)[y][x] ]-- //correct for own tile
	return cnts['.'], cnts['|'], cnts['#']
}
func (f *Field) next(){
	old_positions := *(f.positions)

	new_positions := make([][]byte, len(old_positions))
	for y, line := range old_positions{
		new_positions[y] = make([]byte, len(line))

		for x, b := range line {

			_, cnt_tree, cnt_lumber := f.adjacent(y,x)

			if b == '.' && cnt_tree >= 3 {
				new_positions[y][x] = '|'
			} else if b == '|' && cnt_lumber >= 3 {
				new_positions[y][x] = '#'
			} else if b == '#' && (cnt_lumber == 0 || cnt_tree == 0) {
				new_positions[y][x] = '.'
			} else {
				new_positions[y][x] = old_positions[y][x]
			}
		}
	}

	f.positions = &new_positions
}
func (f Field) count() int{

	cnt_tree := 0
	cnt_lumber := 0
	for _, line := range *f.positions {
		for _, b := range line {

			if b == '|' { 
				cnt_tree++ 
			} else if b == '#' {
				cnt_lumber++
			}
		}
	}


	return cnt_tree * cnt_lumber
}


func part1(){
	input := `.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.`

	input = readFile("input.txt")
	field := parseField(input)
		for i := 0; i < 10 ; i++ {
		field.next()
	}
	field.print()
	fmt.Println("Part1:", field.count())
}

func main() {


	
	part1()

	
	input := readFile("input.txt")
	field := parseField(input)

	/*	
	//find pattern 8000 == 1000 == 15000 => repetition each 7000 patterns
	for j := 1; j <= 21000; j+= 1 {
		field.next()
		if j % 1000 == 0 { fmt.Println(j, " => ", field.count()) }
	}

	//reset
	input = readFile("input.txt")
	field = parseField(input)
	*/


	iterations := 1000000000
	iterations %= 7000 //simplify
	for i := 0; i < iterations ; i++ {
		field.next()
	}
	fmt.Println("Part 2:", field.count())

}