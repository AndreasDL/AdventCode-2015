package main

import (
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
	"sort"
)

//let's not use a tree like structure !
//maps ftw !

func main(){
	input := readInput("input.txt")

	fmt.Println(part1(input))

	fmt.Println(part2(input))


}


func part2(input string) int{
	children, weights := getChildrenAndWeights(input)
	parents := getParents(children)
	root := getRoot(parents, children)

	//fmt.Println(weights)
	balance(root, children, weights)
	//fmt.Println(weights)
	
	return 0
}


func part1(input string) string{
	children, _ := getChildrenAndWeights(input)
	//fmt.Println( *children )
	parents := getParents(children)
	//fmt.Println( *parents )

	root := getRoot(parents, children)
	return root
}

//map[parent] = list of children
func getChildrenAndWeights(s string) (*map[string][]string, *map[string]int) {
	children := map[string][]string{}
	weights := map[string]int{}

	s = strings.Replace(s, "," , "", -1)
	for _, line := range strings.Split(s, "\n"){

		parts := strings.Split(line, " ")
		weight, _ := strconv.Atoi(parts[1][1:len(parts[1])-1])
		weights[parts[0]] = weight

		if len(parts) > 3 {
			children[parts[0]] = append(children[parts[0]], parts[3:]...)
		} else {
			children[parts[0]] = []string{}
		}

	}

	return &children, &weights
}
//Reverse the map => map[child] = parent
func getParents(children *map[string][]string) *map[string]string {
	parents := map[string]string{}
	for parent, childs := range *children{
		for _, c := range childs{
			parents[c] = parent
		}
	}
	return &parents
}
//parent should have at least one child
//loop over nodes with children & check if their parent exists
func getRoot(parents *map[string]string, children *map[string][]string)string {

	//assumption parent has a child
	par, chi := *parents, *children
	for c := range chi{
		if _, ex := par[c] ; !ex { return c}
	}

	return ""
}

func calculateWeight(parent string, children *map[string][]string, weights *map[string]int) int{

	if len( (*children)[parent]) == 0 {
		return (*weights)[parent]
	}

	weight := (*weights)[parent]
	for _, c := range (*children)[parent] {		
		weight += calculateWeight(
			c,
			children,
			weights,
		)
	}

	return weight
}
func balance(parent string, children *map[string][]string, weights *map[string]int){
	chi := (*children)[parent]

	if len(chi) <= 1 { return }

	//balance children first !
	for _, c := range chi {
		balance(
			c,
			children,
			weights,
		)
	}

	//check if balanced
	sort.Slice( chi, 
		func(i,j int)bool { 
			return calculateWeight(chi[i], children, weights) < 
				calculateWeight(chi[j], children, weights) 
			},
	)

	nameUnBalanced, wUnBal := chi[0]           , calculateWeight(chi[0], children, weights)
	nameBalanced  , wBal   := chi[ len(chi)-1 ], calculateWeight(chi[len(chi)-1], children, weights)

	//not balanced 
	if len(chi) >2 && wBal != wUnBal { 

		//Adjust balanced / unbalanced if needed unbalanced should be the one that differs
		if len(chi) > 2 && wUnBal == calculateWeight(chi[1], children, weights){
			nameBalanced, nameUnBalanced = nameUnBalanced, nameBalanced
			wBal, wUnBal = wUnBal, wBal
		}
		
		//adjust
		delta := wBal - wUnBal
		(*weights)[nameUnBalanced] += delta
		
		//ugly way of returning & exitting
		fmt.Println( (*weights)[nameUnBalanced] )
		panic(nil)
	}
}
	

func readInput(fname string) string {
	s, _ := ioutil.ReadFile(fname)

	return string(s)
}