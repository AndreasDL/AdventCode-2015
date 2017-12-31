package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	ctr1, ctr2 := 0, 0
	input := readInput("input.txt")
	for _, line := range strings.Split(input, "\n") {
		if isNice1(line) { ctr1++ }
		if isNice2(line) { ctr2++ }

	}
	fmt.Println(ctr1, ctr2)

	//Part 2 doesn't work => regexp doesn't support referral of named groups yet
	//switching to perl
	fmt.Println( isNice2("qjhvhtzxzqqjkmpb") )
	fmt.Println( isNice2("xxyxx") )
	fmt.Println( isNice2("uurcxstgmygtbstg") )
	fmt.Println( isNice2("ieodomkazucvgmuy") )

}

func readInput (fname string) string{
	res , _ := ioutil.ReadFile(fname)
	return string(res)
}
func isNice1(s string) bool {

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
func isNice2(s string) bool{
	fmt.Print(s, " => ")

	r1 := regexp.MustCompile("(?P<foo>.).*(?P=foo)") //not supported!
	//as are backregerences etc!
	//https://github.com/google/re2/wiki/Syntax => check later
	oneLetterRepeat := r1.MatchString(s)

/*
	r2 := regexp.MustCompile("(?P<first>..).*(?P<first>)")
	twiceTwice := r2.MatchString(s)
*/

	return oneLetterRepeat //&& twiceTwice
}


