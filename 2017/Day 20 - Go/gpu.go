package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type particle struct{
	px, py, pz int
	vx, vy, vz int
	ax, ay, az int
}
func (p *particle) step(){
	p.vx += p.ax
	p.vy += p.ay
	p.vz += p.az

	p.px += p.vx
	p.py += p.vy
	p.pz += p.vz
}
func (p *particle) distanceOrigin() int {
	return abs(p.px)+abs(p.py)+abs(p.pz)
}
func (p *particle) accelerationOrigin() int{
	return abs(p.ax)+abs(p.ay)+abs(p.az)
}
func abs(i int) int{
	if i < 0 { return -i }
	return i
}
func (p *particle) SameLocation(q *particle) bool{
	return p.px == q.px && p.py == q.py && p.pz == q.pz
}

func main() {
	input := readInput("input.txt")
	particles := *(parseInput(input))

	//part 1
	min := 0
	for i, p := range particles {
		if (*particles[min]).accelerationOrigin() > (*p).accelerationOrigin() { min = i }
	}
	fmt.Println(min)

	//part2
	//remove colissions at start 
	for kp, p := range particles {
		for kq, q := range particles {
			if kp == kq { //cannot collide with own particle
				continue 
			} else if p.SameLocation(q){ 
				delete(particles, kq)
				delete(particles, kp)
			}
		}
	}
	for step := 0 ; step < 1000 ; step++ {
		//do steps
		for _, p := range particles { p.step() }

		//remove colissions
		for kp, p := range particles {
			for kq, q := range particles {
				if kp == kq { //cannot collide with own particle
					continue 
				} else if p.SameLocation(q){ 
					delete(particles, kq)
					delete(particles, kp)
				}
			}
		}
	}
	fmt.Println( len(particles) )
}

func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}
func parseInput(s string) *map[int]*particle{
	
	res := map[int]*particle{}
	for i, line := range strings.Split(s, "\n") {

		line = line[:len(line)-1]	
		parts := strings.Split(line, ">, ")
		p, v, a := parsePart(parts[0]), parsePart(parts[1]), parsePart(parts[2])

		res[i] = &particle{
			p[0], p[1], p[2],
			v[0], v[1], v[2],
			a[0], a[1], a[2],
		}
	}

	return &res
}
func parsePart(s string) [3]int{
	s = s[3:]

	res := [3]int{}
	for i, part := range strings.Split(s, ","){
		res[i], _ = strconv.Atoi(part)
	}
	return res
}