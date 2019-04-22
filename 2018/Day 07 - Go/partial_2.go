package main


import (
	"fmt"
	"sort"
	"io/ioutil"
	"strings"
)


//task
type task struct{
	name string
	
	requirements []string
	next []string
	
	finished bool //done
	started bool //work in progress
	scheduled bool //in options list
	
	end int
}
func (t task) canStart(_tasks *map[string]*task) bool{
	tasks := *_tasks

	if t.started { 
		//already running
		return false
	} else if len(t.requirements) == 0 { 
		//root node
		return true 
	}
	//check requirements
	for _, r := range t.requirements {
		if ! (*tasks[r]).finished { return false }
	}
	return true
}
func (t *task) finishWhenDone(time int) bool{

	if time >= t.end { 
		t.finished = true
	}

	return t.finished
}
func (t *task) start(time int){
	t.started = true
	t.end = time + int(t.name[0] - 'A' + 1) + ADD_TASK_TIME
}

//IO / input
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
				bef, []string{}, []string{aft}, false, false, false, 0,
			}
		} else {
			(*tasks[bef]).next = append( (*tasks[bef]).next, aft)
		}

		//after node
		if _, exists := tasks[aft]; !exists {
			tasks[aft] = &task{
				aft, []string{bef}, []string{}, false, false, false, 0,
			}
		} else {
			(*tasks[aft]).requirements = append( (*tasks[aft]).requirements, bef)
		}
	}

	return &tasks, root
}


//search related
func getNextSteps(_tasks *map[string]*task, time int) []*task{
	tasks := *_tasks

	possible_tasks := []*task{}
	for _, task := range tasks {
		if (*task).canStart(_tasks){
			possible_tasks = append(possible_tasks, task)
		}
	}

	sort.Slice(possible_tasks, func(i, j int) bool {
	  return (*possible_tasks[i]).name < (*possible_tasks[i]).name
	})

	return possible_tasks
}
func getPath(_tasks *map[string]*task, root string) (string, int) {

	workplaces := [CAPACITY]*task{}

	path := ""
	time := 0; for ; len(path) < len(*_tasks); time++ {

		//fmt.Println(time, "start", workplaces)

		//finish completed tasks && free up space
		for i, _task := range workplaces {
			if _task == nil { continue }

			if (*_task).finishWhenDone(time){
				workplaces[i] = nil
				path += (*_task).name
			}
		}

		//get next steps
		possible_steps := getNextSteps(_tasks, time)
		//fmt.Println(time, "cleaned and prioritised", workplaces, possible_steps)

		if len(possible_steps) == 0 { continue }

		//schedule
		for i, _spot := range workplaces {
			
			if _spot != nil { continue } //some other task already running
			
			//if we have an empty spot start another task
			workplaces[i] = possible_steps[0]
			(possible_steps[0]).start(time)

			//remove task from possible_steps
			possible_steps = possible_steps[1:]

			if len(possible_steps) == 0 { break }
		}

		//fmt.Println(time, "scheduled", workplaces, possible_steps)
		//fmt.Println()
	}

	//for loop will do time++ before checking the len(path) < len(*_tasks) so we need to correct for that
	time-- 
	return path, time
}


const CAPACITY int = 2 //number of workers, 2 in example, 5 in ex
const ADD_TASK_TIME int = 0 //required additional time per time 0 in example, 60 in ex

func main(){
	input := &[]string{
		"Step C must be finished before step A can begin.",
		"Step C must be finished before step F can begin.",
		"Step A must be finished before step B can begin.",
		"Step A must be finished before step D can begin.",
		"Step B must be finished before step E can begin.",
		"Step D must be finished before step E can begin.",
		"Step F must be finished before step E can begin.",
	}

	//input := readInput("input.txt")
	tasks, root := parseInput(input)
	path, duration := getPath(tasks, root)
	fmt.Println("Part 2:", path, duration)
}