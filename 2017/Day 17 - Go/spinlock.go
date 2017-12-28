package main

import "fmt"


func main() {
	
	//fmt.Println(part1(2017, 3  ))
	//fmt.Println(part1(2017, 370))

	for i := 0 ; i < 10 ; i++ { fmt.Println( part2(i, 3) ) }
	
	fmt.Println( part2(50000000, 370) )
}

func part1(iterations, stepsForward int) *[]int{
	currPos := 0
	buffer := &[]int{0}
	for i := 1; i <= iterations ; i++ {

		currPos = (currPos+stepsForward) % len( *buffer )
		currPos++
		
		buffer = insertSlice(buffer, currPos, i)

		//fmt.Println(i, currPos, buffer)
		if i % 10000 == 0 { fmt.Println(i) }
	}
	return buffer
}
func insertSlice(input *[]int, pos,val int) *[]int{
	orig := *input
	res := make([]int, len(orig)+1)

	for i := 0 ; i < pos ; i++ { res[i] = orig[i] }
	res[pos] = val
	for i := pos ; i < len(orig) ; i++ { res[i+1] = orig[i] }
	
	return &res
}

func part2(iterations, stepsForward int) int{
	value, currPos := 0,0
	for i := 1; i <= iterations ; i++ {

		currPos = (currPos+stepsForward+1) % i
		if currPos == 0 { value = i	}

		if i % 1000000 == 0 { fmt.Println(i) }
	}
	return value
}