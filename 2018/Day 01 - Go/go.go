package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)


func parseInstructions(input string) *[]int{
	res := strings.Split(input, "\n")

	result := make([]int, len(res), len(res)) //init ti right size

	for i, r := range res{
		result[i], _ = strconv.Atoi(r)
	}

	return &result
}
func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}

func main(){
	
	steps := *parseInstructions(
		readInput("input.txt"),
	)

	freq := 0
	found := false
	seen := make(map[int]bool)
	for i := 0; !found; i++ { //loop while not found, keeping track of the iteration
		for _, step := range steps {

			freq += step
			
			if _, exists := seen[freq]; exists{
				fmt.Println("part 2:", freq)
				found = true
				if i > 0 { break }
			}

			seen[freq] = true
		}

		if i == 0 { //part1 => end of first iteration
			fmt.Println("part 1:", freq)
		}
	}

}