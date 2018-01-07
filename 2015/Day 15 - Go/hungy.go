package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

type Ingredient struct {
	capacity    int
	duratbility int
	flavor      int
	texture  	int
	calories 	int
}
func readInput(fname string) string {
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}
func parseInput(input string) []Ingredient {

	lines := strings.Split(input, "\n")
	res   := make([]Ingredient, len(lines))

	for i, line := range lines {

		line = strings.Replace(line, ",", "", -1)
		fields := strings.Fields(line)

		cap, _ := strconv.Atoi(fields[ 2])
		dur, _ := strconv.Atoi(fields[ 4])
		fla, _ := strconv.Atoi(fields[ 6])
		tex, _ := strconv.Atoi(fields[ 8])
		cal, _ := strconv.Atoi(fields[10])

		res[i] = Ingredient{
			capacity   : cap,
			duratbility: dur,
			flavor     : fla,
			texture    : tex,
			calories   : cal,
		}
	}

	return res
}

func main() {
	input := readInput("input.txt")

	ingredients := parseInput(input)

	fmt.Println( OptimalBalance(ingredients) )

}


func OptimalBalance(ingredients []Ingredient) int{

	maxScore := 0
	for a := 0 ; a < 100 ; a++ {
		for b := 0 ; b < 100 ; b++ {
			for c := 0 ; c < 100 ; c++ {

				if a+b+c > 100 { continue }
				d := 100 -a-b-c

				cap := a * ingredients[0].capacity
				cap += b * ingredients[1].capacity
				cap += c * ingredients[2].capacity
				cap += d * ingredients[3].capacity
				if cap <= 0 { continue }

				dur := a * ingredients[0].duratbility
				dur += b * ingredients[1].duratbility
				dur += c * ingredients[2].duratbility
				dur += d * ingredients[3].duratbility
				if dur <= 0 { continue }

				fla := a * ingredients[0].flavor
				fla += b * ingredients[1].flavor
				fla += c * ingredients[2].flavor
				fla += d * ingredients[3].flavor
				if fla <= 0 { continue }

				tex := a * ingredients[0].texture
				tex += b * ingredients[1].texture
				tex += c * ingredients[2].texture
				tex += d * ingredients[3].texture
				if tex <= 0 { continue }

				score := cap * dur * fla * tex

				if score > maxScore { maxScore = score }
			}
		}
	}

	return maxScore	
}