package main

import(
	"fmt"
)



//object
type State struct{
	e1, e2 int
	scores []byte
}
func (s *State) next(){

	//combine recipies to get new scores
	sum := s.scores[s.e1] + s.scores[s.e2]

	if sum > 9 {
		score1 := sum /10
		score2 := sum %10
		s.scores = append(s.scores, score1, score2)	
	} else {
		s.scores = append(s.scores, sum)
	}
	
	//move
	s.e1 = (s.e1 + int(s.scores[s.e1]) + 1) % len(s.scores)
	s.e2 = (s.e2 + int(s.scores[s.e2]) + 1) % len(s.scores)
}
func (s State) contains(input []byte) (bool,int){

	if len(s.scores) < len(input){ return false, 0 }

	offset := len(s.scores) - len(input)
	found := compSlices(input,s.scores[offset:])
	
	//two elements can be added
	if !found && offset > 1{
		offset--
		found = compSlices(input,s.scores[offset:offset+len(input)])
	}	

	return found, offset
}

//helper functions
func compSlices(a, b []byte) bool{

	if len(a) != len(b) { return false }

	for i, aa := range a{
		if aa != (b)[i] { return false }
	}

	return true
}

//Parts
func part1(n int) []byte{

	scores := make([]byte, 0, 100000)
	scores = append(scores, 3, 7)
	state := State{ 0,1, scores }

	for i:= 0; len(state.scores) < n+10 ; i++{
		state.next()
	}
	
	return state.scores[n:n+10]
}
func part2(searched []byte) int{
	scores := make([]byte, 0, 200000000)
	scores = append(scores, 3, 7)
	state := State{ 0,1, scores }

	found, offset := state.contains(searched)
	for i := 0 ; !found && len(state.scores) < 100000000; i++ {
		state.next()
		found, offset = state.contains(searched)
	}
	
	if !found { fmt.Println("not complete!") }
	return offset
}

func main(){
	fmt.Println("Part 1:", part1(824501))
	fmt.Println("Part 2:", part2([]byte{8,2,4,5,0,1}))
}