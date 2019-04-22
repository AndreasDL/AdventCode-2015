package main

import (
	"fmt"
	"strconv"
	"regexp"
	"strings"
	"io/ioutil"
)

type Nanobot struct{
	x, y, z int
	r int
}
func (n Nanobot) inRange(nn *Nanobot)bool {

	dist := abs(n.x - nn.x)
	dist += abs(n.y - nn.y)
	dist += abs(n.z - nn.z)

	return dist <= n.r
}

//helper functions
func toInt(s string) int{
	v, _ := strconv.Atoi(s)
	return v
}
func abs(i int) int {
	if i < 0 { return -i }
	return i
}
func readFile(fname string) string{
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}

//input parsing
var re = regexp.MustCompile("-?\\d+")
func parseLine(line string) *Nanobot{
	parts := re.FindAllString(line, 4)
	pi := make([]int, len(parts))
	for i, p := range parts{ pi[i] = toInt(p) }

	return &Nanobot{
		pi[0], pi[1], pi[2],
		pi[3],
	}
}
func parseInput(line string) []*Nanobot{
	lines := strings.Split(line, "\n")
	result := make([]*Nanobot, len(lines))
	for i, line := range lines {
		result[i] = parseLine(line)
	}
	return result
}

//part 1
func findStrongest(bots []*Nanobot) *Nanobot {
	max := bots[0]
	for i := 1; i < len(bots) ; i++ {
		if bots[i].r > max.r{
			max = bots[i]
		}
	}
	return max
}
func part1(bots []*Nanobot){
	
	strongest := findStrongest(bots)
	fmt.Println("strongest:", strongest)

	cnt := 0
	for _, b := range bots {
		if strongest.inRange(b) { cnt++ }
	}

	fmt.Println("Part1 : ", cnt)
}


func main() {
		
	input := readFile("input.txt")
	bots := parseInput(input)
	part1(bots)
	
}
