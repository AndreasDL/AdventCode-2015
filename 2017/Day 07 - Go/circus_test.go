package main

import (
	"testing"
)

var tests = []struct{
	input string
	output string
}{
	{`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`,
	"tknk"},
}



func TestPart1(t *testing.T){

	for _, test := range tests {

		actual := part1(test.input)
		
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %q returned %s expecting %s.", 
				test.input, 
				actual, 
				test.output,
			)
		}
	}
}

//give it some input!
func TestPart2(t *testing.T){

	for _, test := range tests {

		part2( test.input )
	}
}