package main

import (
	"testing"
)


var tests = []struct{
	input string
	output int
}{
	{`b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`,
	1 },
}


func TestRegedit(t *testing.T){
	
	for _, test := range tests {

		actual := Regedit(test.input)
		
		if actual != test.output {
			t.Fatalf("Mistakes were made.. %q returned %d expecting %d.", 
				test.input, 
				actual, 
				test.output,
			)
		}
	}
}