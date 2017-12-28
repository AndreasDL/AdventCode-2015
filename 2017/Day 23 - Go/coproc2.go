package main

import "fmt"


func main() {
	B, C := 106700, 123700

	ctr := 0
	for i := B ; i <= C ; i+=17 {
		
		if !IsPrime(i) { 
			ctr++ 
			//fmt.Println(i, "is not a prime !", ctr)
		}/* else {
			fmt.Println(i, "is a prime !", ctr)
		}*/

	}

	fmt.Println(ctr)

}


func IsPrime(value int) bool {

	if value % 2 == 0 { return false }

	div := 3 ; for ; div*div < value && value % div != 0 ; div+=2 {}

	return value % div != 0
}