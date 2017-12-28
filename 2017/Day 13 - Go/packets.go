package main

import (
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
	"time"
)

func main(){
	input := readInput("input.txt")

	fmt.Println( Part1(input) )

	fmt.Println( Part2(input) )
}

func Part1(input string) int{
	firewall,length := parseInput(input)

	return Severity(firewall, length, 0)
}
func Severity(firewall map[int]int, length int, delay int) int {
	severity := 0
	for i := 0; i <= length; i++ {
		

		layerDepth := firewall[i]
		if layerDepth == 0 { continue }

		scannerPos := (i+delay) % layerDepth
		if scannerPos == 0 {
			severity += i * (layerDepth+2)/2
		}
	}

	return severity
}

func Part2(input string) int{
	defer track(time.Now(),"Part 2")
	firewall, length := parseInput(input)

	i := 0
	caught := GetCaught(firewall, length, i)
	for caught {

		i++
		caught = GetCaught(firewall, length, i)
	}

	return i
}
func GetCaught(firewall map[int]int, length int, delay int) bool{
	
	for i := 0; i <= length; i++ {
		
		layerDepth := firewall[i]
		if layerDepth == 0 { continue }

		scannerPos := (i+delay) % layerDepth
		if scannerPos == 0 {
			return true
		}
	}

	return false
}

func parseInput(s string) (map[int]int, int){
	s = strings.Replace(s, " ", "", -1)
	max := -1

	firewall := map[int]int{}
	for _, line := range strings.Split(s, "\n") {
		parts := strings.Split(line, ":")
		
		i  , _ := strconv.Atoi(parts[0])
		val, _ := strconv.Atoi(parts[1])

		firewall[i]= val + val - 2

		if val > max { max = val }
		if i   > max { max = i }
	}

	return firewall, max
}
func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}
func track(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf(" %s took %s \n", name, elapsed)
}