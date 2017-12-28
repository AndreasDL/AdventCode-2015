package main

import (
	"io/ioutil"
	"fmt"
)

var deltas = map[byte]struct{dx,dy int}{
	'^': { 0,-1},
	'V': { 0, 1},
	'<': {-1, 0},
	'>': { 1, 0},

	'v': { 0, 1},
}

func main() {
	input := readInput("input.txt")

	fmt.Println(Visited(input))
}


func Visited(s string) int {
	beenThere := map[string]bool{"0;0": true}

	x, y := 0, 0 
	for _, c := range []byte(s){
		delta := deltas[c]

		x += delta.dx
		y += delta.dy

		key := fmt.Sprintf("%d;%d",x,y)
		beenThere[key] = true

		//fmt.Println(string(c), key, beenThere)
	}

	return len(beenThere)
}


func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}