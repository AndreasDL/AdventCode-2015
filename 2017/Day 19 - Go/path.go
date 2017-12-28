package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

type walker struct{
	field [][]byte
	x,y int
	dx, dy int // direction of last step
}

func main() {
	//input := readInput("testInput.txt")
	input := readInput("input.txt")

	res, ctr := "", 1 //start pos => already 1 step
	walker := new(input)
	char, finish := walker.step()
	for ; !finish ; char, finish = walker.step() {
		
		res += char
		ctr++
		/*fmt.Println( string(char), 
			finish, 
			walker.y , walker.x, 
			walker.dy, walker.dx,
			" => ", res,
		)*/
	}
	res += char

	fmt.Println(res)
	fmt.Println(ctr)
}

func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}
func new(s string) *walker {
	lines := strings.Split(s, "\n")

	res := make([][]byte, len(lines))
	for i, line := range lines {
		res[i] = []byte(line)
	}

	x := 0 ; for ; x < len(res[0]) && res[0][x] != '|' ; x++ {}

	return &walker{
		field : res,
		y     : 0,
		x     : x,
		dx    : 0,
		dy    : 1,
	}
}
func (f *walker) print() {
	for _, line := range f.field {
		fmt.Println(string(line))
	}
}
func (f *walker) step() (char string, finish bool) {
	
	//markPath
	if (f.dy != 0 && f.field[f.y][f.x] == '|') || (f.dx != 0 && f.field[f.y][f.x] == '-') { f.field[f.y][f.x] = '=' }

	char = ""
	//fmt.Print( string(f.field[f.y][f.x]), " => ")

	if f.field[f.y][f.x] == '+' { //change dir
		       
		if        f.dy !=  1 && f.y-1 > 0      && f.field[f.y-1][f.x] != ' ' { //up
			f.dx =  0
			f.dy = -1
		} else if f.dy != -1 && f.y+1 < len(f.field) && f.field[f.y+1][f.x] != ' ' { //down
			f.dx =  0
			f.dy =  1
		} else if f.dx !=  1 && f.x-1 > 0      && f.field[f.y][f.x-1] != ' ' { //left
			f.dx = -1
			f.dy =  0
		} else if f.dx != -1 && f.x+1 < len(f.field[0]) && f.field[f.y][f.x+1] != ' ' { //right
			f.dx =  1 
			f.dy =  0
		}
	} else if !strings.ContainsAny( string(f.field[f.y][f.x]), "=+-|" ){
		//return char at position & boolean that indicates if we are at the finish !
		char   = string(f.field[f.y][f.x])
		finish = f.field[f.y+f.dy][f.x+f.dx] == ' '//check if done
		//f.print()
	}

	//Step
	f.y += f.dy
	f.x += f.dx

	return 
}