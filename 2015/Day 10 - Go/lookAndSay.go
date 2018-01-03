package main

import (
	"strconv"
	"fmt"
)

func next(s string) string{
	res := ""
	j := 1
	for i := 0 ; i < len(s) ; i+= j-i {
		j = i+1 ; for ; j < len(s) && s[j]==s[i] ; j++ {}
		
		res += strconv.Itoa(j-i) + string(s[i])
	}

	return res
}

func main() {
	/*fmt.Println(next("1"))
	fmt.Println(next("11"))
	fmt.Println(next("21"))
	fmt.Println(next("1211"))
	fmt.Println(next("111221"))
	fmt.Println(next("312211"))*/

	s := "1113222113"
	for i := 0 ; i < 40 ; i++ { s = next(s) }
	fmt.Println(len(s))

	//takes some time
	for i := 0 ; i < 10 ; i++ { s = next(s) ; fmt.Println(i, len(s)) }
	fmt.Println(len(s))
}