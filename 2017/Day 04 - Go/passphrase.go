package main

import (
	"fmt"
	"strings"
	"sort"
	"ioutil"
)

func main(){
	intput := ioutil.ReadFile("input.txt")
	fmt.Println("aa bb cc dd ee", "->", isValid("aa bb cc dd ee"))
	fmt.Println("aa bb cc dd aa", "->", isValid("aa bb cc dd aa"))
	fmt.Println("aa bb cc dd aaa", "->", isValid("aa bb cc dd aaa"))

	ctr := 0
	for _, line := range strings.Split(input, "\n"){

		if isValid(line){ ctr++	}
	}

	fmt.Println(ctr)


	//part 2
	fmt.Println("aa bb cc dd ee" , "->", isValid2("aa bb cc dd ee"))
	fmt.Println("aa bb cc dd aa" , "->", isValid2("aa bb cc dd aa"))
	fmt.Println("ab ba cc dd aaa", "->", isValid2("ab ba cc dd aaa"))

	ctr = 0
	for _, line := range strings.Split(input, "\n"){

		if isValid2(line){ ctr++ }
	}

	fmt.Println(ctr)

}


func isValid(s string) bool {
	
	seen := map[string]int{}
	for _, s := range strings.Split(s, " "){
		seen[s]++

		if seen[s] > 1 {
			return false
		}
	}

	return true
}

func isValid2(str string) bool {
	seen := map[string]int{}
	for _, s := range strings.Split(str, " "){

		sorted := strings.Split(s, "")
		sort.Strings( sorted )

		key := strings.Join(sorted, "")
//		fmt.Println(s , " => ", sorted, " => ", key)
		
		seen[key]++

		if seen[key] > 1 { return false	}
	}

	return true
}