package main

import (
	"testing"
)

var tests = []struct{
	input string
	output int
	regions int 
}{
	{
		input: "flqrgnkx",
		output: 8108,
		regions: 1242,
	},{
		input: "hfdlxzhv",
		output: 8230,
		regions: 1103,
	},
}

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


func TestPart2(t *testing.T){
	
	for _, test := range tests {

		actual := Part2(test.input)
		
		if actual != test.regions {
			t.Fatalf("Mistakes were made.. %s returned %d expecting %d.", 
				test.input,
				actual, 
				test.regions,
			)
		}
	}
}

