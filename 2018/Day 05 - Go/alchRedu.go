package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func readInput(fname string) string{
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}
func pass(s string) string{
	//loop over string remove as much possible occurences at once
	res := ""
	i, j := 0, 1
	for ; j < len(s) ; i, j = i+1, j+1 {
		if strings.EqualFold(string(s[i]), string(s[j])) && s[i] != s[j] {
			//skip one char
			i++
			j++
			continue
		}
		//i possibly changed
		if i < len(s) { res += string(s[i]) }
	}
	//last char of string
	if i < len(s){ res += string(s[i]) }
	
	return res
}
func reduce(s string) string{
	res := pass(s)
	for len(res) < len(s) {
		s = res
		res = pass(res)
	}
	return res
}
func rereduce(poly string) int{
	min := len(poly)
	for _, c := range "abcdefghijklmnopqrstuvwxyz"{
		s := strings.Replace(poly, string(c), "", -1)
		s  = strings.Replace(s, strings.ToUpper(string(c)), "", -1)

		r := len(reduce(s))
		if r < min {
			min = r
		}
	}
	return min
}

func main(){

	//Part 1
	input := readInput("input.txt")
	fmt.Println(
		"Part 1:", 
		len(reduce(input)),
	)

	//Part 2
	input = readInput("input.txt")
	fmt.Println("Part 2:", rereduce(input))
}