package main


import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readInput(fname string) *[]string {
	s, _ := ioutil.ReadFile(fname)
	res := strings.Split(
		string(s), 
		"\n",
	)

	return &res
}

//Part 1
func countRunes(s *string) *map[rune]int{
	str := *s
	counts := make(map[rune]int)
	for _, c := range str {
		counts[c] += 1
	}

	return &counts
}
func isDouble(s *string) bool{

	counts := *countRunes(s)

	for _, v := range counts{
		if v == 2 {
			return true
		}
	}

	return false
}
func isTriple(s *string) bool{
	
	counts := *countRunes(s)
	for _, v := range counts {
		if v == 3 {
			return true
		}
	}
	return false
}
func calcChecksum(lines *[]string)int {
	
	doubles, triples := 0, 0
	for _, line := range *lines{

		if isDouble(&line) { doubles++ }
		if isTriple(&line) { triples++ }

	}

	return doubles * triples
}

//Part 2
func getDistance(a, b *string) (int, string) {
	
	//assumption both strings are of equal length
	dist := 0
	common_letters := ""
	for i := 0; i < len(*a); i++ {
		
		ca := (*a)[i] //deref & get rune
		cb := (*b)[i] 

		if cb != ca { 
			dist++ 
		} else {
			common_letters += string(ca)
		}
	}
	return dist, common_letters
}
func getLetters(lines *[]string)string{
	inputs := *lines
	
	for i := 0 ; i < len(inputs); i++ {
		for j := i+1; j < len(inputs); j++ {

			dist, letters := getDistance(
				&inputs[i],
				&inputs[j],
			)

			if dist == 1 {
				return letters
			}
		}
	}
	return ""
}

func main(){

	lines := readInput("input.txt")

	//part 1
	checksum := calcChecksum(lines)
	fmt.Println(checksum)

	//part 2
	letters := getLetters(lines)
	fmt.Println(letters)
}