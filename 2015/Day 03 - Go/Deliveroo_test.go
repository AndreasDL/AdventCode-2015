package main

import "testing"

var testsP1 = []struct{
	input string
	output int
}{
	{">", 2 },
	{"^>V<", 4},
	{"^v", 2},
	{"^v^v^v^v^v", 2},
}

func TestVisited(t *testing.T){

	for _, test := range testsP1{

		actual := Visited(test.input)
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
	{"^v", 3 },
	{"^>v<", 3},
	{"^v^v^v^v^v", 11},
}

func TestVisited2(t *testing.T){

	for _, test := range testsP2{

		actual := Visited2(test.input)
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %q returned %d expecting %d.", 
				test.input, 
				actual, 
				test.output,
			)
		}
	}
}
