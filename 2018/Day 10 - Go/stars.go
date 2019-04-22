package main


import (
	"fmt"
	"strings"
	"regexp"
	"strconv"
	"io/ioutil"
)

var regex = regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)



//node
type node struct{
	x, y int
	vx, vy int
}
func parseNode(line string) node{
	parts := regex.FindStringSubmatch(line)
	x , _ := strconv.Atoi(parts[1])
	y , _ := strconv.Atoi(parts[2])
	vx, _ := strconv.Atoi(parts[3])
	vy, _ := strconv.Atoi(parts[4])

	return node{x, y, vx, vy}
}
func (n *node) move(){
	n.x+= n.vx
	n.y+= n.vy
}
func (n *node) reverseMove(){
	n.x -= n.vx
	n.y -= n.vy
}

//all nodes
func parseNodes(input string) *[]node{
	lines := strings.Split(input, "\n")
	
	nodes := make([]node, len(lines))
	for i, line := range lines{
		nodes[i] = parseNode(line)
	}
	return &nodes
}
func doStep(_nodes *[]node) int{
	nodes := *_nodes
	minx, miny := nodes[0].x, nodes[0].y
	maxx, maxy := nodes[0].x, nodes[0].y
	for i, _ := range nodes {
		nodes[i].move()

		if nodes[i].x < minx { 
			minx = nodes[i].x
		} else if nodes[i].x > maxx {
			maxx = nodes[i].x 
		}

		if nodes[i].y < miny { 
			miny = nodes[i].y 
		} else if nodes[i].y > maxy {
			maxy = nodes[i].y 
		}

	}
	return (maxy - miny) * (maxx - minx)
}
func doReverseStep(_nodes *[]node){
	nodes := *_nodes
	for i, _ := range nodes {
		nodes[i].reverseMove()
	}
}
func min_max(nodes *[]node) (int, int, int, int){
	
	_nodes := *nodes
	minx, miny := _nodes[0].x, _nodes[0].y
	maxx, maxy := _nodes[0].x, _nodes[0].y
	for _, node := range _nodes{
		if node.x < minx { 
			minx = node.x
		} else if node.x > maxx {
			maxx = node.x 
		}

		if node.y < miny { 
			miny = node.y 
		} else if node.y > maxy {
			maxy = node.y 
		}
	}

	return minx, miny, maxx, maxy
}
func plot(nodes *[]node) string{

	minx, miny, maxx, maxy := min_max(nodes)

	//init
	field := make([][]byte, maxy - miny+1)
	for i, _ := range field{
		field[i] = make([]byte, maxx-minx+1)
		for j, _ := range field[i]{ field[i][j] = ' '}
	}

	//set values
	for _, node := range *nodes {
		y := node.y - miny
		x := node.x - minx
		field[y][x] = '#'
	}


	res := ""
	for _, line := range field {
		res += string(line) + "\n"
	}
	return res
}
func bb_size(nodes *[]node) int{
	minx, miny, maxx, maxy := min_max(nodes)

	size := (maxy - miny) * (maxx - minx)
	return size
}

//helper
func readFile(fname string) string{
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}


func main(){
	
	input := `position=< 9,  1> velocity=< 0,  2>
position=< 7,  0> velocity=<-1,  0>
position=< 3, -2> velocity=<-1,  1>
position=< 6, 10> velocity=<-2, -1>
position=< 2, -4> velocity=< 2,  2>
position=<-6, 10> velocity=< 2, -2>
position=< 1,  8> velocity=< 1, -1>
position=< 1,  7> velocity=< 1,  0>
position=<-3, 11> velocity=< 1, -2>
position=< 7,  6> velocity=<-1, -1>
position=<-2,  3> velocity=< 1,  0>
position=<-4,  3> velocity=< 2,  0>
position=<10, -3> velocity=<-1,  1>
position=< 5, 11> velocity=< 1, -2>
position=< 4,  7> velocity=< 0, -1>
position=< 8, -2> velocity=< 0,  1>
position=<15,  0> velocity=<-2,  0>
position=< 1,  6> velocity=< 1,  0>
position=< 8,  9> velocity=< 0, -1>
position=< 3,  3> velocity=<-1,  1>
position=< 0,  5> velocity=< 0, -1>
position=<-2,  2> velocity=< 2,  0>
position=< 5, -2> velocity=< 1,  2>
position=< 1,  4> velocity=< 2,  1>
position=<-2,  7> velocity=< 2, -2>
position=< 3,  6> velocity=<-1, -1>
position=< 5,  0> velocity=< 1,  0>
position=<-6,  0> velocity=< 2,  0>
position=< 5,  9> velocity=< 1, -2>
position=<14,  7> velocity=<-2,  0>
position=<-3,  6> velocity=< 2, -1>`

	input = readFile("input.txt")
	nodes := *parseNodes(input)

	//minimize bounding box as proxy for finding solution
	prev_size := bb_size(&nodes)
	curr_size := doStep(&nodes)
	i:= 0; for ; curr_size < prev_size; i++ {
		prev_size = curr_size
		curr_size = doStep(&nodes)
	}

	//go back one step
	doReverseStep(&nodes)
	fmt.Println("Part 1:")
	fmt.Print(plot(&nodes))
	fmt.Println()
	fmt.Println("Part 2:", i)
}