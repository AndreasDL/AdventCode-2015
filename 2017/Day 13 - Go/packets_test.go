package main

import (
	"testing"
)

var tests = []struct{
	input string
	output int
	output2 int
}{
	{
		input: `0: 3
1: 2
4: 4
6: 4`,
		output: 24,
		output2: 10,
	},
}
/*
func TestPart1(t *testing.T){
	
	for _, test := range tests {

		actual := Part1(test.input)
		
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %s returned %d expecting %d.", 
				test.input,
				actual, 
				test.output,
			)
		}
	}
}

func TestPart1Input(t *testing.T){

	input := readInput("input.txt")
	actual := Part1(input)
	if actual != 2160 {
		t.Fatalf("%d instead of 2160", actual)
	}
}*/

func TestPart2(t *testing.T){
	
	for _, test := range tests {

		actual := Part2(test.input)
		
		if actual != test.output2 {
			t.Fatalf("Mistakes were made.. %s returned %d expecting %d.", 
				test.input,
				actual, 
				test.output2,
			)
		}
	}
}