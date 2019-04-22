package main


import (
	"testing"
)

var checksum_inputs = []struct{
	text string
	isDouble bool
	isTriple bool
}{
	{
		text: "abcdef",
		isDouble: false,
		isTriple: false,
	},{
		text: "bababc",
		isDouble: true,
		isTriple: true,
	},{
		text: "abbcde",
		isDouble: true,
		isTriple: false,
	},{
		text: "abcccd", 
		isDouble: false,
		isTriple: true,
	},{
		text: "aabcdd", 
		isDouble: true,
		isTriple: false,
	},{
		text: "abcdee", 
		isDouble: true,
		isTriple: false,
	},{
		text: "ababab", 
		isDouble: false,
		isTriple: true,
	},
}


func TestIsDouble(t *testing.T){
	for _, i := range checksum_inputs{
		actual := isDouble( &i.text)
		if actual != i.isDouble {
			t.Fatalf(
				"Failure detecting doubles at %s output %t should be %t",
				i.text,
				actual,
				i.isDouble,
			)
		}
	}
}


func TestIsTriple(t *testing.T){
	for _, i := range checksum_inputs{
		actual := isTriple( &i.text)
		if actual != i.isTriple {
			t.Fatalf(
				"Failure detecting Triple at %s output %t should be %t",
				i.text,
				actual,
				i.isTriple,
			)
		}
	}
}

func TestChecksum(t *testing.T){

	input := make([]string, len(checksum_inputs))
	for i, v := range checksum_inputs {
		input[i] = v.text
	}

	expected := 12
	actual := calcChecksum(&input)
	if actual != expected {
		t.Fatalf(
			"checksum is %x should be %x",
			actual,
			expected,
		)
	}
}