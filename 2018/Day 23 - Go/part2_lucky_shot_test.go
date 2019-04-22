package main


import (
	"testing"
)



func TestWalkInRange(t *testing.T){

	var cases = []struct{
		bot Nanobot
		loc Coord
	}{
		{
			Nanobot{ Coord{10,5,0}, 6 },
			Coord{0,0,0},
		},{
			Nanobot{ Coord{11,5,0}, 6 },
			Coord{0,0,0},
		},{
			Nanobot{ Coord{10,5,0}, 2 },
			Coord{0,0,0},
		},{
			Nanobot{ Coord{3,5,9}, 1 },
			Coord{1,-3,3},
		},
	}

	for _, c := range cases {

		if c.bot.inRangeOf(c.loc) {
			t.Fatalf(
				"%v should not be in range of %v",
				c.loc,
				c.bot,
			)
		}

		next_loc := c.loc.moveInRange(c.bot)

		if !c.bot.inRangeOf(next_loc){
			t.Fatalf(
				"%v moved to %v and should now be in range of %v",
				c.loc,
				next_loc,
				c.bot,
			)
		}
	}
}