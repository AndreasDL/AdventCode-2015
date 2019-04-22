package main

import (
	"fmt"
	"strconv"
	"regexp"
	"strings"
	"io/ioutil"
	"container/heap" //https://golang.org/pkg/container/heap/
)


//TBH this wasn't an easy challenge, it requires some efficient brute force!
//I also shamelessly stole some ideas from this woman: https://www.youtube.com/watch?v=C2b4iI0KQPY
//she solved it in 1h15 min :o

type Coord struct {
	x,y,z int
}
func (loc Coord) distance( p Coord) int{
	dist := abs(loc.x - p.x)
	dist += abs(loc.y - p.y)
	dist += abs(loc.z - p.z)
	return dist
}
func (loc Coord) moveInRange(bot Nanobot) Coord{

	//how much do we need to move ?
	distance_to_range := bot.distanceToRange(loc) + 1 

	//how much in each direction ?
	dx := bot.loc.x - loc.x
	dy := bot.loc.y - loc.y
	dz := bot.loc.z - loc.z
	total_dist := abs(dx) + abs(dy) + abs(dz)

	x := loc.x + (distance_to_range*dx)/total_dist
	y := loc.y + (distance_to_range*dy)/total_dist
	z := loc.z + (distance_to_range*dz)/total_dist

	return Coord{x,y,z}
}
func (loc Coord) distToZero() int{
	return loc.distance(Coord{0,0,0})
}


type Nanobot struct{
	loc Coord
	r int
}
func (bot Nanobot) inRangeOf(loc Coord)bool{
	return bot.loc.distance(loc) <= bot.r
}
func (bot Nanobot) distanceToRange(loc Coord) int{
	return loc.distance(bot.loc) - bot.r
}

type Candidate struct{
	loc Coord
	nearby int
}
func (c Candidate) distToZero() int{
	return c.loc.distToZero()
}

//priority queue using heap ensure that we loop in right order (best candidates first)
type PriorityQueue []Candidate
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int)bool {
	//less => better
	if (pq[i].nearby > pq[j].nearby) { //more bots nearby is better
		return true
	} else if (pq[i].nearby < pq[j].nearby){
		return false
	} else {
		//same number of bots, then we should look to the point closest to zero
		return pq[i].distToZero() < pq[j].distToZero()
	}
}
func (pq PriorityQueue) Swap(i,j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(s interface{}){ *pq = append(*pq, s.(Candidate)) }
func (pq *PriorityQueue) Pop() interface{}{ 
	last := len(*pq) -1
	res := (*pq)[last]
	*pq = (*pq)[:last]
	return res
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
func inRange(bots []Nanobot, loc Coord) int {
	cnt := 0
	for _, b := range bots {
		if b.inRangeOf(loc) { cnt++ }
	}
	return cnt
}

//input parsing
var re = regexp.MustCompile("-?\\d+")
func parseLine(line string) Nanobot{
	parts := re.FindAllString(line, 4)
	pi := make([]int, len(parts))
	for i, p := range parts{ pi[i] = toInt(p) }

	return Nanobot{
		Coord{pi[0], pi[1], pi[2]},
		pi[3],
	}
}
func parseInput(line string) []Nanobot{
	lines := strings.Split(line, "\n")
	result := make([]Nanobot, len(lines))
	for i, line := range lines {
		result[i] = parseLine(line)
	}
	return result
}

func main() {
	
	input := `pos=<10,12,12>, r=2
pos=<12,14,12>, r=2
pos=<16,12,12>, r=4
pos=<14,14,14>, r=6
pos=<50,50,50>, r=200
pos=<10,10,10>, r=5`
	
	input = readFile("input.txt")
	bots := parseInput(input)
	
	//find starting points
	//midpoints with most bots nearby!
	fmt.Println("find best midpoints to start with")
	max := 0
	candidates := make(PriorityQueue,0)
	for x, i := range bots {
		for y, j := range bots {

			if y >= x { break }

			midPoint := Coord{
				(i.loc.x + j.loc.x)/2,
				(i.loc.y + j.loc.y)/2,
				(i.loc.z + j.loc.z)/2,
			}
			nearby := inRange(bots, midPoint)

			if nearby > max { //we found a candidate with more bots nearby! => throw out old candidates
				max = nearby
				candidates = candidates[:0]
				heap.Init(&candidates)
			}

			if nearby >= max {
				candidates = append(candidates, Candidate{midPoint, nearby})
			}
		}
	}
	heap.Init(&candidates)

	for _, c := range candidates { fmt.Println(c) }
	fmt.Println(len(candidates))


	//for each candidate, move to get an additional point in range
	fmt.Println("Searching")
	seen := map[Candidate]bool{}
	best := candidates[0]
	for i := 0; i < 10000 && len(candidates) > 0 ; i++ {

		//take current best candidate
		curr := heap.Pop(&candidates).(Candidate)
		fmt.Println(curr)

		//keep track of seen positions
		if _, ex := seen[curr] ; ex { continue }
		seen[curr] = true

		//try to make it a bit better by walking towards another bot
		for _, bot := range bots {

			if bot.inRangeOf(curr.loc) { continue }

			//move in range of this bot
			next_loc := curr.loc.moveInRange(bot)
			next_nearby := inRange(bots, next_loc)
			next := Candidate{
				next_loc, 
				next_nearby,
			}

			//keep track of best seen point
			if next.nearby > best.nearby {
				best = next
				fmt.Println(best)
			}

			//only store better of same quality candidates
			if next_nearby >= curr.nearby { 
				heap.Push(&candidates, next)
			}
		}
	}

	fmt.Println(best)
	fmt.Println(best.distToZero())

	fmt.Println("can we move closer?")
	best_loc := best.loc
	fmt.Println(best_loc, inRange(bots, best_loc))
	best_loc.x--
	fmt.Println(best_loc, inRange(bots, best_loc))
	best_loc.x++
	best_loc.y--
	fmt.Println(best_loc, inRange(bots, best_loc))
	best_loc.y++
	best_loc.z--
	fmt.Println(best_loc, inRange(bots, best_loc))

}

//938 - 31309239 14759931 25415470 dist: 71484640 too low!
//938 - 31309240 14759931 25415471 dist: 71484642 is correct!
//submitted solution not correct: should be (same distance, more in range)
//976 - 31309240 15032338 25143064 dist 7184642

//i give up ..
//the algorithm should explore move once it is done getting most drones in range 
//(now it is in a local optiman that happend to be the same distance away from zero as the solution)

//TODO
//check this blogpost for better approach 
// https://ragona.com/posts/advent_of_code_day_23

//or use this code to check results
//https://github.com/lizthegrey/adventofcode/blob/master/2018/day23.go