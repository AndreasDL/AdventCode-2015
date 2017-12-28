package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)

func main(){
	in1 := []int{0,3,0,1,-3}
	in2 := readInput("input.txt")

/*	
	fmt.Println( doStep(&in1, 0, 0) )
	fmt.Println()
	fmt.Println( doStep(in2, 0, 0) )
	fmt.Println()
*/

	//IN PLACE!! don't run both at same time !

	fmt.Println( doStep2(&in1, 0, 0) )
	fmt.Println()

	fmt.Println( doStep2(in2, 0, 0) )
	fmt.Println()


}


func doStep(ptr *[]int, pos int, ctr int)int {

	//if ctr == 25 { panic(nil) }

	input := *ptr

	//fmt.Println(pos, input, ctr)
	if pos >= len(input){
		return ctr
	}

	newPos := pos + input[pos]

	if pos < len(input) {
		input[pos]++
		return doStep(ptr, newPos, ctr+1)
	}

	return ctr
}

func doStep2(ptr *[]int, pos int, ctr int)int {
	input := *ptr

	for pos < len(input){

		//if ctr == 25 { panic(nil) } //debug

		newPos := pos + input[pos]
		
		if input[pos] >= 3 {
			input[pos]--
		} else {
			input[pos]++
		}

		pos = newPos
		ctr++

		//fmt.Println(pos, input, ctr) //debug
	}

	return ctr
}

func readInput(fname string) *[]int {
	s, _ := ioutil.ReadFile(fname)

	res := []int{}
	for _, line := range strings.Split( string(s), "\n"){

		nbr, _ := strconv.Atoi(line)
		res = append(res, nbr)
	}

	return &res
}