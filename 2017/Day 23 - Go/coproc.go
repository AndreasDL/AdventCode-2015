package main

import (
	"strconv"
	"strings"
	"fmt"
	"io/ioutil"
)

type prog struct{
	//folding in sublime txt
	registers map[string]int
}
func NewProg() *prog{
	return &prog{
		registers: map[string]int{},
	}
}
func (p *prog) set(x string, y string){
	if val, err := strconv.Atoi(y) ; err == nil {
		p.registers[x] = val	
	} else {
		p.registers[x] = p.registers[y]
	}
}
func (p *prog) sub(x string, y string){
	if val, err := strconv.Atoi(y) ; err == nil {
		p.registers[x] -= val
	} else {
		p.registers[x] -= p.registers[y]
	}
}
func (p *prog) mul(x string, y string){
	if val, err := strconv.Atoi(y) ; err == nil {
		p.registers[x] *= val	
	} else {
		p.registers[x] *= p.registers[y]
	}
}
func (p *prog) jnz(x string, y string) int{

	xval, err := strconv.Atoi(x)
	if err != nil { xval = p.registers[x] }
	if xval == 0 { return 0 }

	yval, err := strconv.Atoi(y)
	if err != nil  { yval = p.registers[y] }

	return yval - 1  //pos++ follows later on
}
func (p *prog) run(instructions []string){

	pos := 0
	mul := 0
	for pos < len(instructions) && pos >= 0 {

		instr := strings.Split(instructions[pos], " ")
		switch instr[0] {
			case "set": p.set(instr[1], instr[2])
			case "sub": p.sub(instr[1], instr[2])
			case "mul": p.mul(instr[1],instr[2]) ; mul++
			case "jnz": pos += p.jnz(instr[1], instr[2])
		}
		//fmt.Println(pos, instr, p.registers)

		pos++
	}
	
	fmt.Println(mul)
}

func main() {
	instructions := *(parseInstructions( readInput("input.txt") ))
	prog := NewProg()

	prog.run(instructions)

}
func readInput(fname string) *string{
	input, _ := ioutil.ReadFile(fname)
	res := string(input)
	return &res
}
func parseInstructions(input *string) *[]string{
	res := strings.Split(*input, "\n")
	return &res
}