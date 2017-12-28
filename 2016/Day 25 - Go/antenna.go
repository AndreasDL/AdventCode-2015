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
func (p Prog) strToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil { val = p.memory[s]}
	return val
}
func (p *Prog) Execute(l int) string {

	res := ""
	for pos := 0 ; pos >= 0 && pos < len(p.instructions) && len(res) < l ; pos++ {
		instr := p.instructions[pos]

		fields := strings.Fields(instr)

		//inspired by https://markheath.net/post/aoc-2016-day23
		//https://markheath.net/post/aoc-2016-day25
		//and https://www.reddit.com/r/adventofcode/comments/5k6yfu/2016_day_25_solutions/
		if pos == 3 { //shortcut a += (c*d)
			p.memory["d"] += abs(p.memory["b"]) * abs(p.memory["c"])
			pos += 4
			continue //skip other stuff !
		}

		switch fields[0]{
			case "cpy": p.cpy(fields[1], fields[2])
			case "inc": p.inc(fields[1])
			case "dec": p.dec(fields[1])
			case "jnz": pos += p.jnz(fields[1], fields[2])
			case "out":	res += strconv.Itoa(p.strToInt(fields[1]))
		}
	}

	return res
}

func main() {
	
	sample := readInput("input.txt")

	inputs := make(chan int, 1000000)
	outputs := make(chan int, 1000)
	done := make(chan bool, 1)

	for i := 0 ; i < 8 ; i++ {
		go func(input <-chan int, outputs chan<- int){
			for a := range input{
				prog := Prog{
					memory       : map[string]int{"a": a},
					instructions : strings.Split(sample, "\n"),
				}
				res := prog.Execute(8)

				if res == "01010101" || res == "10101010" {
					outputs <- a
				}
				if a%1000 == 0 { fmt.Println(a) }

				select{
					case <-done: return //done!
					default: //non blocking!
				}
			}
			return
		}(inputs, outputs)
	}
	
	for i := 0 ; i < 1000000 ; i++ { inputs <- i }
	close(inputs)
	results := make(chan int, 100)

	for i := 0 ; i < 2 ; i++ {
		go func(input <-chan int, outputs chan<- int){
			for a := range input{
				prog := Prog{
					memory       : map[string]int{"a": a},
					instructions : strings.Split(sample, "\n"),
				}
				res := prog.Execute(100)

				if res == strings.Repeat("01", 50) || res == strings.Repeat("10", 50) {
					outputs <- a
					done <- true
					close(outputs)
					close(done)
				}
				fmt.Println("tested", a)

				select{
					case <-done: return //done!
					default: //non blocking
				}
			}
			return
		}(outputs, results)
	}

	for res := range results {
		fmt.Println("found!: ", res)
	}
}

func readInput(fname string) string {
	res , _ := ioutil.ReadFile(fname)
	return string(res)
}
func abs(v int) int{
	if v < 0 { return -v }
	return v
}