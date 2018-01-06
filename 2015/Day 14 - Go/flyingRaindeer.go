package main

import (
	"strings"
	"io/ioutil"
	"strconv"
	"fmt"
)

func main() {

	//input := readInput("sampleInput.txt") ; time := 1000
	input := readInput("input.txt")         ; time := 2503
	deers := parseInput(input)

	max := 0
	for _, d := range deers { 
		dist := d.distanceAt(time)
		if dist > max { max = dist }
	}
	fmt.Println("Part1: ", max)


	points := map[int]int{}
	for sec := 1 ; sec <= time ; sec++ {
			
		//find max dist or tie
		maxDist := 0
		maxDeer := []int{}
		for i, d := range deers { 
			if dist := d.distanceAt(sec) ; dist > maxDist { 
				maxDist = dist 
				maxDeer = []int{i}
			} else if dist == maxDist {
				maxDeer = append(maxDeer, i)
			}
		}

		//award points
		for _, d := range maxDeer { points[d]++	}
		
	}

	maxPoints := 0
	for _, p := range points { 
		if p > maxPoints { maxPoints = p }
	}
	fmt.Println("Part2: ", maxPoints)


}

type reindeer struct {
	speed    int
	duration int
	rest     int
}

/*
Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
*/
func (r reindeer) distanceAt(sec int) int {

	burstDist     := r.speed * r.duration
	cycleTime     := r.duration + r.rest

	noBursts      := int(sec / cycleTime)
	timeRemaining := sec - (noBursts * cycleTime)

	if timeRemaining > r.duration { timeRemaining = r.duration }


	return int(noBursts) * burstDist + int(timeRemaining) * r.speed
	//distance from bursts + remaining burst
}

func parseInput(input string) []reindeer{
	lines := strings.Split(input, "\n")
	res := make([]reindeer, len(lines))

	for i, line := range strings.Split(input, "\n"){

		fields := strings.Fields(line)

		sp, _ := strconv.Atoi(fields[3])
		du, _ := strconv.Atoi(fields[6])
		re, _ := strconv.Atoi(fields[13])

		res[i] = reindeer{
			speed    : sp,
			duration : du,
			rest     : re,
		}
	}

	return res
}

func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}