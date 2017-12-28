package main

import "fmt"

func main() {
	fmt.Println(
		Part1("11101000110010100", 272),
	)

	fmt.Println(
		Part1("11101000110010100", 35651584),
	)
}

func Part1(input string, length int) string{

	//example
	a := []byte(input)

	//extend
	for len(a) < length {
		a = step(a)
	}

	//trim
	a = a[:length]

	return checkSum(a)
}

func step(a []byte) []byte{
	
	b := make([]byte, len(a))
	for i, c := range a { b[i] = c }

	b = reverse(b)
	b = swap(b)

	a = append(a, '0')
	a = append(a, b...)
	return a
}
func reverse(s []byte) []byte{
	for i, j := 0 , len(s)-1 ; i < j ; i,j = i+1,j-1{
		s[i], s[j] = s[j], s[i]
	}
	return s
}
func swap(s []byte) []byte{
	for i, c := range s {
		if c == '0' { 
			s[i] = '1'
		} else {
			s[i] = '0'
		}
	}
	return s
}

func checkSum(data []byte) string{
	
	checksum := checkStep(data)
	for len(checksum) % 2 == 0 && len(checksum) > 0 {
		checksum = checkStep(checksum)
	}

	return string(checksum)
}
func checkStep(data []byte) []byte{
	res := make([]byte, int(len(data)/2))

	for i,b,e := 0,0,1 ; e <= len(data) ; i,b,e = i+1,b+2,e+2 {
		
		if data[b] == data[e] {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return res
}