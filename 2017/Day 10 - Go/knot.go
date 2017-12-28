package main

import "fmt"

func main(){
	input := []byte{63,144,180,149,1,255,167,84,125,65,188,0,2,254,229,24}
	
	list := make([]byte, 256)
	for i := 0 ; i < 256; i++ { list[i] = byte(i) }
	fmt.Println(HashRound(0,0, list,input))

	
	inputString := "63,144,180,149,1,255,167,84,125,65,188,0,2,254,229,24"
	fmt.Println(Hash(inputString))
}

func Hash(input string) string{
	list := make([]byte, 256)
	for i := 0 ; i < 256; i++ { list[i] = byte(i) }

	inputLengths := handleInput(input)
	//fmt.Println(inputLengths)

	currPos, skipSize := 0,0
	for r := 0; r < 64 ; r++{
		_, currPos, skipSize = HashRound(currPos, skipSize, list,inputLengths)

		//fmt.Println(r, currPos, skipSize, list)
	}

	res := denseIt(list)
	return makeString(res)
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