package main

import (
	"strconv"
	"strings"
	"fmt"
	"io/ioutil"
)


type Prog struct {
	memory map[string]int
	instructions []string
	toggles []bool
}

func (p *Prog) cpy(x,y string){
	p.memory[y] = p.strToInt(x)
}
func (p *Prog) inc(x string){
	p.memory[x]++
}
func (p *Prog) dec(x string){
	p.memory[x]--
}
func (p *Prog) jnz(x, y string)int{
	val := 0
	if p.strToInt(x) != 0 {
		val = p.strToInt(y)
		val-- //+1 for pos after every instruction
	}
	return val
}
func (p *Prog) tgl(pos int, x string){
	val := pos + p.strToInt(x)
	if val >= 0 && val < len(p.toggles) {
		p.toggles[val] = !p.toggles[val]
	}
}
func (p Prog) strToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil { val = p.memory[s]}
	return val
}
func (p *Prog) Execute() int {
	for pos := 0 ; pos >= 0 && pos < len(p.instructions) ; pos++ {
		instr := p.instructions[pos]

		fields := strings.Fields(instr)

		//inspired by https://markheath.net/post/aoc-2016-day23
		if pos == 5 || pos == 21 { //shortcut a += (c*d)
			p.memory["a"] += abs(p.memory["c"]) * abs(p.memory["d"])
			pos += 4
			continue //skip other stuff !
		}

		if !p.toggles[pos] {
			switch fields[0]{
				case "cpy": p.cpy(fields[1], fields[2])
				case "inc": p.inc(fields[1])
				case "dec": p.dec(fields[1])
				case "jnz": pos += p.jnz(fields[1], fields[2])
				case "tgl": p.tgl(pos, fields[1])
			}
		} else {
			switch fields[0]{
				case "cpy": pos += p.jnz(fields[1], fields[2])
				case "inc": p.dec(fields[1])
				case "dec": p.inc(fields[1])
				case "jnz": p.cpy(fields[1], fields[2])
				case "tgl": p.inc(fields[1])
			}
		}
	}

	return p.memory["a"]
}

func main() {
	sample := `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`
	
	sample = readInput("input.txt")

	instructions := strings.Split(sample, "\n")
	prog := Prog{
		memory       : map[string]int{"a": 12},
		instructions : instructions,
		toggles      : make([]bool, len(instructions)),
	}

	fmt.Println(prog.Execute())
}

func readInput(fname string) string {
	res , _ := ioutil.ReadFile(fname)
	return string(res)
}
func abs(v int) int{
	if v < 0 { return -v }
	return v
}