package main

import (
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
	"sort"
)

func main() {
	input := readInput("input.txt")
	
	totalSur, totalRib := 0, 0
	for _, line := range strings.Split(input, "\n"){
		totalSur += Surface(line)
		totalRib += Ribbon(line)
	}
	
	fmt.Println(totalSur, totalRib)
}


func Surface(s string) int{
	l,w,h := parseLine(s)
	lw := l*w
	wh := w*h
	hl := h*l
	return 2*(lw+wh+hl) + Min(lw,wh,hl)
}
func Min(a,b,c int) int{
	if a < b && a < c { 
		return a 
	} else if b < c {
		return b
	}
	return c
}
func parseLine(s string) (int, int, int){
	parts := strings.Split(s, "x")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	c, _ := strconv.Atoi(parts[2])
	return a,b,c
}

func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}

func Ribbon(s string) int{
	l,w,h := parseLine(s)
	bow := l*w*h

	sides := []int{l,w,h}
	sort.Ints(sides)

	wrap := sides[0]+sides[0]+sides[1]+sides[1]

	return wrap+bow
}