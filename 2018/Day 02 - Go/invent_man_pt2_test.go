package main


import (
	"testing"
)

var distanceInputs = []struct{
	a string
	b string
	expected int
	common_letters string
}{
	{
		a: "abcde",
		b: "axcye",
		expected: 2,
		common_letters: "ace",
	}, {
		a: "fghij",
		b: "fguij",
		expected: 1,
		common_letters: "fgij",
	},
}


func testDistance(t *testing.T){

	for _, testcase := range distanceInputs{

		dist, letters := getDistance(&testcase.a, &testcase.b)
		if dist != testcase.expected {
			t.Fatalf(
				"distance between %s and %s = %x, should be %x",
				testcase.a,
				testcase.b,
				dist,
				testcase.expected,
			)
		}

		if letters != testcase.common_letters {
			t.Fatalf(
				"letters wrong %s should be %s",
				letters,
				testcase.common_letters,
			)
		}


	}
}


var letter_inputs = []string{
	"abcde",
	"fghij",
	"klmno",
	"pqrst",
	"fguij",
	"axcye",
	"wvxyz",
}

var expected string = "fgij"

func TestGetLetters(t *testing.T){

	actual := getLetters(&letter_inputs) 
	if actual != expected {
		t.Fatalf(
			"letters not correct, %s should be %s",
			actual,
			expected,
		)
	}
}