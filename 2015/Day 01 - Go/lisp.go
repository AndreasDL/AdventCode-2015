package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	input := readInput("input.txt")
	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}


func Part1(s string) int{
	pos := 0 
	for _, c := range []byte(s){
		if c == '(' { 
			pos++ 
		} else if c == ')'{
			pos--
		}
	}
	return pos
}

func Part2(s string) int{
	pos := 0
	for i, c := range []byte(s){
		if c == '(' { 
			pos++ 
		} else if c == ')'{
			pos--
		}

		if pos < 0 {
			return i+1
		}
	}
	
	return -1
}


func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}