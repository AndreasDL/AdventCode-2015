package main

import (
	"strings"
	"fmt"
	"io/ioutil"
)

func isNice( s string) bool {

	if strings.Contains(s, "ab") || strings.Contains(s, "cd") || 
	   strings.Contains(s, "pq") || strings.Contains(s, "xy") {
		return false
	}
	
	vowels := 0
	for i := 0 ; i < len(s) && vowels < 3; i++ {
		if strings.Contains("aeiou", string(s[i])) { vowels++ }
	}

	appearTwice := false 
	for i , j := 0, 1 ; j < len(s) && !appearTwice ; i,j = i+1, j+1 {
		if s[i]==s[j] { appearTwice = true }
	}

	return appearTwice && vowels >= 3
}


func main() {
	ctr := 0
	input := readInput("input.txt")
	for _, line := range strings.Split(input, "\n") {
		if isNice(line) { ctr++ }
	}
	fmt.Println(ctr)
}


func readInput (fname string) string{
	res , _ := ioutil.ReadFile(fname)
	return string(res)
}