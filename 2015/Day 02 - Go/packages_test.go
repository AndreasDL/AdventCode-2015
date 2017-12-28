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

var testsP2 = []struct{
	input string
	output int
}{
	{"2x3x4"   , 34 },
	{"1x1x10"  , 14 },
	{"4x2x3"   , 34 },
	{"1x10x1"  , 14 },
}

func TestRibbon(t *testing.T){

	for _, test := range testsP2{

		actual := Ribbon(test.input)
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %q returned %d expecting %d.", 
				test.input, 
				actual, 
				test.output,
			)
		}
	}
}
