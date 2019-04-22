package main


import (
	"fmt"
	"strconv"
	"container/heap" //https://golang.org/pkg/container/heap/
)

type Point struct{
	x,y int
}
var caveTypes = []string{".","=","|"}

//Field represents the cave & uses cache to keep track of geology types
type Field struct{
	levels map[Point]int
	depth int
	target_x, target_y int
}
func createField(target_x, target_y, depth int) *Field{
	levels := map[Point]int{}
	levels[Point{0,0}] = 0
	levels[Point{target_x, target_y}] = 0

	return &Field{ 
		levels,
		depth,
		target_x, target_y,
	}
}
func (f *Field) getLevel(x,y int) int{
	
	//already calculated
	if lvl, ex := f.levels[Point{x,y}]; ex {
		return lvl
	}

	//calculate value
	var lvl int
	if y == 0 {
		lvl = x*16807
	} else if x == 0 {
		lvl = y * 48271	
	} else {
		//recurse
		lvl = f.getLevel(x-1,y) * f.getLevel(x,y-1)
	}

	lvl += f.depth
	lvl %= 20183
	f.levels[Point{x,y}] = lvl
	
	return lvl
}
func (f *Field) fill(){
	for y := 0; y <= f.target_y; y++ {
		for x := 0; x <= f.target_x ; x++ {
			f.getLevel(x,y)
		}
	}	
}
func (f *Field) print(size int){
	for y := 0; y < size; y++ {
		for x := 0; x < size ; x++ {
			
			caveType := f.getLevel(x,y) % 3
			fmt.Println(caveTypes[caveType])			
		}
		fmt.Println()
	}
}
func (f *Field) getType(x,y int) int{
	return f.getLevel(x,y) % 3
}
func (f *Field) riskLevel() int {
	sum := 0
	for y := 0 ; y <= f.target_y; y++ {
		for x := 0; x <= f.target_x ; x++ {
			sum += f.getType(x,y) % 3
		}
	}
	return sum
}
func (f *Field) next(s State) []State{
	delta_x := [4]int{0,-1,0,1}
	delta_y := [4]int{1,0,-1,0}

	possible_next := make([]State, 0,16)
	for i := 0; i < 4 ; i++ {

		//new coordinates
		nx := s.x + delta_x[i]
		ny := s.y + delta_y[i]
		if nx < 0 || ny < 0 { continue }

		//possible next states
		possible_next = append( //no change of gear
			possible_next, 
			State{
				nx, ny, s.time+1,
				s.climbing_gear,
				s.torch,
		})
		possible_next = append( //change to neither
			possible_next, 
			State{
				nx, ny, s.time+8,
				false,
				false,
		})
		possible_next = append( //change to climbing gear
			possible_next,
			State{
				nx, ny, s.time+8,
				true,
				false,
		})
		possible_next = append( //change to torch
			possible_next, 
			State{
				nx, ny, s.time+8,
				false,
				true,
		})
	}

	//only keep valid states
	nextStates := make([]State, 0, 16)
	for _, ns := range possible_next {

		validnow := validGear(ns.climbing_gear, ns.torch, f.getType(s.x, s.y))
		validnext := validGear(ns.climbing_gear, ns.torch, f.getType(ns.x, ns.y))

		if validnow && validnext {
			nextStates = append(nextStates, ns)	
		}
	}

	return nextStates
}

//state for bfs
type State struct{
	x,y int
	time int
	climbing_gear bool
	torch bool
}
func (s State) hash() string{
	res := strconv.Itoa(s.x) + "|" 
	res += strconv.Itoa(s.y) + "|" 
	if s.climbing_gear { 
		res += "1"
	} else {
		res += "0"
	}
	
	res += "|"

	if s.torch {
		res += "1"
	} else {
		res += "0"
	}

	return res
}

//priority queue using heap ensure that we loop in right order (fastest paths first)
type PriorityQueue []State
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int)bool { 
	if (pq[i].time < pq[j].time) {
		return true
	} else if (pq[i].time > pq[j].time){
		return false
	} else {
		first_torch := pq[i].torch 
		second_torch := pq[j].torch

		if first_torch {
			return true
		} else if !first_torch && second_torch {
			return false
		} else { //!first_torch && !second_torch
			return true
		}
	}
}
func (pq PriorityQueue) Swap(i,j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(s interface{}){
	*pq = append(*pq, s.(State))
}
func (pq *PriorityQueue) Pop() interface{}{
	last := len(*pq) -1
	res := (*pq)[last]
	*pq = (*pq)[:last]
	
	return res
}

//checks if gear is valid for a gived caveType
func validGear(climbingGear, torch bool, caveType int) bool{
	if caveType == 0 {
		return climbingGear || torch
	} else if caveType == 1 {
		return !torch
	} else if caveType == 2 {
		return !climbingGear
	}

	panic("type " + string(caveType) + " unknown")
}



func part1(target_x, target_y, depth int) int {
	f := createField(target_x, target_y, depth)
	return f.riskLevel()
}


func (f *Field) bfs(start State) int { //best first search

	todo := make(PriorityQueue, 0, 10000)
	todo = append(todo, start)
	heap.Init(&todo)

	seen := map[string]bool{}
	for cnt := 0; len(todo) > 0; cnt++ {

		//todo.pop()
		v := heap.Pop(&todo).(State)

		if s, ex := seen[v.hash()] ; s || ex { continue }

		//check if we are at target!
		if v.x == f.target_x && v.y == f.target_y && v.torch {
			fmt.Println("found", v)
			return v.time
		}	

		//seen
		seen[v.hash()] = true

		//next states
		for _, w := range f.next(v) {
			if s, ex := seen[w.hash()]; s || ex { continue }
			
			heap.Push(&todo, w)
		}
	}

	fmt.Println(todo)
	fmt.Println(seen)
	panic("not found")
}


func main(){


	if part1(10,10,510) != 114 { panic("broken </3") }


	depth := 8112
	target_x := 13
	target_y := 743
	fmt.Println("Part 1:", part1(target_x, target_y, depth))
	
	field := createField(target_x, target_y, depth)
	start := State{0,0,0, false, true}
	fmt.Println("Part 2:", field.bfs(start))
}