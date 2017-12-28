package main

import (
	"fmt"
)

/*
		65 64 63 62 61 60 59 58 57 
		66 37 36 35 34 33 32 31 56
		67 38 17 16 15 14 13 30 55
		68 39 18  5  4  3 12 29 54
		69 40 19  6  1  2 11 28 53
		70 41 20  7  8  9 10 27 52
		71 42 21 22 23 24 25 26 51
		72 43 44 45 46 47 48 49 50
		73 74 75 76 77 78 79 80 81
*/


func main(){

	spiral(1)
	spiral(12)
	spiral(23)
	spiral(1024)
	spiral(325489)

}

func spiral(v int)int{
	fmt.Println(v)
	ring, min, max := onRing(v)
	fmt.Println("\tring:", ring, "which contains values between", min, "and", max)

	//distance => moves to quandrant center + ring
	distOffset := quadrantCenterOffset(v, min, max)
	fmt.Println("\tdist =>", distOffset, "+" ,ring, " => ", distOffset+ring)

	return distOffset + ring
}

//onRing(100) => ring = 5, ringMin = 82, ringMax = 121
func onRing(v int) (int, int, int){
	ring := 0
	rMin, rMax := 1, 1
	
	for rMax < v{
		ring++
		rMin = rMax + 1
		rMax = ringMax(ring)
	}

	return ring, rMin, rMax
}

func ringMax(r int) int{
	rMax := 1+2*r
	return rMax * rMax
}

func quadrantCenterOffset(v, rMin, rMax int) int{

	//size of each edge
	edgeSize := (rMax - rMin) / 4 +1

	//which side are we on ?
	side := (v - rMin)/edgeSize

	//center of side if halway the side !
	centerOffset := edgeSize/2
	sideCenter := (side+side+1)* centerOffset + rMin-1

	res := sideCenter - v
	if res < 0 { res = -res }

	return res
}