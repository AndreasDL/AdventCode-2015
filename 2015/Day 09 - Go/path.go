package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	pq "github.com/Workiva/go-datastructures/queue" //modified to use item.String() => hashable maps
)

func main() {
	chart := newChart(readInput("input.txt"))
	fmt.Println(chart)

	shortest := chart.findShortest()
	fmt.Println(shortest.path, shortest.length)

	longest := chart.findLongest()
	fmt.Println(longest.path, longest.length)
}

type Chart struct {
	//cities[from][to] = distance
	cities map[string]map[string]int
}
func newChart(input string) *Chart{
	res := map[string]map[string]int{} //res[from][to] = distance
	for _, line := range strings.Split(input, "\n"){
		fields := strings.Fields(line)

		from        := fields[0]
		to          := fields[2]
		distance, _ := strconv.Atoi(fields[4])

		//init maps if needed
		if _, ex := res[from] ; ! ex { res[from] = map[string]int{} }
		if _, ex := res[to]   ; ! ex { res[to]   = map[string]int{} }

		res[from][to] = distance
		res[to][from] = distance
	}

	return &Chart{
		cities: res,
	}
}
func (c Chart) findShortest() Path {
	//Assumption: solution doesn't contain loops => requirement for part 2

	//init states
	todos := pq.NewPriorityQueue(50, false)
	for k, _ := range c.cities{
		todos.Put(Path{
			visited: map[string]bool{k: true},
			length : 0,
			path   : []string{k},
		})
	}
	fmt.Println(todos)

	//search
	for !todos.Empty() {

		item, err := todos.Get(1)
		if err != nil { panic(nil) }
		curr, ok := item[0].(Path)
		if !ok { 
			panic("we should be able to cast to Path") 
		} else if c.IsFinal(curr){
			return curr
		}

		for _, p := range *(c.NextPaths(curr)) {
			todos.Put(p)
		}
	}

	return Path{}
}
func (c Chart) findLongest() Path {
	//Requirement: solution doesn't contain loops
	//koop looping over all paths!

	//init states
	todos := pq.NewPriorityQueue(50, false)
	for k, _ := range c.cities{
		todos.Put(Path{
			visited: map[string]bool{k: true},
			length : 0,
			path   : []string{k},
		})
	}
	fmt.Println(todos)

	//search
	var max Path
	for !todos.Empty() {

		item, err := todos.Get(1)
		if err != nil { panic(nil) }
		curr, ok := item[0].(Path)
		if !ok { 
			panic("we should be able to cast to Path") 
		} else if c.IsFinal(curr) && curr.length > max.length{
			max = curr
		}

		for _, p := range *(c.NextPaths(curr)) {
			todos.Put(p)
		}
	}

	return max
}

func (c Chart) IsFinal(p Path) bool {
	//final when we have visited all cities
	return len(p.visited) == len(c.cities)
}
func (c Chart) NextPaths(p Path) *[]Path {

	res := []Path{}
	for nextCity, dist := range c.cities[p.currCity()] {
		if _, ex := p.visited[nextCity] ; !ex { //don't go back to visited cities!

			newPath := make([]string, len(p.path), cap(p.path))
			copy(newPath, p.path)
			newPath = append(newPath, nextCity)

			newVisited := map[string]bool{}
			for k, v := range p.visited { newVisited[k] = v }
			newVisited[nextCity] = true

			res = append(res, Path{
				length  : p.length + dist,
				path    : newPath,
				visited : newVisited,
			})

		}
	}
	return &res
}

type Path struct{
	visited map[string]bool
	length int
	path []string
}
func (p Path) Compare(pp pq.Item)int {

	pother, ok := pp.(Path)
	if !ok { return 0 }

	if p.length > pother.length { 
		return 1 
	} else if p.length < pother.length {
		return -1
	}
	return 0
}
func (p Path) currCity() string{
	//current city => last visited city
	return p.path[len(p.path)-1]
}
func (p Path) String() string{
	//used to enable hashing of Path object in modified priority queue
	res := strings.Join(p.path, "->")
	res += strconv.Itoa(p.length) + "|"
	for k, v := range p.visited {
		if v { res += k + "|" }
	}
	return res
}

func readInput(fname string) string {
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}