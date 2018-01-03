package main

import (
	"strings"
	"fmt"
)

func isValid(s string) bool {
	//does not contain i o l
	if strings.ContainsAny(s, "iol") { return false }

	//two different pairs
	b, e := 0, 1 ; for ; e < len(s) && s[b]!=s[e]; b,e = b+1, e+1 {}
	if e >= len(s) { 
		return false 
	} else { //repeat starting with adjusted positions
		b+=2
		e+=2
		for ; e < len(s) && s[b]!=s[e]; b,e = b+1, e+1 {}
		if e >= len(s) {return false }
	}

	//Straight of three letters
	straight := false
	for i,j,k := 0, 1, 2 ; k < len(s) && !straight ; i,j,k = i+1,j+1,k+1 {

		c := s[i]
		straight = (s[j] == c+1 && s[k] == c+2)
	}

	return straight
}

func increment(s string) string {
	pass := []byte(s)
	for i, _ := range pass { pass[i] -= 'a' }

	//increment
	pass[len(pass)-1]++

	//fix wraparound
	for i := len(pass)-1 ; i > 0 ; i-- {
		pass[i-1] += byte(pass[i] / 26)
		pass[i] %= 26
	}

	for i, _ := range pass { pass[i] += 'a' }
	return string(pass)
}

func main() {
	start := increment("abcdefgh") ; for !isValid(start) { start = increment(start) }
	fmt.Println("Sample1: ", start)

	start  = increment("ghijklmn") ; for !isValid(start) { start = increment(start) }
	fmt.Println("Sample2: ", start)

	start  = increment("hepxcrrq") ; for !isValid(start) { start = increment(start) }
	fmt.Println("Part1: ", start)

	start  = increment(start)      ; for !isValid(start) { start = increment(start) }
	fmt.Println("Part2: ", start)
}