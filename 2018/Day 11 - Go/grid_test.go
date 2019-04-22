package main


import (
	"testing"
)



func Test_powerlevel(t *testing.T){
	
	cases := []struct{
		x,y int
		serial int
		lvl int
	}{
		{3,5,8,4},
		{122, 79, 57 , -5},
		{217, 196, 39, 0},
		{101, 153, 71, 4},		
	}

	for _, c := range cases {
		actual := powerlevel(c.serial, c.y, c.x)
		if actual != c.lvl{
			t.Fatalf("%d should be %d", actual, c.lvl)
		}
	}
}

func Test_part1(t *testing.T){
	cases := []struct{
		serial int
		x,y int
	}{
		{18, 33,45},
		{42, 21,61},
	}

	for _, c := range cases {
		actual_x, actual_y := part1(c.serial)

		if actual_x != c.x || actual_y != c.y{
			t.Fatalf(
				"%d;%d should be %d;%d",
				actual_x, actual_y,
				c.x, c.y,
			)
		}
	}
}


func Test_part2(t *testing.T){
	cases := []struct{
		serial int
		x,y, size int
	}{
		{18, 90,269,16},
		{42, 232,251,12},
	}

	for _, c := range cases {
		x,y,size := part2(c.serial)

		if x != c.x || y != c.y || size != c.size {
			t.Fatalf(
				"%d;%d;%d should be %d;%d;%d",
				x, y, size,
				c.x, c.y, c.size,
			)
		}
	}
}