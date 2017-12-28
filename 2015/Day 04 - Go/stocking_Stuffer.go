package main

import (
	"fmt"
	"crypto/md5"
	"strconv"
)

var noThreads = 10

func main() {
	Search("abcdef", Hash)
	Search("pqrstuv", Hash)

	Search("iwrupvqb", Hash)
	Search("iwrupvqb", Hash2)
}

func Search(key string, f func(key string, i int)bool){
	fmt.Println("Searching for ", key)

	inputs := make(chan int, 100)
	output := make(chan int, noThreads)
	done := make(chan int, noThreads)

	//fill up inputs
	go filler(inputs, done)

	//start workers
	for j := 0 ; j < noThreads ; j++ {
		go worker(key, inputs, output, done, f)
	}

	//await outputs
	for res := range output {
		fmt.Println(res)
	}
}

func filler(in chan<- int, done <-chan int){
	for i := 0 ; i >= 0 ; i++ { 
		in <- i
		if len(done) > 0 { close(in) ; return }
	}
}
func worker(key string, in <-chan int, out, done chan<- int, f func(string, int)bool){
	i := <- in
	for !f(key, i){ 
		if len(done) > 0 { return }
		if i % 1000000 == 0 { fmt.Println(i) }
		i = <-in
	}

	out <- i
	done <- i
	close(out)
	return
}

func Hash(key string, i int) bool{
	key += strconv.Itoa(i)
	hash := fmt.Sprintf("%x", 
		md5.Sum([]byte(key)),
	)
	return hash[:5] == "00000"
}
func Hash2(key string, i int) bool{
	key += strconv.Itoa(i)
	hash := fmt.Sprintf("%x", 
		md5.Sum([]byte(key)),
	)
	return hash[:6] == "000000"
}
