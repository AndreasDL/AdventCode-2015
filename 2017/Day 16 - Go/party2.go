package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
)

type partners struct{
	offset int
	parties []rune
}

func (p * partners) ParseAndRun(cmd string) {
	switch rune(cmd[0]) {
	case 's':
		x, _ := strconv.Atoi(cmd[1:])
		p.Spin(x)
	case 'x':
		parts := strings.Split(cmd, "/")
		a, _ := strconv.Atoi(parts[0][1:])
		b, _ := strconv.Atoi(parts[1])
		p.Exchange(a,b)
	case 'p':
		parts := strings.Split(cmd, "/")
		a,b := rune(parts[0][1:][0]), rune(parts[1][0])
		p.Partner(a,b)
	}
}
func (p *partners) Spin(x int) {
	p.offset = (p.offset + p.length() - x) % p.length()
}
func (p *partners) Exchange(a,b int) {
	a = (p.offset + a) % p.length()
	b = (p.offset + b) % p.length()

	p.parties[a], p.parties[b] = p.parties[b], p.parties[a]
}
func (p *partners) Partner(a,b rune) {
	posA, posB := p.getPos(a), p.getPos(b)
	p.parties[posA], p.parties[posB] = p.parties[posB], p.parties[posA]	
}
func (p *partners) ToString() string{
	return string( 
		append(
			p.parties[p.offset:], 
			p.parties[:p.offset]...,
		),
	)
}
func (p *partners) getPos(a rune) int{
	i := 0
	for i < p.length() && p.parties[i] != a { i++ }
	return i
}
func (p *partners) length() int{
	return len(p.parties)
}
func (p *partners) print() {
	fmt.Println( p.ToString() )
}
func (p *partners) normalize(){
	//execute the spin & reset offset
	p.parties = append(p.parties[p.offset:], p.parties[:p.offset]...)
	p.offset = 0
}


func main(){

	//examples()
	
	input := readInput("input.txt")
	instructions := strings.Split(input,",")
	
	//part1
	Run(instructions, 1)


	//part 2
	//reduce instructions

	//we that if we execute all instructions 60 times we get the same starting point
/*	start := partners{0, []rune("abcdefghijklmnop")}
	seen := map[string]int{}
	for i := 0 ; i < 1000 ; i++ {
		for _, instr := range instructions {
			start.ParseAndRun(instr)
		}
		
		if pos, ex := seen[start.ToString()] ; ex {
			fmt.Println("loop! ", pos, " => ", i)
			// 0->60, 1 => 61 etc ...
		} else {
			seen[start.ToString()] = i
		}
	}
	
	Run(instructions, 60) //returns "abcedefghijklmnop"
*/
	

	//thus we can cheat and don't run all 1B iterations
	iterations := 1000000000 % 60
	Run(instructions, iterations)
}

func examples(){
	start := partners{
		offset: 0,
		parties: []rune("abcde"),
	}
	start.ParseAndRun("s1")
	start.print()

	start.ParseAndRun("x3/4")
	start.print()
	
	start.ParseAndRun("pe/b")
	start.print()
	
	fmt.Println()
	fmt.Println()
}

func Run(instructions []string, iterations int){
	//part1
	start := partners{0, []rune("abcdefghijklmnop")}
	for i := 0 ; i < iterations ; i ++ {
		for _, instr := range instructions {
			start.ParseAndRun(instr)
		}
	}
	start.print()
}

func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}