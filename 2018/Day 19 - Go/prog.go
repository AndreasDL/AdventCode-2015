package main


import (
	"strconv"
	"regexp"
	"strings"
	"fmt"
	"io/ioutil"
	"math"
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


func part1(input string){

	instr, ip := parseInput(input)
	registers := [6]int{0,0,0,0,0,0}
	for i := 0; registers[ip] < len(instr) && registers[ip] >= 0; i++ {

		ins := instr[registers[ip]]

		fmt.Print(ins, registers, " => ")
		registers = ins.execute(registers)
		registers[ip]++
		fmt.Println(registers)
	}
	fmt.Println("Part 1:", registers[0])

}

func part2(input string){
	/*
	get trace and figure out what it is doing
	because register[0] is now 1 => we initialise 2 big numbers 

	&{0x4b3680 4 14 4 muli 31} [1 909 31 0 23550 0] => [1 909 32 0 329700 0]
	&{0x4b35f0 4 2 4 mulr 32} [1 909 32 0 329700 0] => [1 909 33 0 10550400 0]
	&{0x4b34e0 1 4 1 addr 33} [1 909 33 0 10550400 0] => [1 10551309 34 0 10550400 0]
	&{0x4b3990 0 5 0 seti 34} [1 10551309 34 0 10550400 0] => [0 10551309 35 0 10550400 0]
	&{0x4b3990 0 8 2 seti 35} [0 10551309 35 0 10550400 0] => [0 10551309 1 0 10550400 0]
	&{0x4b3990 1 1 5 seti 1} [0 10551309 1 0 10550400 0] => [0 10551309 2 0 10550400 1]
	&{0x4b3990 1 1 3 seti 2} [0 10551309 2 0 10550400 1] => [0 10551309 3 1 10550400 1]
	//A = 0
	//B = 10551309
	//IP = 0
	//C = 1
	//D = 10550400
	//E = 0

	we actually do the following program
	for c = range (1,B+1):
		for E range (1, B+1):
			if c*e == B:
				a += e
	*/

	instr, ip := parseInput(input)
	registers := [6]int{1,0,0,0,0,0}
	prev_a := registers[0]
	for i := 0; registers[ip] < len(instr) && registers[ip] >= 0; i++ {

		ins := instr[registers[ip]]

		//fmt.Print(ins, registers, " => ")
		registers = ins.execute(registers)
		registers[ip]++
		//fmt.Println(registers)

		if registers[0] != prev_a {
			fmt.Println(registers)
			prev_a = registers[0]	
		}

		//if i > 1e10 { break }
	}

	//find factors (optimized)
	a := 0
	b := 10551309
	lim := int(math.Sqrt(float64(b))) // factor cannot be greater than square root
	for c := 1; c <= lim ; c++ {

		if b % c == 0 { //find two factors at once, c and b/c
			a += c
			a += b/c

			fmt.Println(c, b/c)
		}
	}
	fmt.Println("Part 2: ", a)


}

func main(){


	input := `#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5`
	
	input = readFile("input.txt")
	//part1(input)

	part2(input)
	
}