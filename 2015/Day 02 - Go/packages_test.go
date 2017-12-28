package main

import "testing"

var testsP1 = []struct{
	input string
	output int
}{
	{"2x3x4"   , 58 },
	{"1x1x10"  , 43 },
}

func TestSurface(t *testing.T){

	for _, test := range testsP1{

		actual := Surface(test.input)
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %q returned %d expecting %d.", 
				test.input, 
				actual, 
				test.output,
			)
		}
	}
}
/*
var testsP2 = []struct{
	input string
	output int
}{
	{")"    , 1 },
	{"()())", 5 },
}

func TestPart2(t *testing.T){

	for _, test := range testsP2{

		actual := Part2(test.input)
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %q returned %d expecting %d.", 
				test.input, 
				actual, 
				test.output,
			)
		}
	}
}
*/