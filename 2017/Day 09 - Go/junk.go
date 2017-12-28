package main

import (
	"fmt"
	"io/ioutil"
)


func main(){
	input := *readInput("input.txt")
	fmt.Println(ParseScore(input))
	fmt.Println(CollectGarbage(input))
}


func ParseScore(s string) int{
	//clean strings
	s = removeCancels(s)

	lvl := 0
	score :=  0
	inGarbage := false
	for _, c := range s {

		if c == rune('{') && !inGarbage {
			lvl++
			score+= lvl
		} else if c == rune('}') && !inGarbage {
			lvl--
		} else if c == rune('<') {
			inGarbage = true
		} else if c == rune('>') {
			inGarbage = false
		}
	}

	return score
}
func CollectGarbage(s string) int{
	s = removeCancels(s)

	ctr := 0 
	inGarbage := false

	for _, c := range s {

		if inGarbage{
			ctr++
		}

		//are we moving in or out garbage ?
		if c == rune('<') {
			inGarbage = true
		} else if c == rune('>') {
			inGarbage = false
			ctr-- //correct
		}
	}

	return ctr
}


func removeCancels(s string) string{
	res := ""
	for i := 0; i < len(s) ; i++{
		c := s[i]

		if c == byte('!') {
			i++ //skip 
		} else {
			res += string(c)
		}

	}

	return res
}
func readInput(fname string) *string {
	s, _ := ioutil.ReadFile(fname)
	res := string(s)
	return &res
}