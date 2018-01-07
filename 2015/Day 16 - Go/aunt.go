package main

import (
	"strings"
	"io/ioutil"
	"fmt"
)

var properties = map[string]string {
	"children"    : "3",
	"cats"		  : "7",
	"samoyeds"    : "2",
	"pomeranians" : "3",
	"akitas"	  : "0",
	"vizslas"	  : "0",
	"goldfish"	  : "5",
	"trees"		  : "3",
	"cars"		  : "2",
	"perfumes"    : "1",
}


func readInput(fname string) string {
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}
func main() {
	input := readInput("input.txt")

	//part1
	candidates1 := []int{}
	for i, line := range strings.Split(input, "\n"){
		if isPossible1(line) { candidates1 = append(candidates1, i+1) }		
	}
	fmt.Println("Part1: ", candidates1)


	//part2
	candidates2 := []int{}
	for i, line := range strings.Split(input, "\n"){
		if isPossible2(line) { candidates2 = append(candidates2, i+1) }		
	}
	fmt.Println("Part2: ", candidates2)

}

func isPossible1(line string) bool{
	line = strings.Replace(line, ",", "", -1)
	line = strings.Replace(line, ":", "", -1)

	fields := strings.Fields(line)
	for j := 2 ; j < len(fields) ; j += 2 {
		if properties[ fields[j] ] != fields[j+1] { return false }
	}

	return true
}
func isPossible2(line string) bool{
	line = strings.Replace(line, ",", "", -1)
	line = strings.Replace(line, ":", "", -1)

	fields := strings.Fields(line)
	for j := 2 ; j < len(fields) ; j += 2 {

		item     := fields[j]
		shouldBe := properties[ item ]
		actual   := fields[j+1]

		switch item{
		case "cats", "trees" : if actual <= shouldBe  { return false }
		case "pomeranians", "goldfish" : if actual >= shouldBe { return false }
		default: if actual != shouldBe { return false }
		}
	}

	return true
}

