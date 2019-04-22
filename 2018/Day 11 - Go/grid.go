package main


import (
	"fmt"
)


func powerlevel(serial, y,x int) int {
	rank := x+10

	power := rank * y
	power += serial
	power *= rank

	power /= 100
	power %= 10

	power -= 5

	return power
}
func initField(serial int) *[][]int{
	size := 300

	grid := make([][]int, size)
	for i, _ := range grid {
		grid[i] = make([]int, size)
	}

	for y, line := range grid {
		for x, _ := range line {
			grid[y][x] = powerlevel(serial, y+1, x+1)
		}
	}
	return &grid
}

func cellPower(grid *[][]int, y,x int, size int) int{
	_grid := *grid

	sum := 0
	for j := y; j < y+size; j++ {
		for i := x; i < x+size; i++ {
			sum += _grid[j][i]
		}
	}
	return sum
}
func findSquare(grid *[][]int, serial, size int) (int,int, int){
	_grid := *grid

	maxy, maxx, max_power := -1,-1,-1
	for y := 0; y < len(_grid)-size; y++ {
		for x := 0; x < len(_grid[y])-size; x++{
			if power := cellPower(grid,y,x,size); power > max_power{
				max_power = power
				maxy = y
				maxx = x
			}
		}
	}

	//indexing starts at 1 </3
	return maxy+1, maxx+1, max_power
}

func part1(serial int) (int, int){
	grid := initField(serial)
	y, x, _ := findSquare(grid, serial,3)
	
	return x, y
}
func part2(serial int)(int, int, int){
	
	grid := initField(serial)
	max_power, maxx, maxy, max_size := -1, 0,0,0
	for size := 1; size < 300; size++{

		//fmt.Println(size)
		y, x, power := findSquare(grid, serial, size)
		if power > max_power {
			max_power = power
			maxx = x
			maxy = y
			max_size = size
		}	
	}

	return maxx,maxy, max_size
}

func main(){
	serial := 6303

	x, y := part1(serial)
	fmt.Println("Part1:", x,y)

	x,y,size := part2(serial)
	fmt.Println("Part2", x,y,size)
}