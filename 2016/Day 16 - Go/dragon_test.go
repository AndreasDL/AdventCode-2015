package main

import "testing"

var tests = []struct{
	input, output string
}{
	{"1", "100"},
	{"0", "001"},
	{"11111", "11111000000"},
	{"111100001010", "1111000010100101011110000"},
	{"10000", "10000011110"},
	{"10000011110", "10000011110010000111110"},
}


func TestStep(t *testing.T){
	for _, test := range tests {
		actual := string(step( []byte(test.input) ))
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %s returned %s expecting %s.", 
				test.input,
				actual, 
				test.output,
			)
		}
	}
}

var checkTests = []struct{
	input, step, output string
}{
	{"110010110100", "110101", "100"},
	{"10000011110010000111", "0111110101", "01100"},
}

func TestCheckStep(t *testing.T){
	for _, test := range checkTests{
		actual := string(checkStep( []byte(test.input) ))
		if actual != test.step {
			t.Fatalf("Mistakes were made.. %s returned %s expecting %s.", 
				test.input,
				actual, 
				test.step,
			)
		}
	}
}

func TestCheckSum(t *testing.T){
	
	for _, test := range checkTests {
		actual := checkSum( []byte(test.input) )
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %s returned %s expecting %s.", 
				test.input,
				actual, 
				test.output,
			)
		}
	}
}

var testsPrt1 = []struct{
	length int
	input, output string
}{
	{20, "10000", "01100"},
}


func TestPart1(t *testing.T){
	for _, test := range testsPrt1 {
		actual := Part1(test.input, test.length)
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %s returned %s expecting %s.", 
				test.input,
				actual, 
				test.output,
			)
		}
	}
}
