package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)


func readFile( fname string) string{
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}

type Field struct{
	positions map[int]map[int]byte
	y_min, y_max int
}
func (f Field) print(){

	for y := 0 /*f.y_min*/; y <= f.y_max; y++{
		for x:= 400; x < 600; x++{

			if b, ex := f.positions[y][x] ; ex {
				fmt.Print(string(b))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func (f *Field) flow(y,x int){

	b, ex := f.positions[y][x]
	for ; y <= f.y_max && !ex ; {
		f.positions[y][x] = '|'

		y++
		b, ex = f.positions[y][x] //look to next position
	}

	if y <= f.y_max && (b == '#' || b == '~') { //when we hit ground or existing water we can fill
	 	f.fill(y-1,x) 
	}
}
func (f Field) canFill(y, x int) (bool, bool) { //canFill(y,x) => canWaterFlow ?, overflow?
	
	//water can flow sideways as long as there is something to run on
	_ , ex2 := f.positions[y+1][x] 
	if !ex2 { return false, true }


	//check if we hit a wall
	b1, ex1 := f.positions[y][x]
	return !ex1 || b1 != '#', false
}
func (f *Field) fill(y,x int){ //fill bucket

	if y <= 0 || y >= f.y_max { return }

	isFree := false
	left_x , left_overflow  := x, false
	right_x, right_overflow := x, false
	for ; y > 0 && !left_overflow && !right_overflow ; y-- {

		//left
		left_x = x
		isFree, left_overflow = f.canFill(y,left_x)
		for ; isFree ; isFree, left_overflow = f.canFill(y, left_x) {
			f.positions[y][left_x] = '~'
			left_x--
		}

		//right
		right_x = x
		isFree, right_overflow = f.canFill(y,right_x)
		for ; isFree ; isFree, right_overflow = f.canFill(y, right_x) {
			f.positions[y][right_x] = '~'
			right_x++
		}
	}

	if left_overflow || right_overflow {
		y++ //correct y => y is updated AFTER every loop

		//set overflow line to '|'
		for x := left_x + 1; x < right_x ; x++ {
			f.positions[y][x] = '|'
		}

		//recurse, ignore overflow when they have already happend!
		if left_overflow && f.positions[y+1][left_x+1] != '|' {
			f.flow(y,left_x) 
		}
		if right_overflow && f.positions[y+1][right_x-1] != '|' { 
			f.flow(y,right_x) 
		}
	}
}
func (f Field) count() int{
	cnt := 0
	for y := f.y_min ; y <= f.y_max ; y++ {
		for _, v := range f.positions[y] {
			
			if v == '~' || v == '|' { cnt++ }
		}
	}

	return cnt
}
func (f Field) count2() int {
	cnt := 0
	for y := f.y_min ; y <= f.y_max ; y++ {
		for _, v := range f.positions[y] {
			
			if v == '~' { cnt++ }
		}
	}

	return cnt
}
func parseLine(line string) (int, int, int, bool){

	parts := strings.Split(line, ",")
	first_parts := strings.Split(parts[0], "=")
	x, _ := strconv.Atoi(first_parts[1])
	is_reversed := first_parts[0] == "y"

	parts = strings.Split(parts[1], "=")
	y_start, _ := strconv.Atoi(strings.Split(parts[1], "..")[0])
	y_stop , _ := strconv.Atoi(strings.Split(parts[1], "..")[1])

	return x, y_start, y_stop, is_reversed
}
func parseField(input string) *Field{

	lines := strings.Split(input, "\n")
	result := make(map[int]map[int]byte, len(lines))

	x, y_min, y_max, is_reversed := parseLine(lines[0])
	if is_reversed { 
		y_min, y_max = x,x
	}
	
	for _, line := range lines {

		x, y_start, y_stop, is_reversed := parseLine(line)

		if is_reversed{
			if _, ex := result[x] ; !ex { result[x] = map[int]byte{} }
			for y := y_start; y <= y_stop; y++{ result[x][y] = '#' }
		} else {

			if y_start < y_min { y_min = y_start }
			if y_stop  > y_max { y_max = y_stop  }

			for y := y_start; y <= y_stop; y++{ 
				if _, ex := result[y] ; !ex { result[y] = map[int]byte{} }
				result[y][x] = '#' 
			}
		}
		
	}

	//initiliase all lines
	for y := 0 ; y <= y_max ; y++ {
		if _, ex := result[y]; !ex { result[y] = map[int]byte{} }
	}

	f := Field{
		result,
		y_min,
		y_max,
	}
	return &f
}


func main(){
	input := `x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504`

	input = readFile("input.txt")

	f := parseField(input)

	/*
	f.y_max = 137
	for k, _ := range f.positions{
		if k > f.y_max  { delete(f.positions, k) }
	}
	//*/

	f.flow(0,500)
	//f.print()
	fmt.Println("Part1:", f.count())
	fmt.Println("Part2:", f.count2())


}