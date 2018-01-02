package main

import (
	"io/ioutil"
	"fmt"
	"strings"
)


func main() {

	input := readInput("input.txt")//*/"sampleInput.txt")


	cLength, sLength, sLength2 := 0, 0, 0
	for _, line := range strings.Split(input, "\n"){
		cLength  += len(line)
		sLength  += strLength(line)
		sLength2 += strLength2(line)

		fmt.Println(strLength2(line))
	}

	fmt.Println(cLength, sLength, " => ", cLength - sLength,  " | ", sLength2, " => ", sLength2 - cLength)
}


func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}

func strLength(s string) int {
	ctr := 0
	for i := 1 ; i < len(s)-1 ; i++ {

		ctr++

		if s[i] == '\\'{
			switch s[i+1] {
				case 'x' : i += 3
				case '"' : i++
				case '\\': i++
			}
		}
	}
	return ctr
}

func strLength2(s string) int {
	ctr := 0
	for i := 0 ; i < len(s) ; i++ {

		ctr++

		if s[i] == '"' { 
			ctr++ 
		} else if s[i] == '\\'{
			ctr++
		} 

	}
	return ctr +2 
}