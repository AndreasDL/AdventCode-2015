package main 

import "testing"

var testsP1 = []struct{
	input string
	output int
}{
	{"(())"   ,  0 },
	{"()()"   ,  0 },
	{"((("    ,  3 },
	{"(()(()(",  3 },
	{"))(((((",  3 },
	{"())"    , -1 },
	{"))("    , -1 },
	{")))"    , -3 },
	{")())())", -3 },
}

func TestPart1(t *testing.T){

	for _, test := range testsP1{

		actual := Part1(test.input)
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

