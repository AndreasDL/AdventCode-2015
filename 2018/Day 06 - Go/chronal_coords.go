package main

import (
	"fmt"
	"strings"
	"strconv"
	"io/ioutil"
)

type Point struct{
	id int
	y, x int
}

func readFile(fname string) string{
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}
func parseInput(s string) (*[]Point, int, int){

	max_x, max_y := -1, -1
	lines := strings.Split(s, "\n")
	points := make([]Point, len(lines))

	for i, line := range lines {

		parts := strings.Split(line, ",")

		x, _:= strconv.Atoi(parts[0])
		y, _:= strconv.Atoi(parts[1][1:]) //get rid of , 
		
		points[i] = Point{
			i+1,
			y,
			x,
		}

		//keep track of field size
		if y > max_y { max_y = y }
		if x > max_x { max_x = x }
	}

	return &points, max_y, max_x
}

func getDist(_p1, _p2 *Point) int{
	p1, p2 := *_p1, *_p2
	dx := p1.x - p2.x
	dy := p1.y - p2.y

	if dx < 0 { dx = -dx }
	if dy < 0 { dy = -dy }
	return dx + dy
}
func getNearest(points *[]Point, center *Point) (int, bool){
	
	//init
	min_id, min_dist := -1, -1
	valid := true

	for _, p := range *points {
		dist := getDist(center, &p)
		
		if dist == 0 { continue } //skip comparing to own point

		if dist < min_dist || min_dist == -1 {
			min_dist = dist
			min_id = p.id
			valid = true
		} else if dist == min_dist {
			valid = false
		}
	}

	return min_id, valid
}
func getTotalDist(points *[]Point, center *Point, cutoff int) (int, bool) {
	dist := 0
	for _, p := range *points {
		dist += getDist(center, &p)

		if dist >= cutoff {
			return dist, false
		}
	}
	return dist, true
}

func createField(max_y, max_x int) *[][]int{
	field := make([][]int, max_y+1)
	for y, _ := range field {
		field[y] = make([]int, max_x+1)
	}
	return &field
}
func initField(_field *[][]int, points *[]Point){
	//set points
	field := *_field
	for _, p := range *points {
		field[p.y][p.x] = p.id
	}
}
func expandField(_field *[][]int, points *[]Point){
	field := *_field

	for y, line := range field {
		for x, val := range line {

			if val != 0 { continue }

			if id, ok := getNearest(points, &Point{0, y, x}); ok {
				field[y][x] = id
			}
		}
	}
}
func removeInfFields(_field *[][]int){

	field := *_field
	ids := map[int]bool{}
	height, width := len(field), len(field[0])

	//upper & bottom row
	for x := 0 ; x < width ; x++ {
		ids[field[0][x]] = true
		ids[field[height-1][x]] = true
	}	
	//left and right column
	for y := 0 ; y < height ; y++ {
		ids[field[y][0]] = true
		ids[field[y][width-1]] = true
	}

	//remove from field ==> set to zero
	for y, line := range field {
		for x, val := range line {

			if _, exists := ids[val]; exists {
				field[y][x] = 0
			}
		}
	}
}
func count(field *[][]int) int{
	max_cnt := -1
	cnt := map[int]int{} //map[id] = cnt
	for _, line := range *field {
		for _, val := range line {
			cnt[val]++

			if cnt[val] > max_cnt && val != 0{
				max_cnt = cnt[val]
			}
		}
	}

	return max_cnt
}
func distanceField(_field *[][]int, points *[]Point, cutoff int){
	field := *_field

	for y, line := range field {
		for x, _ := range line {
			if _, ok := getTotalDist(points, &Point{0,y,x}, cutoff); ok {
				field[y][x] = 1
			}
		}
	}
}
func printField(field *[][]int){
	for _, line := range *field { 
		fmt.Println(line) 
	}
}

func main(){
		
	//Part 1
	input := readFile("input.txt")
	points, max_y, max_x := parseInput(input)
	field := createField(max_y, max_x) //create empty field
	initField(field, points) //set points
	expandField(field, points) //calculate closest points
	removeInfFields(field) //remove infinite fields
	res := count(field) //get sizes
	fmt.Println("Part 1:", res)

	//Part2
	cutoff := 10000
	//input = readFile("input.txt")
	//points, max_y, max_x = parseInput(input)
	field = createField(max_y, max_x) //create empty field
	distanceField(field, points, cutoff) //mark where the distances are below cutoff
	res = count(field) //get size
	fmt.Println("Part 2:", res)
}