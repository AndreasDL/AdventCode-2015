package main

import (
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
	"regexp"
)

func addr(registers [4]int, a,b,c int) [4]int{
	registers[c] = registers[a] + registers[b]
	return registers
}
func addi(registers [4]int, a,b,c int) [4]int{
	registers[c] = registers[a] + b
	return registers
}
func mulr(registers [4]int, a,b,c int) [4]int{
	registers[c] = registers[a]*registers[b]
	return registers
}
func muli(registers [4]int, a,b,c int) [4]int{
	registers[c] = registers[a]*b
	return registers
}
func banr(registers [4]int, a,b,c int) [4]int{
	registers[c] = registers[a] & registers[b]
	return registers
}
func bani(registers [4]int, a,b,c int) [4]int{
	registers[c] = registers[a] & b
	return registers
}
func borr(registers [4]int, a,b,c int) [4]int{
	registers[c] = registers[a] | registers[b]
	return registers
}
func bori(registers [4]int, a,b,c int) [4]int{
	registers[c] = registers[a] |  b
	return registers
}
func setr(registers [4]int, a,b,c int) [4]int{
	registers[c] = registers[a]
	return registers
}
func seti(registers [4]int, a,b,c int) [4]int{
	registers[c] = a
	return registers
}
func gtir(registers [4]int, a,b,c int) [4]int{
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0 
	}
	return registers
}
func gtri(registers [4]int, a,b,c int) [4]int{
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0 
	}	
	return registers
}
func gtrr(registers [4]int, a,b,c int) [4]int{
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0 
	}
	return registers
}
func eqir(registers [4]int, a,b,c int) [4]int{
	if a == registers[b]{
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}
func eqri(registers [4]int, a,b,c int) [4]int{
	if registers[a] == b{
		registers[c] = 1
	} else {
		registers[c] = 0
	}	
	return registers
}
func eqrr(registers [4]int, a,b,c int) [4]int{
	if registers[a] == registers[b]{
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}


func compSlices(a,b [4]int) bool{
	for i, aa := range a {
		if aa != b[i] { return false }
	}
	return true
}
func toInt(s string) int{
	v, _ := strconv.Atoi(s)
	return v
}
func parseInts(line string) [4]int{
	res := [4]int{}
	for i, v := range re.FindAllString(line, 4){ 
		res[i] = toInt(v)
	}
	return res
}

type Trace struct {
	before [4]int
	op,a,b,c int
	after [4]int
}
var re = regexp.MustCompile("\\d+")
func parseTrace(lines []string) *Trace{

	before := parseInts(lines[0])

	parts := parseInts(lines[1])
	op,a,b,c := parts[0], parts[1], parts[2], parts[3]
	
	after := parseInts(lines[2])

	return &Trace{
		before,
		op,a,b,c,
		after,
	}
}
func parseFile(fname string) []*Trace{
	file, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(file), "\n")

	traces := make([]*Trace, 0, len(lines)/4)
	for i,j := 0,3; j < len(lines); i,j=i+4,j+4{
		traces = append(traces, parseTrace(lines[i:j]))
	}
	return traces
}

var functions = []func([4]int,int,int,int) [4]int{
	addr,addi,
	mulr,muli,
	banr,bani,
	borr,bori,
	setr,seti,
	gtir,gtri,gtrr,
	eqir,eqri,eqrr,
}
var names = []string{
	"addr","addi",
	"mulr","muli",
	"banr","bani",
	"borr","bori",
	"setr","seti",
	"gtir","gtri","gtrr",
	"eqir","eqri","eqrr",
}

func part1(traces []*Trace) int{
	
	part1 := 0
	for _, t := range traces {
	
		cnt:=0
		for _, f := range functions{
			actual := f(t.before, t.a,t.b,t.c)
			//fmt.Println(t.before, " => ", names[i], " => ", actual)
			if compSlices(actual, t.after){ 
				cnt++ 
			}
		}

		if cnt >= 3 { part1++ }
	}
	return part1
}

func decode_opcodes(traces []*Trace) ([]func([4]int,int,int,int)[4]int, []string) {
	//init mappings & candidates
	unmapped_functions := make([]func([4]int,int,int,int)[4]int, len(functions))
	unmapped_function_names := make([]string, len(functions))
	for i, f := range functions {
		unmapped_functions[i] = f
		unmapped_function_names[i] = names[i]
	}
	opcode_translation := make([]func([4]int,int,int,int)[4]int, len(functions))
	opcode_translation_names := make([]string,len(functions))
	
	//bruteforce loop
	for len(unmapped_functions) > 0 {
		for _, t := range traces {
		
			cnt, fnd := 0, -1
			for i, f := range unmapped_functions{
				actual := f(t.before, t.a,t.b,t.c)
				
				if compSlices(actual, t.after){ 
					fnd = i
					cnt++ 
				}
			}
			if cnt == 1 { 
				opcode_translation[t.op] = unmapped_functions[fnd]
				opcode_translation_names[t.op] = unmapped_function_names[fnd]

				//remove elements
				unmapped_functions[fnd] = unmapped_functions[len(unmapped_functions)-1]
				unmapped_functions = unmapped_functions[:len(unmapped_functions)-1]

				unmapped_function_names[fnd] = unmapped_function_names[len(unmapped_function_names)-1]
				unmapped_function_names = unmapped_function_names[:len(unmapped_function_names)-1]
			}
		}
	}

	return opcode_translation, opcode_translation_names
}

func parseInputpart2(fname string) [][4]int{
	file, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(file), "\n")
	res := make([][4]int, len(lines))
	for i, line := range lines {
		res[i] = parseInts(line)
	}

	return res
}

func part2(decoded_functions []func([4]int,int,int,int)[4]int, instructions [][4]int) int{

	registers := [4]int{0,0,0,0}
	for _, instr := range instructions{

		op := instr[0]
		a,b,c := instr[1], instr[2], instr[3]

		f := decoded_functions[op]
		registers = f(registers, a,b,c)
	}

	return registers[0]
}

func main() {

	traces := parseFile("input_part1.txt")
	fmt.Println("Part1: ", part1(traces))
	
	decoded_functions, _ := decode_opcodes(traces)
	instructions := parseInputpart2("input_part2.txt")
	fmt.Println("Part2: ", part2(decoded_functions, instructions))

}