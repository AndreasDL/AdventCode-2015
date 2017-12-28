package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

type Node struct{
	x, y int
	size, used, free int
}
func NewNode(line string) Node{
	parts := strings.Fields(line)

	locParts := strings.Split(parts[0], "-")
	x, _     := strconv.Atoi(locParts[1][1:])
	y, _     := strconv.Atoi(locParts[2][1:])

	size, _  := strconv.Atoi(parts[1][:len(parts[1])-1])
	used, _  := strconv.Atoi(parts[2][:len(parts[2])-1])
	free, _  := strconv.Atoi(parts[3][:len(parts[3])-1])

	return Node{
		x    : x,
		y    : y,
		size : size,
		used : used,
		free : free,
	}
}
func (a *Node) IsViablePair(b Node) bool{
	
	if a.used == 0 { 
		return false 
	} else if a.x == b.x && a.y == b.y {
		return false 
	}

	return a.used <= b.free
}
func (a *Node) hash() string{
	return fmt.Sprintf("%d-%d-%d-%d", a.x, a.y, a.used, a.size)
}
func (n Node) String() string {
	if n.used > 100 {
		return "#"
	}
	if n.used == 0 {
		return "_"
	}
	if n.x == 36 && n.y == 0 {
		return "G"
	}
	return "."
}
	
func main() {
	Part1()
	fmt.Println()

	fmt.Println("Field looks like this:")
	Part2()
}

func Part1(){
	input := readInput("realInput.txt")
	lines := strings.Split(input, "\n")[2:]

	nodes := make([]Node, len(lines))
	for i, line := range lines{
		nodes[i] = NewNode(line)
	}

	ctr := 0
	for i := 0 ; i < len(nodes) ; i++ {
		for j := 0 ; j < len(nodes) ; j++ {
			if i == j { continue }
			
			if nodes[i].IsViablePair(nodes[j]) { ctr++ }
		}
	}
	fmt.Println("Viable pairs:", ctr)
}
func Part2(){
	input := readInput("realInput.txt")
	lines := strings.Split(input, "\n")[2:]

	nodes := make([][]Node, 25)
	for i := range nodes { nodes[i] = make([]Node, 37) }

	for _, line := range lines{
		node := NewNode(line)
		nodes[node.y][node.x] = node
	}

	//output
	for _, line := range nodes {
		for _, n := range line {
			fmt.Print(n)
		}
		fmt.Println()
	}
}
func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}