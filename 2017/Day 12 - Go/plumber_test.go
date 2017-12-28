package main

import (
	"testing"
)

var tests = []struct{
	input string
	output int
	groups int
}{
	{
		input: `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`,
		output: 6,
		groups: 2,
	},
}

func TestInGroup(t *testing.T){
	
	for _, test := range tests {

		actual := InGroup(test.input)
		
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %s returned %d expecting %d.", 
				test.input,
				actual, 
				test.output,
			)
		}
	}
}

func TestCountGroups(t *testing.T){
	
	for _, test := range tests {

		actual := CountGroups(test.input)
		
		if actual != test.groups {
			t.Fatalf("Mistakes were made.. %s returned %d expecting %d.", 
				test.input,
				actual, 
				test.groups,
			)
		}
	}
}

