package main

import (
	"strings"
	"fmt"
	"io/ioutil"
)

var sample string = `.#.
..#
###`

var rules = `../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`


func main() {
/*	
	//sample
	image := NewPainting(sample, rules)
	image.print()

	image.evolve() ; image.print()
	image.evolve() ; image.print()
	fmt.Println(image.countOn())
*/

/*	//part1
	realRules := readInput("rules.txt")
	image := NewPainting(sample, realRules)
	for i := 0 ; i < 5 ; i++ { image.evolve() }
	fmt.Println( image.countOn() )
*/

	//part 2
	realRules := readInput("rules.txt")
	image := NewPainting(sample, realRules)
	for i := 0 ; i < 18 ; i++ { image.evolve() }
	fmt.Println( image.countOn() )


}

type Painting struct{
	chars [][]byte
	rules map[string][][]byte
}
func (p *Painting) size() int{
	//folding !!
	return len(p.chars)
}
func (p *Painting) print(){
	for _, line := range p.chars{
		fmt.Println( string(line) )
	}
	fmt.Println()
	fmt.Println()
}
func (p *Painting) countOn() int{
	ctr := 0
	for _, line := range p.chars{
		for _, c := range line {
			if c == '#' { ctr++ }
		}
	}
	return ctr
}
func (p *Painting) evolve(){
	osize := 3
	if p.size() % 2 == 0 { osize = 2 }
	
	nsize := (osize+1) * (p.size()/osize)

	//memalloc
	res := make([][]byte, nsize)
	for i, _ := range res { res[i] = make([]byte, nsize) }

	//loop over blocks
	for i, yb, ye := 0, 0, osize ; ye <= p.size() ; i, yb, ye = i+1, yb+osize, ye+osize {
		for j, xb, xe := 0, 0, osize ; xe <= p.size() ; j, xb, xe = j+1, xb+osize, xe+osize {

			//get lookup string
			from := ""
			for i := yb ; i < ye ; i++ { //masking of i, cuz YOLO
				for j := xb ; j < xe ; j++ { from += string(p.chars[i][j]) }
				from += "/"
			}
			from = from[:len(from)-1]

			//get mapping
			to,_ := p.rules[from]
			//fmt.Println(from, " => ", to)

			//put mapping in place
			for y, line := range to {
				for x, c := range line {
					res[i*(osize+1) +y][j*(osize+1) +x] = c
				}
			}

		}
	}

	p.chars = res
}
func NewPainting(start, rules string) *Painting{
	res := Painting {
		chars: *loadChars(start),
		rules: *loadRules(rules),
	}
	return &res
}
func loadChars(s string) *[][]byte{
	lines := strings.Split(s, "\n")
	res := make([][]byte, len(lines))
	for i, line := range lines { res[i] = []byte(line) }
	return &res
}
func loadRules(s string) *map[string][][]byte{
	
	res := map[string][][]byte{}
	for _, line := range strings.Split(s, "\n") {
		parts := strings.Split(line, " => ")
		
		from := parts[0]
		to := toBytes(parts[1])
		
		res[from], res[flip(from)] = to, to
		for i := 1 ; i < 4 ; i++ {
			from = rotate(from)
			res[from], res[flip(from)] = to, to
		}
	}
	return &res
}
func flip(s string) string{
	parts := strings.Split(s, "/")
	res := make([][]byte, len(parts))
	for i,p := range parts { res[i] = []byte(flipline(p)) }

	return toString(res)
}
func flipline(s string) string{
	res := []byte(s)
	for i, j := 0, len(s)-1 ; i <  j ; i,j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
func rotate(s string) string{
	
	parts := strings.Split(s, "/")

	//allocate memory
	res := make([][]byte, len(parts))
	for i, line := range parts{ res [i] = make([]byte, len(line)) }

	//put everythin in place
	for y, i := 0, len(parts)-1 ; y < len(parts) ; y,i = y+1, i-1 {
		for x := 0 ; x < len(parts[y]) ; x++ {
			res[x][i] = parts[y][x]
		}
	}

	return toString(res)
}
func toString(in [][]byte) string{
	res := ""
	for _, line := range(in){
		res += string(line)
		res += "/"
	}
	return res[:len(res)-1]
}
func toBytes(s string) [][]byte{
	parts := strings.Split(s, "/")
	res := make([][]byte, len(parts))
	for i,p := range parts { res[i] = []byte(p) }
	return res
}
func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)
	return string(s)
}