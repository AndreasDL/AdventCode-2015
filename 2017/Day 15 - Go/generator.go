package main

import "fmt"

func main() {
	a, b := 65, 8921
	fmt.Println( Part1(a,b,40000000) )
	fmt.Println( Part2(a,b,5000000) )

	a,b = 512, 191
	fmt.Println( Part1(a,b,40000000) )
	fmt.Println( Part2(a,b,5000000) )
}

func Part1(a,b int, iterations int) int{
	factA, factB := 16807, 48271

	ctr:= 0
	for i := 0 ; i < iterations ; i++ {
		a = genNext(a, factA)
		b = genNext(b, factB)
	
		if a%65536 == b%65536 { ctr++ }
	}
	return ctr
}
func genNext(val, fact int) int{
	return (val*fact) % 2147483647
}


func Part2(a,b int, iterations int) int{
	factA, factB := 16807, 48271
	multA, multB := 4, 8

	ctr := 0
	for i := 0 ; i < iterations ; i++ {
		a = genNextMult(a, factA, multA)
		b = genNextMult(b, factB, multB)

		if a % 65536 == b %65536 { ctr++ }
	}

	return ctr
}
func genNextMult(val, fact, multiple int) int{
	next := genNext(val, fact)
	for next % multiple != 0 {
		next = genNext(next, fact)
	}

	return next
}