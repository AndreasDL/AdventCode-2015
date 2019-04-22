package main


import (
	"testing"
)

func compStates(a,b *State) bool{

	if (*a).e1 != (*b).e1 || (*a).e2 != (*b).e2 { return false }
	return compSlices(a.scores, b.scores)

}


func Test_next(t *testing.T){

	cases := []State{
		State{  0, 1, []byte{3, 7 }},
		State{  0, 1, []byte{3, 7, 1, 0 }},
		State{  4, 3, []byte{3, 7, 1, 0, 1, 0 }},
		State{  6, 4, []byte{3, 7, 1, 0, 1, 0, 1 }},
		State{  0, 6, []byte{3, 7, 1, 0, 1, 0, 1, 2 }},
		State{  4, 8, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4 }},
		State{  6, 3, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5 }},
		State{  8, 4, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1 }},
		State{  1, 6, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5 }},
		State{  9, 8, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8 }},
		State{  1,13, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9 }},
		State{  9, 7, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9, 1, 6 }},
		State{ 15,10, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9, 1, 6, 7 }},
		State{  4,12, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9, 1, 6, 7, 7 }},
		State{  6, 2, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9, 1, 6, 7, 7, 9 }},
		State{  8, 4, []byte{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9, 1, 6, 7, 7, 9, 2}},
	}


	scores := make([]byte, 0, 20)
	scores = append(scores, 3, 7)
	state := State{ 0,1, scores }

	for i, expected := range cases{

		if !compStates(&state, &expected){
			t.Fatalf(
				"at %d => %v should be %v",
				i,
				state,
				expected,
			)
		}

		state.next()
	}
}


func Test_part1(t *testing.T){
	cases := []struct{
		n int
		res []byte
	}{
		{    9, []byte{5,1,5,8,9,1,6,7,7,9}},
		{    5, []byte{0,1,2,4,5,1,5,8,9,1}},
		{   18, []byte{9,2,5,1,0,7,1,0,8,5}},
		{ 2018, []byte{5,9,4,1,4,2,9,8,8,2}},
	}

	for _, c := range cases {
		actual := part1(c.n)

		if !compSlices(actual,c.res){
			t.Fatalf(
				"%d => %v should be %v",
				c.n,
				actual,
				c.res,
			)
		}
	}
}

func Test_part2(t *testing.T){
	cases := []struct{
		searched []byte
		n int
	}{
{ []byte{5,1,5,8,9}, 9},
{ []byte{0,1,2,4,5}, 5},
{ []byte{9,2,5,1,0}, 18},
{ []byte{5,9,4,1,4}, 2018},
	}

	for _, c := range cases {
		actual := part2(c.searched)

		if actual != c.n{
			t.Fatalf(
				"%v, %d should be %d",
				c.searched,
				actual,
				c.n,
			)
		}
	}
}