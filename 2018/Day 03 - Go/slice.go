package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"io/ioutil"
)
var regex_split = regexp.MustCompile("[#@,x:\\s]+")

func parseLine(line string) (int, int, int, int, int){
	
    parts := regex_split.Split(line, -1)

    id, _ := strconv.Atoi(parts[1])
    x , _ := strconv.Atoi(parts[2])
    y , _ := strconv.Atoi(parts[3])
    dx, _ := strconv.Atoi(parts[4])
    dy, _ := strconv.Atoi(parts[5])

    return id, x, y, dx, dy
}
func readLines(fname string) []string{
	file, _ := ioutil.ReadFile(fname)

	lines := strings.Split(
		string(file),
		"\n",
	)

	return lines
}

//Part 1
func handleLine(patch *[1000][1000]int, line string){
	
	_, x, y, dx, dy := parseLine(line)
	for i := x ; i < x+dx ; i++ {
		for j := y; j < y+dy ; j++{
			patch[j][i]++
		}
	}
}
func count(patch *[1000][1000]int) int{
	
	cnt := 0
	patch_obj := *patch
	for _, line := range patch_obj{
		for _, val := range line{
	
			if val > 1 { cnt++ }
		}
	}

	return cnt
}

//Part 2
func checkLine(patch *[1000][1000]int, line string) (bool, int) {
	id, x, y, dx, dy := parseLine(line)
	for i := x ; i < x+dx ; i++ {
		for j := y; j < y+dy ; j++{
			
			if patch[j][i] > 1 {
				return false, -1
			}
		}
	}

	return true, id
}

func main(){
	
	patch := [1000][1000]int{}
	lines := readLines("input.txt")

	for _, line := range lines {
		handleLine(&patch, line)
	}
	
	fmt.Println("Part 1:", count(&patch))

	for _, line := range lines {

		if found, id := checkLine(&patch, line) ; found {
			fmt.Println("Part 2:", id)
			break
		}

	}

}