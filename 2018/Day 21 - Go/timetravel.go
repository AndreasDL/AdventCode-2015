package main


import (
	"strconv"
	"regexp"
	"strings"
	"fmt"
	"io/ioutil"
)


//instructions
var strToFunc = map[string]func([6]int,int,int,int)[6]int {
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}
func addr(registers [6]int, a,b,c int) [6]int{
	registers[c] = registers[a] + registers[b]
	return registers
}
func addi(registers [6]int, a,b,c int) [6]int{
	registers[c] = registers[a] + b
	return registers
}
func mulr(registers [6]int, a,b,c int) [6]int{
	registers[c] = registers[a]*registers[b]
	return registers
}
func muli(registers [6]int, a,b,c int) [6]int{
	registers[c] = registers[a]*b
	return registers
}
func banr(registers [6]int, a,b,c int) [6]int{
	registers[c] = registers[a] & registers[b]
	return registers
}
func setr(registers [6]int, a,b,c int) [6]int{
	registers[c] = registers[a]
	return registers
}
func seti(registers [6]int, a,b,c int) [6]int{
	registers[c] = a
	return registers
}
func gtir(registers [6]int, a,b,c int) [6]int{
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0 
	}
	return registers
}
func gtri(registers [6]int, a,b,c int) [6]int{
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0 
	}	
	return registers
}
func gtrr(registers [6]int, a,b,c int) [6]int{
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0 
	}
	return registers
}
func eqir(registers [6]int, a,b,c int) [6]int{
	if a == registers[b]{
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}
func eqri(registers [6]int, a,b,c int) [6]int{
	if registers[a] == b{
		registers[c] = 1
	} else {
		registers[c] = 0
	}	
	return registers
}
func eqrr(registers [6]int, a,b,c int) [6]int{
	if registers[a] == registers[b]{
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}
func bani(registers [6]int, a,b,c int) [6]int{
	registers[c] = registers[a] & b
	return registers
}
func borr(registers [6]int, a,b,c int) [6]int{
	registers[c] = registers[a] | registers[b]
	return registers
}
func bori(registers [6]int, a,b,c int) [6]int{
	registers[c] = registers[a] |  b
	return registers
}

//int conversion
func toInt(s string) int{
	v, _ := strconv.Atoi(s)
	return v
}
var re = regexp.MustCompile("\\d+")
func parseInts(line string) []int{
	parts := re.FindAllString(line, 3)
	result := make([]int, len(parts))
	for i, p := range parts{
		result[i] = toInt(p)
	}

	return result
}

//instruction struct
type Instr struct{
	f func([6]int,int,int,int)[6]int
	a,b,c int
	name string
	line int
}
func(i Instr) execute(registers [6]int) [6]int{
	return i.f(registers, i.a, i.b, i.c)
}
func parseInstr(line string, nbr int) *Instr{

	parts := strings.Split(line, " ")
	numbers := parseInts(line)

	return &Instr{
		strToFunc[parts[0]],
		numbers[0],
		numbers[1],
		numbers[2],
		parts[0],
		nbr,
	}
}
func parseInput(input string) ([]*Instr, int){

	lines := strings.Split(input, "\n")
	
	result := make([]*Instr, len(lines)-1)
	for i, line := range lines[1:]{
		result[i] = parseInstr(line, i)
	}

	ip := parseInts(lines[0])[0]

	return result, ip
}
func readFile(fname string) string {
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}


func part1(instr []*Instr, ip int){

	registers := [6]int{0,0,0,0,0,0}
	for i := 0; registers[ip] < len(instr) && registers[ip] >= 0; i++ {

		ins := instr[registers[ip]]
		registers = ins.execute(registers)
		registers[ip]++

		if ins.line == 28 { break }
	}

	part1 := registers[3]
	fmt.Println(registers, " => r0 should be", part1)

	//check if that is really the case
	//can be removed
	registers = [6]int{part1, 0,0,0,0,0}
	for i := 0; registers[ip] < len(instr) && registers[ip] >= 0; i++ {

		ins := instr[registers[ip]]
		registers = ins.execute(registers)
		registers[ip]++
	}

	fmt.Println("Part 1: ", part1, " made the program end")
}


func part2(instr []*Instr, ip int){

	//find lower int that makes the program halt
	//part 1 => r0 == r3 halts => halt at first check
	//now exit as late at possible => stop once we detect loops on value r3
	//loop => means we will run forever
	//so once we detect a loop we set r0 to the last r3 before the loop happens
	//that way we'll end right before the loop starts

	last := 0
	seen := map[int]bool{}
	registers := [6]int{0,0,0,0,0,0}

	//run program
	for i := 0; registers[ip] < len(instr) && registers[ip] >= 0; i++ {

		ins := instr[registers[ip]]
		registers = ins.execute(registers)
		registers[ip]++

		if ins.line == 28 { 

			if _, loop := seen[registers[3]]; loop {
				fmt.Println("Part2", last, registers)
				return
			}
			seen[registers[3]] = true
			last = registers[3]
		}
	}
	
}

func main(){

	input := readFile("input.txt")
	instr, ip := parseInput(input)
	part1(instr, ip)
	part2(instr, ip)
}