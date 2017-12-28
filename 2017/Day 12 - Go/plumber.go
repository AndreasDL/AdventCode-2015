package main

import (
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
)

func main(){
	input := readInput("input.txt")

	fmt.Println( InGroup(input) )
	fmt.Println( CountGroups(input) )
}


func InGroup(s string) int{
	graph := parseInput(s)
	res := connectedTo(0, graph)
	return len(*res)
}
func CountGroups(s string) int{
	graph := *(parseInput(s))
	
	//create stack with numbers in order
	candidates := stack{}
	for k, _ := range graph{
		candidates.push(k)
	}
	
	ctr := 0
	seen := map[int]bool{}
	for start := candidates.pop() ; start >= 0 ; start = candidates.pop(){
		if seen[start] { continue }

		ctr++
		groupMembers := *(connectedTo(start, &graph))
		for k, _ := range groupMembers {
			seen[k] = true
		}

	}

	return ctr
}

func parseInput(s string) *map[int]map[int]bool{
	res := map[int]map[int]bool{}
	for _, line := range strings.Split(s, "\n"){

		parts := strings.Split(line, " <-> ")
		
		from, _ := strconv.Atoi(parts[0])

		//init map
		if _, ex := res[from] ; !ex { res[from] = map[int]bool{} }
		
		//save the to values
		for _, t := range strings.Split(parts[1], ", ") {
			to, _ := strconv.Atoi(t)
			
			res[from][to] = true
		}

		
	}

	return &res
}
func connectedTo(start int, m *map[int]map[int]bool) *map[int]bool {
	graph := *m


	seen := map[int]bool{start: true}
	
	todos := stack{}
	for k, _ := range graph[start] {
		todos.push(k)
	}


	for key := todos.pop() ; key >= 0 ; key = todos.pop(){

		if seen[key] { continue }

		seen[key] = true
		for k, _ := range graph[key] {
			if !seen[k] { todos.push(k) }
		}

	}

	return &seen
}
func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}


type stack []int
func (s *stack) pop() int{
	l := len(*s)
	if l == 0 {
		return -1
	} else { 
		c := (*s)[l-1]
		*s = (*s)[:l-1] 
		return c
	}
}
func (s *stack) push(c int) {
	*s = append(*s, c)
}