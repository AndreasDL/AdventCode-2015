package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)

type prog struct{
	id int
	registers map[string]int
	in  <-chan int //input from other prog
	out chan<- int //output to other prog
	sch chan<- int //send counter to main thread
	ctr int
}
func NewProg(id int, in <-chan int, out chan<- int, sch chan<- int) *prog{
	return &prog{
		id		 : id,
		registers: map[string]int{"p": id}, 
		in: in,
		out: out,
		sch: sch,
		ctr: 0,
	}
}
func (p *prog) set(x string, y string){
	if val, err := strconv.Atoi(y) ; err == nil {
		p.registers[x] = val	
	} else {
		p.registers[x] = p.registers[y]
	}
}
func (p *prog) add(x string, y string){
	if val, err := strconv.Atoi(y) ; err == nil {
		p.registers[x] += val
	} else {
		p.registers[x] += p.registers[y]
	}
}
func (p *prog) mul(x string, y string){
	if val, err := strconv.Atoi(y) ; err == nil {
		p.registers[x] *= val	
	} else {
		p.registers[x] *= p.registers[y]
	}
}
func (p *prog) mod(x string, y string){ 
	if val, err := strconv.Atoi(y) ; err == nil {
		p.registers[x] %= val	
	} else {
		p.registers[x] %= p.registers[y]
	}
}
func (p *prog) jgz(x string)int { //abuse this function => reg value or value on its own ?
	if val, err := strconv.Atoi(x) ; err == nil {
		return val
	} else {
		return p.registers[x]
	}
}
func (p *prog) snd(x string){
	p.out <- p.jgz(x)
	p.ctr++
	p.sch <- p.ctr
}
func (p *prog) rcv(x string){
	p.registers[x] = <- p.in
}
func (p *prog) run(instructions []string){

	pos := 0
	for pos < len(instructions) && pos >= 0 {

		instr := strings.Split(instructions[pos], " ")
		switch instr[0] {
			case "set": p.set(instr[1],instr[2])
			case "add": p.add(instr[1],instr[2])
			case "mul": p.mul(instr[1],instr[2])
			case "mod": p.mod(instr[1],instr[2])
			case "snd": p.snd(instr[1])
			case "rcv": p.rcv(instr[1])
			case "jgz":
				if p.jgz(instr[1]) > 0 {
					pos += p.jgz(instr[2]) - 1 //pos++ comes later => -1
				}
		}
		pos++
		//fmt.Println(p.id, instr, p.registers, pos)
	}
	
	close(p.out)
}



func main() {
	//input := readInput("testInput2.txt")
	input := readInput("input.txt")
	instructions := *parseInstructions(input)
	
	ch0to1, ch1to0 := make(chan int, 1000), make(chan int, 1000)
	snd0, snd1 := make(chan int), make(chan int)

	prog0 := NewProg(0, ch1to0, ch0to1, snd0)
	prog1 := NewProg(1, ch0to1, ch1to0, snd1)

	go prog0.run(instructions)
	go prog1.run(instructions)

	go func(){ 
		for _, ok := <-snd0 ; ok ; _,ok = <-snd0 {} //consume snd0
	}() 

	//write out all snd1 values => last one is the correct one
	for val, ok := <-snd1 ; ok ; val,ok = <-snd1 {
		fmt.Println(val)
	}
}

func parseInstructions(input string) *[]string{
	res := strings.Split(input, "\n")
	return &res
}
func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}