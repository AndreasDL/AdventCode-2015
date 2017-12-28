package main

import "fmt"

func main() {
	a, b := 65, 8921
	fmt.Println( Part2(a,b,5000000) )

	a,b  = 512, 191
	fmt.Println( Part2(a,b,5000000) )
}

func Part2(a,b int, iterations int) int{
	outA, outB, outS := make(chan int, 1), make(chan int, 1), make(chan int, 1)
	
	go genNextMult(a, 16807, 4, outA)
	go genNextMult(b, 48271, 8, outB)
	go judge(outA, outB, iterations, outS)

	return <-outS
}
func genNextMult(val, fact, multiple int, output chan int){
	for {
		if val % multiple == 0 { output <- val%65536 }
		val = (val*fact) % 2147483647
	}
}
func judge(outA,outB chan int, iterations int, output chan int){
	ctr := 0
	for i := 0 ; i < iterations ; i++ {
		if <-outA == <-outB { ctr++ }
	}
	output <- ctr
}