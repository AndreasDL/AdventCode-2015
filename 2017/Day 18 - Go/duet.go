package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)

var registers map[string]int
var lastValue int

func init(){ registers = map[string]int{} } //init the map on package load !

//register read & writes
func snd(x string){ 
	lastValue = registers[x]
}
func set(x string, y string){
	if val, err := strconv.Atoi(y) ; err == nil {
		registers[x] = val	
	} else {
		registers[x] = registers[y]
	}
}
func add(x string, y string){
	if val, err := strconv.Atoi(y) ; err == nil {
		registers[x] += val
	} else {
		registers[x] += registers[y]
	}
}
func mul(x string, y string){
	if val, err := strconv.Atoi(y) ; err == nil {
		registers[x] *= val	
	} else {
		registers[x] *= registers[y]
	}
}
func mod(x string, y string){ 
	if val, err := strconv.Atoi(y) ; err == nil {
		registers[x] %= val	
	} else {
		registers[x] %= registers[y]
	}
}
func jgz(x string)int { //reg value or value on its own ?
	if val, err := strconv.Atoi(x) ; err == nil {
		return val
	} else {
		return registers[x]
	}
}

func main() {
	//input := readInput("testInput.txt")
	input := readInput("input.txt")
	
	instructions := *parseInstructions(input)
	pos := 0

	for pos < len(instructions) && pos >= 0 {

		instr := strings.Split(instructions[pos], " ")

		switch instr[0] {
			case "snd": snd(instr[1])
			case "set": set(instr[1],instr[2])
			case "add": add(instr[1],instr[2])
			case "mul": mul(instr[1],instr[2])
			case "mod": mod(instr[1],instr[2])
			case "rcv": 
				if jgz(instr[1]) != 0 {
					fmt.Println(lastValue)
					fmt.Println(registers)
					panic(nil) //fuckit lets quit !
				}
			case "jgz":
				if jgz(instr[1]) > 0 {
					pos += jgz(instr[2]) - 1 //pos++ comes later => -1
				}
		}
		
			
		pos++
		fmt.Println(instr, registers, lastValue, pos)
	}

	fmt.Println(lastValue)
	fmt.Println(registers)
}


func parseInstructions(input string) *[]string{
	res := strings.Split(input, "\n")
	return &res
}
func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}