package main


import (
	"fmt"
	"io/ioutil"
)

type Point struct{
	x,y int
}
func (n Point) north() Point {
	return Point{n.x, n.y-1}
}
func (n Point) east() Point {
	return Point{n.x+1,n.y}
}
func (n Point) south() Point {
	return Point{n.x, n.y+1}
}
func (n Point) west() Point {
	return Point{n.x-1,n.y}
}


func buildGraph(input string) (Point, *map[Point]map[Point]bool) {

	root := Point{ 0,0 }
	graph := map[Point]map[Point]bool{
		root: map[Point]bool{},
	}
	current_roots := []Point{ root }

	for _, c := range input {

		v := current_roots[len(current_roots)-1]

		//fmt.Println(string(c), v, current_roots, graph)
		if c == 'N' || c == 'E' || c == 'S' || c == 'W' {
		
			var w Point
			if c == 'N' {
				w = v.north()
			} else if c == 'E' {
				w = v.east()
			} else if c == 'S' {
				w = v.south()
			} else if c == 'W' {
				w = v.west()
			}

			if _, ex := graph[v]; !ex { graph[v] = map[Point]bool{} }
			graph[v][w] = true

			//pop()
			current_roots = current_roots[:len(current_roots)-1]
			//add 
			current_roots = append(current_roots, w)
			
		} else if c == '(' {
			current_roots = append(current_roots, v)
		} else if c == ')' {
			current_roots = current_roots[:len(current_roots)-1]
		} else if c == '|' {
			current_roots = current_roots[:len(current_roots)-1]
			current_roots = append(current_roots, current_roots[len(current_roots)-1])
		}
	}

	return root, &graph
}


func bfs(root Point, graph *map[Point]map[Point]bool) map[Point]int{
	_graph := *graph

	seen := map[Point]bool{ root: true }
	distances := map[Point]int{ root: 0 } //keep track of shortest distances
	roots := map[Point]Point{} //needed to get distance for new nodes
	
	//init todo so we don't need to 'hack' init of roots and distances for root
	todo := []Point{}
	for k, _ := range _graph[root] {
		//skip seen entries
		if s, ex := seen[k]; s || ex { continue }		
		todo = append(todo, k)
		roots[k] = root
	}


	for len(todo) > 0 {

		//todo.pop()
		v := todo[0]
		todo = todo[1:]

		//fmt.Println(v, seen, distances)

		//skip seen entries => Breadth first means we won't ever find a shorter path to current node, since breadth first goes from short paths to long paths
		if s, ex := seen[v] ; s || ex { continue }

		//seen
		seen[v] = true

		//distances
		root := roots[v]
		distances[v] = distances[root] + 1

		//roots
		for k, _ := range _graph[v] {
			//skip seen entries
			if s, ex := seen[k]; s || ex { continue }
			
			todo = append(todo, k)
			roots[k] = v
		}
	}

	return distances
}

func part1(input string) int {
	root, graph := buildGraph(input)
	distances := bfs(root, graph)

	max := -1
	for _, v := range distances{
		if v > max { max = v}
	}

	return max
}

func part2(input string, threshold int) int{
	root, graph := buildGraph(input)
	distances := bfs(root, graph)

	cnt := 0
	for _, v := range distances{
		if v >= threshold { cnt++ }
	}

	return cnt
}


func main() {


	//minimalistic testing framework!
	if part1("WNE") != 3 { panic("not working")}
	if part1("ENWWW(NEEE|SSE(EE|N)") != 10 { panic("not working")}
	if part1("ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN") != 18 { panic("not working")}
	if part1("ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))") != 23 { panic("not working")}
	if part1("WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))") != 31 { panic("not working")}

	file , _ := ioutil.ReadFile("input.txt")
	input := string(file)

	fmt.Println("Part1:", part1(input))
	fmt.Println("Part2:", part2(input, 1000))

}