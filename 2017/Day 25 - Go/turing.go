package main

import "fmt"

func main() {
	
	machine := TMachine{
		memory  : map[int]int{},
		state   : 'A',
		currPos : 0,
	}

	
	for i := 0 ; i < 12317297 ; i++ {
		machine.step()
	}

	fmt.Println(machine.checksum())

}
type TMachine struct {
	memory map[int]int
	state byte
	currPos int	
}

func (m *TMachine) checksum() int{
	ctr := 0
	for _, v := range m.memory{ ctr+=v	}
	return ctr
}
func (m *TMachine) step(){
	switch m.state {
		case 'A': m.stepA()
		case 'B': m.stepB()
		case 'C': m.stepC()
		case 'D': m.stepD()
		case 'E': m.stepE()
		case 'F': m.stepF()
	}
}
func (m *TMachine) stepA(){
	if m.memory[m.currPos] == 0 {
		m.memory[m.currPos] = 1
		m.currPos++
		m.state = 'B'
	} else {
		m.memory[m.currPos] = 0
		m.currPos--
		m.state = 'D'
	}
}
func (m *TMachine) stepB(){
	if m.memory[m.currPos] == 0 {
		m.memory[m.currPos] = 1
		m.currPos++
		m.state = 'C'
	} else {
		m.memory[m.currPos] = 0
		m.currPos++
		m.state = 'F'
	}
}
func (m *TMachine) stepC(){
	if m.memory[m.currPos] == 0 {
		m.memory[m.currPos] = 1
		m.currPos--
		m.state = 'C'
	} else {
		m.memory[m.currPos] = 1
		m.currPos--
		m.state = 'A'
	}
}
func (m *TMachine) stepD(){
	if m.memory[m.currPos] == 0 {
		m.memory[m.currPos] = 0
		m.currPos--
		m.state = 'E'
	} else {
		m.memory[m.currPos] = 1
		m.currPos++
		m.state = 'A'
	}
}
func (m *TMachine) stepE(){
	if m.memory[m.currPos] == 0 {
		m.memory[m.currPos] = 1
		m.currPos--
		m.state = 'A'
	} else {
		m.memory[m.currPos] = 0
		m.currPos++
		m.state = 'B'
	}
}
func (m *TMachine) stepF(){
	if m.memory[m.currPos] == 0 {
		m.memory[m.currPos] = 0
		m.currPos++
		m.state = 'C'
	} else {
		m.memory[m.currPos] = 0
		m.currPos++
		m.state = 'E'
	}
}