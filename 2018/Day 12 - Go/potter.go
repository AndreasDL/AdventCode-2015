package main


import (
	"fmt"
	"strings"
	"io/ioutil"
)

func readFile(fname string) *[]string{
	file, _ := ioutil.ReadFile(fname)
	res := strings.Split(string(file), "\n")
	return &res
}
func loadMappings(mappings *[]string) *map[string]string {
	_mappings := *mappings
	
	res := make(map[string]string, len(_mappings))
	for _, line := range _mappings{

		parts := strings.Split(line, " => ")
		res[parts[0]] = parts[1]
	}

	return &res
}

//remove leading and trailing empty pots
func clean(state string, first int) (string, int){
	
	start := 0; for ; state[start] == '.' ; start++ {}
	stop := len(state)-1; for ; state[stop] == '.'; stop-- {}

	return state[start:stop+1], first + start
}
func next(mappings *map[string]string, state string, first int) (string, int){

	state = "....." + state + "....."
	first -= 3 //first two empty pots are only used as input

	next_state := ""
	x,y:=0,5; for ; y <= len(state) ; x,y = x+1,y+1{

		if next, exists := (*mappings)[state[x:y]]; exists{
			next_state += next
		} else {
			next_state += "."
		}
	}

	next_state, first = clean(next_state, first)
	return next_state, first
}

func count(state string, pos int) int{
	//get result
	res := 0
	for _, c := range state {
		if c == '#' { res += pos }
		pos++
	}

	return res
}

func part1(mappings *map[string]string, state string, it int) int{
	pos := 0
	for i := 0; i < it ; i++ {
		state, pos = next(mappings, state, pos)
	}

	return count(state, pos)
}


func run(mappings *map[string]string, state string, it int) int{
	pos := 0
	prev_cnt := count(state, pos)
	prev_delta := -1
	for i := 0; i < it ; i++ {

		state, pos = next(mappings, state, pos)
		cnt := count(state, pos)
		
		if i % 100 == 0 {
			delta := cnt - prev_cnt	
			if delta == prev_delta {
				fmt.Println("same delta at ", i, " start val:", cnt, "delta", delta)
				return cnt + delta * (it-i-1)
			}
			prev_delta = delta
		}
		
		prev_cnt = cnt
		
	}
	return 0
}


func main(){
	

	init_state := "####..##.##..##..#..###..#....#.######..###########.#...#.##..####.###.#.###.###..#.####..#.#..##..#"
	input := readFile("input.txt")
	mappings := loadMappings(input)
	fmt.Println("Part 1:", part1(mappings, init_state, 20))
	fmt.Println("Part 2:", run(mappings, init_state, 50000000000))
}