package main

import (
	"testing"
	"container/ring"
)

func compareSlices(a, b *[]int) bool{

	_a := *a
	_b := *b

	if len(_a) != len(*b){
		return false
	}

	for i, val := range _a {
		if val != _b[i]{
			return false
		}
	}
	return true
}

func Test_insert(t *testing.T){

	r := ring.New(1)
	r.Value = 0

	results := [][]int{
		{1  ,0},
		{2  ,1,  0},
		{3  ,0,  2,  1},
		{4  ,2,  1,  3,  0},
		{5  ,1,  3,  0,  4,  2},
		{6  ,3,  0,  4,  2,  5,  1},
		{7  ,0,  4,  2,  5,  1,  6,  3},
		{8  ,4,  2,  5,  1,  6,  3,  7,  0},
		{9  ,2,  5,  1,  6,  3,  7,  0,  8,  4},
		{10 ,5,  1,  6,  3,  7,  0,  8,  4,  9,  2},
	}

	for i, expected := range results {
		r = insert(r, i+1)

		actual := *toList(r)
		if ! compareSlices(&actual,&expected) {
			t.Fatalf(
				"iteration %d gave %v should be %v",
				i+1,
				actual,
				expected,
			)
		}	
	}
}

func Test_part1(t *testing.T){
	var tests = []struct{
		player_cnt int
		marble_cnt int
		highscore int
	}{
		{
			9, 25, 32,
		},{
			10, 1618, 8317,
		},{
			13, 7999, 146373,
		},{
			17, 1104, 2764,
		},{
			21, 6111, 54718,
		},{
			30, 5807, 37305,
		},
	}


	for _, game := range tests{
		actual := part1(game.player_cnt, game.marble_cnt)

		if actual != game.highscore {
			t.Fatalf(
				"Failed %d players and %d marbles gave %d, should be %d",
				game.player_cnt,
				game.marble_cnt,
				actual,
				game.highscore,
			)
		}
	}
}