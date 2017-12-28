package main

import (
	"strconv"
	"fmt"
)


func main(){
	fmt.Println( Part1("hfdlxzhv") )
	fmt.Println( Part2("hfdlxzhv") )
}

func Part1(s string) int{
	res := 0
	for i := 0 ; i < 128 ; i++ {
		bytes := Hash(s + "-" + strconv.Itoa(i))
		res += CountDucks(bytes)
	}

	return res
}
func CountDucks(hash []byte) int{
	res := 0
	for _, bits := range hash {
		if bits & 1 == 1 { res++ }
		if bits & 2 == 2 { res++ }
		if bits & 4 == 4 { res++ }
		if bits & 8 == 8 { res++ }
		
		if bits & 16  == 16  { res++ }
		if bits & 32  == 32  { res++ }
		if bits & 64  == 64  { res++ }
		if bits & 128 == 128 { res++ }
	}

	return res
}

func Part2(s string) int{

	field := make([][]bool, 128)
	for i := 0 ; i < 128 ; i++ {
		input := s + "-" + strconv.Itoa(i)
		bytes := Hash(input)
		field[i] = makeLine(bytes)
	}
	return countZones(&field)
}
func makeLine(b []byte) []bool{
	
	res := make([]bool, 128)
	for i, bits := range b {
		for j, pos := 7, byte(1) ; j >= 0 ; j, pos = j-1, pos<<1 {
			res[i*8 + j] = (bits & pos == pos)	
		}
	}
	return res
}
func countZones(f *[][]bool) int{
	field := *f
	res := 0

	for y := 0 ; y < 128 ; y++ {
		for x := 0 ; x < 128 ; x++ {
			if field[y][x] { 
				clearPos(f, y, x)
				res++
			}
		}
	}
	return res
}
func clearPos(f *[][]bool, y,x int){
	field := *f
	if !field[y][x] { return }
	field[y][x] = false

	for _, d := range []struct{dy,dx int}{ {-1,0}, {0,-1},{0,1}, {1,0} }{
		yPos, xPos := y + d.dy, x+d.dx

		if yPos < 0 || yPos >= 128 || xPos < 0 || xPos >= 128 { continue }
		if field[yPos][xPos] { clearPos(f, yPos, xPos) }
	}
}
func printField(f *[][]bool){ //debug
	field := *f
	for _, line := range field {

		for _, val := range line {
			if val { 
				fmt.Print("1") 
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}


//copied from day 10
func Hash(input string) []byte{
	list := make([]byte, 256)
	for i := 0 ; i < 256; i++ { list[i] = byte(i) }

	inputLengths := handleInput(input)
	currPos, skipSize := 0,0
	for r := 0; r < 64 ; r++{
		_, currPos, skipSize = HashRound(currPos, skipSize, list,inputLengths)
	}

	return denseIt(list)
}
func HashRound(currPos, skipSize int, list, inputLengths []byte) (int,int,int){

	for _, l := range inputLengths {
		reverseSliceInPlace( list, currPos, int(l) )
		
		currPos += int(l) + skipSize
		skipSize++
	}

	return int(list[0]) * int(list[1]), currPos, skipSize
}
func reverseSliceInPlace(slice []byte, start, size int){
	//copy
	r := make([]byte, size)
	for i := 0; i < size ; i++ {
		r[i] = slice[ (start+i)%len(slice) ]
	}

	//reverse
	for b, e := 0, len(r)-1 ; b<e ; b,e = b+1, e -1 {
		r[b], r[e] = r[e], r[b]
	}

	//put back
	for i := 0; i < size; i++ {
		slice[ (start+i)%len(slice) ] = r[i]
	}
}
func handleInput(input string) []byte{
	
	res := make([]byte, len(input), len(input)+5)
	
	for i, c := range input{
		res[i] = byte(c)
	}

	res = append(res, []byte{17, 31, 73, 47, 23}...)
	return res
}
func denseIt(list []byte) []byte{
	
	res := make([]byte, 0, 16)

	for b,e := 0,16 ; e <= len(list) ; b, e = b+16, e+16 {

		xor := list[b]
		for i := b+1 ; i < e ; i++ {
			xor ^= list[i]
		}
		res = append(res, xor)
	}

	return res
}
func makeString(hash []byte) string{
	res := ""
	for _, b := range hash {
		c := fmt.Sprintf("%x", b)
		
		if len(c) < 2 { c = "0" + c}
		res += c
	}
	return res
}