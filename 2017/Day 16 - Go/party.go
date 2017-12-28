package main

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
)

func Spin(input string, x int) string{
	res := make([]rune, len(input))
	
	for b, i := 0, len(input) -x ; i < len(input) ; b, i = b+1, i+1 { 
		res[b] = rune(input[i])
	}

	for i := 0 ; i < len(input) -x ; i++{
		res[i+x] = rune(input[i])
	}

	return string(res)
}
func Exchange(input string, a,b int) string{
	res := []rune(input)
	res[a], res[b] = res[b], res[a]
	return string(res)
}
func Partner(input string, a,b string) string{
	posA, posB := getPos(input, a), getPos(input, b)
	return Exchange(input, posA, posB)
}
func getPos(input, a string) int{
	i := 0
	for i < len(input) && input[i] != a[0] { i++ }
	return i
}

func ParseAndRun(cmd, input string) string {
	switch rune(cmd[0]) {
	case 's':
		x, _ := strconv.Atoi(cmd[1:])
		return Spin(input, x)
	case 'x':
		parts := strings.Split(cmd, "/")
		a, _ := strconv.Atoi(parts[0][1:])
		b, _ := strconv.Atoi(parts[1])
		return Exchange(input, a,b)
	case 'p':
		parts := strings.Split(cmd, "/")
		a,b := parts[0][1:], parts[1]
		return Partner(input, a,b)
	}

	return input
}

func main(){
	start := "abcde"
	
	start = ParseAndRun("s1", start)
	fmt.Println(start)
	start = ParseAndRun("x3/4", start)
	fmt.Println(start)
	start = ParseAndRun("pe/b", start)
	fmt.Println(start)
	fmt.Println()
	fmt.Println()


	input := readInput("input.txt")
	start  = "abcdefghijklmnop"
	for _, instr := range strings.Split(input, ","){
		start = ParseAndRun(instr, start)
	}
	fmt.Println(start)

}
func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}