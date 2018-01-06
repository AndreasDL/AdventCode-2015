package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	//guests := NewGuests(readInput("sampleInput.txt"))

	guests := NewGuests( readInput("input.txt") )
	fmt.Println("Part1: ", guests.OptimalConfig() )


	//part 2
	guests.scores["me"] = map[string]int{}
	for key, _ := range guests.scores {
		guests.scores["me"][key] = 0
		guests.scores[key]["me"] = 0
	}

	fmt.Println("Part2: ", guests.OptimalConfig() )	
}

type Guests struct{
	scores map[string]map[string]int
}
func NewGuests(input string) Guests {

	res := map[string]map[string]int{}

	for _, line := range strings.Split(input, "\n"){

		fields := strings.Fields(line)

		from   := fields[0]
		val, _ := strconv.Atoi(fields[3])
		if fields[2] == "lose" { val = -val }
		to     := fields[10]
		to      = to[:len(to)-1]

		if _, ex := res[from] ; ! ex { res[from] = map[string]int{} }
		res[from][to] = val
	}

	return Guests{
		scores: res,
	}
}
func (g Guests) Names() []string {
	names := make([]string, 0, len(g.scores))
	for name := range g.scores {
		names = append(names, name)
	}
	return names
}
func (g Guests) Permutations()[][]string{
	//https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
    var helper func([]string, int)
    res := [][]string{}

    helper = func(arr []string, n int){
        if n == 1{
            tmp := make([]string, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    names := g.Names()
    helper(names, len(names))
    return res
}
func (g Guests) SeatingCosts(seating []string) int{
	res := 0
	for i := 0 ; i < len(seating) ; i++ {

		next := (i + 1) % len(seating)

		from := seating[i]
		to   := seating[next]

		res += g.scores[from][to]
		res += g.scores[to][from]
	}

	return res
}
func (g Guests) OptimalConfig() int{
	//calculate all costs & keep lowest!
	maxScore := 0
	for _, p := range g.Permutations(){

		score := g.SeatingCosts(p)
		//fmt.Println(p, " => ", score)

		if score > maxScore { maxScore = score }
	}

	return maxScore
}





func readInput(fname string) string{
	res, _ := ioutil.ReadFile(fname)
	return string(res)
}