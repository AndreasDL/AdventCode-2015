package main

import "fmt"
import "math"

/*
	147	142	133	122  59
	304   5   4   2  57
	330  10   1   1  54 
	351  11  23  25  26 1914
	362 747 806 880 931 957
*/


func main(){
	//loopTill(1)
	loopTill(7)
	loopTill(24)
	loopTill(25)
	loopTill(36)

	loopTill(325489)
}

func loopTill(v int){

	fmt.Println(v)

	//allocate spiral
	dim := int(math.Sqrt( float64(v) )) + 3 //additional shells => no out of indexes :D !
	spiral := make([][]int, dim)
	for i := range spiral {	spiral[i] = make([]int, dim) }

	x := (dim-1)/2 //longest spiral in allocated memory starts here 
	y := dim/2
	spiral[y][x] = 1
	
	res := startLookingFor(v, y,x, &spiral)

	//fmt.Println(spiral)
	fmt.Println(res)
}

var directions = []struct{
	dx, dy int
}{
	{-1,-1},
	{-1, 0},
	{-1, 1},
	{ 0,-1},
	{ 0, 1},
	{ 1,-1},
	{ 1, 0},
	{ 1, 1},
}

func sumSurroundings(y,x int, ptr *[][]int) int{

	spiral := *ptr
	
	if spiral[y][x] != 0 { return spiral[y][x] }

	val := 0
	for _, dir := range directions{
		val += spiral[y+dir.dy][x+dir.dx]
	}

	return val
}

func startLookingFor(val int, y,x int, ptr *[][]int) int{
	spiral := *ptr

	//round 1 goes like => 1 step right, then 1 up, 2 left, 2 down , 2 right
	//each round each direction increases by 2
	up, left, down, right := 1,2,2,2
	for spiral[y][x] < val{
		x++
		spiral[y][x] = sumSurroundings(y,x, &spiral)
		if spiral[y][x] > val { return spiral[y][x] }

		//up
		for i := 0 ; i < up ; i++ {
			y--
			spiral[y][x] = sumSurroundings(y,x, &spiral)
			if spiral[y][x] > val { return spiral[y][x] }
		}

		//left
		for i := 0 ; i < left ; i++ {
			x--
			spiral[y][x] = sumSurroundings(y,x, &spiral)
			if spiral[y][x] > val { return spiral[y][x] }
		}

		//down
		for i := 0 ; i < down ; i++ {
			y++
			spiral[y][x] = sumSurroundings(y,x, &spiral)
			if spiral[y][x] > val { return spiral[y][x] }
		}

		//right
		for i := 0 ; i < right ; i++ {
			x++
			spiral[y][x] = sumSurroundings(y,x, &spiral)
			if spiral[y][x] > val { return spiral[y][x] }
		}

		//next round!
		up, left, down, right = up+2, left+2, down+2, right+2
	}
	
	return spiral[y][x]
}