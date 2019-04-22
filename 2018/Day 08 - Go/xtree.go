package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)

func parseInput(fname string) *[]string{
	file, _ := ioutil.ReadFile(fname)
	res := strings.Split(string(file), " ")
	return &res
}

//Part 1
func parseChild(_parts *[]string, pos, sum int, recurse bool) (int, int){
	parts := *_parts

	//fmt.Println("Parsing", pos, sum)
	child_cnt, _ := strconv.Atoi(parts[pos])
	pos++
	meta_cnt , _ := strconv.Atoi(parts[pos])
	pos++

	//parse childs first
	s := 0
	for i := 0; i < child_cnt; i++ {
		pos, s = parseChild(_parts, pos, sum, recurse)

		if recurse { sum = s }
	}

	//get meta
	for i := 0; i < meta_cnt; i++ {
		val, _ := strconv.Atoi(parts[pos])
		//fmt.Println("adding", val)
		sum += val
		pos++
	}

	return pos, sum
}
func part1(){
}

//Part 2
type node struct{
	child_cnt, meta_cnt int
	value int
	values []int
	children []*node
}
func buildTree(_parts *[]string, pos int) (*node, int){
	parts := *_parts

	child_cnt, _ := strconv.Atoi(parts[pos])
	pos++
	meta_cnt , _ := strconv.Atoi(parts[pos])
	pos++

	current := node{
		child_cnt, 
		meta_cnt, 
		0,
		make([]int, meta_cnt),
		make([]*node, child_cnt),
	}

	//parse childs first
	for i := 0; i < child_cnt; i++ {
		current.children[i], pos = buildTree(_parts, pos)
	}

	//get meta
	for i := 0; i < meta_cnt; i++ {
		val, _ := strconv.Atoi(parts[pos])
		current.value += val
		current.values[i] = val
		pos++
	}

	return &current, pos
}
func (root node) checksum() int{

	//if a node has no childnodes, it's value is the sum of its metavalues
	if len(root.children) == 0 {
		return root.value
	}

	//childnodes => use metadata as indexes
	//sum the checksum values
	sum := 0
	for _, val := range root.values {
		if val-1 < len(root.children) {
			sum += root.children[val-1].checksum()
		}
	}

	return sum
}


func main(){
	/*
	_parts := strings.Split(
		"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", 
		" ",
	)
	parts := &_parts
	*/
	
	parts := parseInput("input.txt")

	_, sum := parseChild(parts, 0, 0, true)
	fmt.Println("Part 1:", sum)

	root, _ := buildTree(parts, 0)
	fmt.Println("Part 2:", root.checksum())
}