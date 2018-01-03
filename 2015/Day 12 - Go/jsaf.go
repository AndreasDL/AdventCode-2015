package main

import (
	"strconv"
	"regexp"
	"fmt"
	"io/ioutil"
)

func scanDigits(json string) int {
	//throw out all strings
	cr := regexp.MustCompile("\".+?\"")
	json = cr.ReplaceAllString(json, "")

	//find all numbers in the remaining
	r := regexp.MustCompile("-?[0-9]+")
	res := 0
	for _, s := range r.FindAllString(json, -1) {
		val, _ := strconv.Atoi(s)
		res += val
	}

	return res
}


func main() {
	fmt.Println(scanDigits(`{"a":2,"b":4}`))
	fmt.Println(scanDigits(`[1,2,3]`))

	input := readInput("input.txt")
	fmt.Println(scanDigits(input))
}

func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}