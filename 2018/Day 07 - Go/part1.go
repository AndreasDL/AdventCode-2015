package main


import (
	"fmt"
	"sort"
	"io/ioutil"
	"strings"
)

func readInput(fname string) *[]string{
	file, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(file), "\n")
	return &lines
}
func parseLine(line string) (string, string) {
	return string(line[5]), string(line[36])
}
func buildTrees(_lines *[]string) (*map[string][]string, *map[string][]string, string) {
	lines := *_lines

	//resulting map gives the possible paths at a node
	//also keep track of the root / starting point => starting point should never appear as a before step X can begin
	root, _ := parseLine(lines[0])
	after   := map[string][]string{}
	before  := map[string][]string{}
	for _, line := range lines {

		bef, aft := parseLine(line)

		//build tree
		after[bef]  = append(after[bef] , aft)
		before[aft] = append(before[aft], bef)

		if aft == root {
			root = bef
		}
	}
	return &after, &before, root
}
func isPossible(_before *map[string][]string, _seen *map[string]bool, node string) bool{
	before := *_before
	seen := *_seen

	//root note
	if _, exists := before[node]; !exists { return true }

	//else check if all element before node are already seen
	for _, n := range before[node] {
		if _, exists := seen[n]; !exists {
			return false
		}
	}
	return true
}
func popNext(_possibilities *[]string, _before *map[string][]string, _seen *map[string]bool) string{
	possibilities := *_possibilities

	//sort list => always pick one that is first in alphabetical order
	sort.Strings(possibilities)

	//find first one that is possible
	for i, p := range possibilities {
		if isPossible(_before, _seen, p){
			//remove elemnet from list => make sure we overwrite pointer!
			*_possibilities = append(possibilities[:i], possibilities[i+1:]...)
			//return the element
			return p
		}
	}
	//this shouldn't happen!
	panic("no next step!")
}
func getPath(_after *map[string][]string, _before *map[string][]string, root string)string {
	after := *_after

	//best first search of path
	path := ""
	seen := map[string]bool{}
	possibilities := []string{root}
	for len(possibilities) > 0 {

		//pop next possible element from list
		next := popNext(&possibilities,_before,&seen)
		
		//skip elements that have been seen before
		if _, exists := seen[next]; exists { continue }

		//add path
		path += next
		//fast lookup of previous steps
		seen[next] = true

		//add paths to possibilities
		for _, t := range after[next]{
			possibilities = append(possibilities, t)
		}

		//fmt.Println(path, next, seen, possibilities)
	}

	return path
}


func main(){

	/*
	input := []string{
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin.",
	}
	after, before, root := buildTrees(&input)
	fmt.Println("Part 1:", getPath(after, before, root))
	*/

	input := readInput("input.txt")
	after, before, root := buildTrees(input)
	fmt.Println("Part 1:", getPath(after, before, root))
}