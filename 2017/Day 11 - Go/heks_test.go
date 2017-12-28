package main

import (
	"testing"
)

var tests = []struct{
	input string
	output int
}{
	{
		input: "ne,ne,ne",
		output: 3,	
	},{
		input: "ne,ne,sw,sw",
		output: 0,	
	},{
		input: "ne,ne,s,s",
		output: 2,	
	},{
		input: "se,sw,se,sw,sw",
		output: 3,
	},
}

func TestHashRound(t *testing.T){
	
	for _, test := range tests {

		actual := Hexed(test.input)
		
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %s returned %d expecting %d.", 
				test.input,
				actual, 
				test.output,
			)
		}
	}
}