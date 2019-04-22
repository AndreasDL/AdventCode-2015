package main


import (
	"fmt"
	"sort"
	"io/ioutil"
	"strings"
)


type task struct{
	name string
	requirements []string
	next []string
	finished bool
}
func (t task) isPossible(_tasks *map[string]*task) bool{
	tasks := *_tasks

	//root node
	if len(t.requirements) == 0 { return true }

	//check requirements
	for _, r := range t.requirements {
		if ! (*tasks[r]).finished { return false }
	}
	return true
}


func readInput(fname string) *[]string{
	file, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(file), "\n")
	return &lines
}
func parseLine(line string) (string, string) {
	return string(line[5]), string(line[36])
}
func parseInput(_lines *[]string) (*map[string]*task, string){
	lines := *_lines

	root, _ := parseLine(lines[0])
	tasks := map[string]*task{}
	for _, line := range lines {

		bef, aft := parseLine(line)

		if aft == root {
			root = bef
		}

		//before node
		if _, exists := tasks[bef]; !exists {
			tasks[bef] = &task{
				bef, []string{}, []string{aft}, false,
			}
		} else {
			(*tasks[bef]).next = append( (*tasks[bef]).next, aft)
		}

		//after node
		if _, exists := tasks[aft]; !exists {
			tasks[aft] = &task{
				aft, []string{bef}, []string{}, false,
			}
		} else {
			(*tasks[aft]).requirements = append( (*tasks[aft]).requirements, bef)
		}
	}

	return &tasks, root
}


func popNext(_options *[]string, _tasks *map[string]*task) string{

	//sort list => always pick one that is first in alphabetical order
	sort.Strings( *_options )

	//find first one that is possible
	for i, p := range *_options {
		if (*_tasks)[p].isPossible(_tasks){
			//remove elemnet from list => make sure we overwrite pointer!
			*_options = append( (*_options)[:i], (*_options)[i+1:]...)
			//return the element
			return p
		}
	}

	//this shouldn't happen!
	panic("no next step!")
}
func getPath(_tasks *map[string]*task, root string)string {
	tasks := *_tasks

	//best first search of path
	path := ""
	options := []string{root}
	for len(options) > 0 {

		//pop next possible element from list
		next := popNext(&options, _tasks)
	
		//skip elements that have been seen before
		if (*tasks[next]).finished { continue }

		//add path
		path += next
		//fast lookup of previous steps
		(*tasks[next]).finished = true

		//add paths to options
		for _, t := range (*tasks[next]).next{
			options = append(options, t)
		}
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
	}*/

	input := readInput("input.txt")
	tasks, root := parseInput(input)
	fmt.Println("Part 1:", getPath(tasks, root))
}