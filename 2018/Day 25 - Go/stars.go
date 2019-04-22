package main


import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"io/ioutil"
)

var re = regexp.MustCompile("-?\\d+")
func toInt(s string) int{
	v, _ := strconv.Atoi(s)
	return v
}
func abs(x int) int{
	if x < 0 { return -x}
	return x
}
func min(a,b int) int{
	if a < b { return a }
	return b
}
func max(a,b int) int{
	if a > b { return a }
	return b
}
func readFile(fname string) string{
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}

type Lookup map[int]map[int]map[int]map[int]int //Lookup[x][y][z][t] = constellation

type Point [4]int
func parsePoint(line string) Point{
	parts := re.FindAllString(line, 4)
	pi := Point{}
	for i, p := range parts{ 
		pi[i] = toInt(p) 
	}
	return pi
}
func parseInput(input string) (Lookup, []Point){
	lines := strings.Split(input, "\n")
	result := Lookup{} 
	points := make([]Point, len(lines))

	for i, line := range lines{
		p := parsePoint(line)
		points[i] = p

		if _, ex := result[p[0]]; !ex { result[p[0]] = map[int]map[int]map[int]int{} }
		if _, ex := result[p[0]][p[1]]; !ex { result[p[0]][p[1]] = map[int]map[int]int{} }
		if _, ex := result[p[0]][p[1]][p[2]]; !ex { result[p[0]][p[1]][p[2]] = map[int]int{} }

		result[p[0]][p[1]][p[2]][p[3]] = i
	}

	return result, points
}
func distance(p1, p2 Point) int{
	dist := 0
	for i, x := range p1 {
		y := p2[i]
		dist += abs(x-y)
	}
	return dist
}
func sameConstellation(p1,p2 Point)bool{
	return distance(p1,p2) <= 3
}

func fillLookup(lookup Lookup, points []Point) Lookup{

	//fast lookup via maps
	//only lookup in dimensions that are no more than 3 away (manhattan distance)
	//skip non existing dimensions
	//could add additional optimization => if x dimension has a difference of 2 than other dimensions can only differ for 1
	for _, p := range points {
		curr := lookup[p[0]][p[1]][p[2]][p[3]]

		for x := p[0]-3 ; x <= p[0]+3 ; x++ {
			if _, ex := lookup[x]; !ex  { continue }

			for y := p[1]-3 ; y <= p[1]+3 ; y++ {
				if _, ex := lookup[x][y]; !ex  { continue }

				for z := p[2]-3 ; z <= p[2]+3 ; z++ {
					if _, ex := lookup[x][y][z]; !ex  { continue }

					for t := p[3]-3 ; t <= p[3]+3 ; t++ {

						if c, ex := lookup[x][y][z][t]; ex  { 
							if sameConstellation(p, Point{x,y,z,t}) {
								merge(lookup, curr, c) //same constellations
							}
						}
					}
				}
			}
		}
	}

	return lookup
}
func merge(lookup Lookup, curr, c int){
	curr, c = min(curr, c), max(curr, c) //always merge to lowest number not really necessary but ocd
	for x, _:= range lookup{
		for y, _ := range lookup[x]{
			for z, _ := range lookup[x][y]{
				for t, val := range lookup[x][y][z]{

					if val == c { 
						lookup[x][y][z][t] = curr
					}

				}
			}
		}
	}
}
func extractConstallations(lookup Lookup) map[int][]Point{
	//loop over lookup and get the distinct constallations
	solution := map[int][]Point{}
	for x, _:= range lookup{
		for y, _ := range lookup[x]{
			for z, _ := range lookup[x][y]{
				for t, val := range lookup[x][y][z]{
					solution[val] = append(solution[val], Point{x,y,z,t})
				}
			}
		}
	}

	return solution
}

func part1(input string) int{
	lookup, points := parseInput(input)
	lookup = fillLookup(lookup, points)
	solution := extractConstallations(lookup)
	return len(solution)
}

func main(){
	 inputs := []string{
`0,0,0,0
 3,0,0,0
 0,3,0,0
 0,0,3,0
 0,0,0,3
 0,0,0,6
 9,0,0,0
12,0,0,0`,

`-1,2,2,0
0,0,2,-2
0,0,0,-2
-1,2,0,0
-2,-2,-2,2
3,0,2,-1
-1,3,2,2
-1,0,-1,0
0,2,1,-2
3,0,0,0`,

`1,-1,0,1
2,0,-1,0
3,2,-1,0
0,0,3,1
0,0,-1,-1
2,3,-2,0
-2,2,0,0
2,-2,0,-1
1,-1,0,-1
3,2,0,2`,

`1,-1,-1,-2
-2,-2,0,1
0,2,1,3
-2,3,-2,1
0,2,3,-2
-1,-1,1,-2
0,-2,-1,0
-2,2,3,-1
1,2,2,0
-1,-2,0,-2`,
}
	//should give 2,4,3,8
	for _, input := range inputs {
		fmt.Println("Part 1:", part1(input))
	}

	input := readFile("input.txt")
	fmt.Println("Part 1:", part1(input))
	


}