package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

var memory = map[string]uint16{}

type Command struct {
	xin, yin, result string
	f func(uint16, uint16) uint16
}
func (c Command) CanExecute() bool{
	_, exx := memory[c.xin]
	_, exy := memory[c.yin]

	if len(c.xin) == 0 { exx = true }
	if len(c.yin) == 0 { exy = true }

	return exx && exy
}
func (c *Command) Execute() bool{

	if !c.CanExecute() { return false } 
	// => xin & yin should exist

	x,y := uint16(0), uint16(0)
	if len(c.xin) > 0 { x = memory[c.xin] }
	if len(c.yin) > 0 { y = memory[c.yin] }
	memory[c.result] = c.f(x,y)

	return true
}
func (c Command) String() string{
	return fmt.Sprintf("xin: %s yin: %s result: %s", c.xin, c.yin, c.result)
}



func main() {

	part1()

}

func part1(){
	commands := parseFile( readInput("input.txt") )
	//fmt.Println(len(commands))

	//panic(nil)
	i := 0
	for len(commands) > 1 {
		if commands[i].Execute() { //execution sucess full => remove instruction
			commands = append(commands[:i], commands[i+1:]...)
			//fmt.Println(i, " => ", len(commands))
		}
		i = (i+1) % len(commands)
	}

	//last one!
	commands[0].Execute()

	fmt.Println("part 1: ", memory["a"])	
}

func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}
func makeCommand(line string) Command{
	parts := strings.Fields(line)

	cmd := Command{
		xin    : "",
		yin    : "",
		result : parts[len(parts)-1],
	}
	if parts[0] == "NOT" {
		cmd.xin = parts[1] ; initMemoryWhenNeeded(parts[1])
		cmd.f   = func (x, y uint16) uint16 { return -x-1 }
	} else if parts[1] == "OR" {
		cmd.xin = parts[0] ; initMemoryWhenNeeded(parts[0])
		cmd.yin = parts[2] ; initMemoryWhenNeeded(parts[2])
		cmd.f   = func (x,y uint16) uint16 { return x | y }
	} else if parts[1] == "AND" {
		cmd.xin = parts[0] ; initMemoryWhenNeeded(parts[0])
		cmd.yin = parts[2] ; initMemoryWhenNeeded(parts[2])
		cmd.f   = func (x,y uint16) uint16 { return x & y }
	} else if parts[1] == "LSHIFT" {
		val, _ := strconv.Atoi(parts[2])
		cmd.xin = parts[0] ; initMemoryWhenNeeded(parts[0])
		cmd.f   = func(x,y uint16) uint16 { return x << uint(val) }
	} else if parts[1] == "RSHIFT" {
		val, _ := strconv.Atoi(parts[2])
		cmd.xin = parts[0] ; initMemoryWhenNeeded(parts[0])
		cmd.f   = func(x,y uint16) uint16 { return x >> uint(val) }
	} else {
		val, err := strconv.Atoi(parts[0])
		if err == nil {
			cmd.f   = func(x,y uint16) uint16 { return uint16(val) }	
		} else {
			cmd.xin = parts[0]
			cmd.f   = func(x,y uint16) uint16 { return x }
		}
		
	}
	
	return cmd
}
func initMemoryWhenNeeded(addr string){
	//if the input is a number, we just create a memory spot containing the number
	if val, err := strconv.Atoi(addr) ; err == nil {
		memory[addr] = uint16(val)
	}
}
func parseFile(s string) []Command{
	commands := []Command{}
	for _, line := range strings.Split(s, "\n") {
		commands = append(commands, makeCommand(line))
	}
	return commands
}