package main

import (
	"testing"
)

var tests = []struct{
	size int
	inputlengths []byte
	score int
}{
	{
		size: 5,
		inputlengths: []byte{3,4,1,5},
		score: 12,
	},{
		size: 256,
		inputlengths: []byte{63,144,180,149,1,255,167,84,125,65,188,0,2,254,229,24},
		score: 4480,
	},
}

func TestHashRound(t *testing.T){
	
	for _, test := range tests {

		list := make([]byte, test.size)
		for i := 0 ; i < test.size; i++ { list[i] = byte(i) }	
		actual, _, _ := HashRound(0,0, list, test.inputlengths)
		
		if actual != test.score {
			t.Fatalf("Mistakes were made.. %d returned %d expecting %d.", 
				test.inputlengths,
				actual, 
				test.score,
			)
		}
	}
}


var tests2 =[]struct{
	input string
	output string
}{
	{
		input: "1,2,3",
		output: "3efbe78a8d82f29979031a4aa0b16a9d",
	},{
		input: "",
		output: "a2582a3a0e66e6e86e3812dcb672a272",
	},{
		input: "AoC 2017",
		output: "33efeb34ea91902bb2f59c9920caa6cd",
	},{
		input: "1,2,4",
		output: "63960835bcdc130f0b66d7ff4f6a5a8e",
	},{
		input: "63,144,180,149,1,255,167,84,125,65,188,0,2,254,229,24",
		output: "c500ffe015c83b60fad2e4b7d59dabc4",
	},
}

func TestHash(t *testing.T){
	
	for _, test := range tests2 {

		actual := Hash(test.input)
		
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %s returned %s expecting %s.", 
				test.input,
				actual, 
				test.output,
			)
		}
	}
}
