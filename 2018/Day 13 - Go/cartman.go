package main


import (
	"fmt"
	"strings"
	"io/ioutil"
)

func readFile(fname string) string{
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}


type Field [][]byte
type Cart struct{
	y, x int
	dy, dx int
	turn int //left = 0, straight = 1, right = 2
	char byte
}

//if statements are too mainstream
var char_dx = map[byte]int { 
	'>'  : 1,
	'<'  : -1,
	'v'  : 0,
	'^'  : 0,
}
var char_dy = map[byte]int { 
	'>'  : 0,
	'<'  : 0,
	'v'  : 1,
	'^'  : -1,
}
var char_moves = map[string]byte {
	">/": '^',
	">\\": 'v',
	"</": 'v',
	"<\\": '^',
	"^/": '>',
	"^\\" : '<',
	"v/": '<',
	"v\\": '>',
}
var turns = map[int]map[byte]byte{ 
	0: { //left
		'^': '<',
		'<': 'v',
		'v': '>',
		'>': '^',
	},
	1: { //straight
		'^': '^',
		'<': '<',
		'v': 'v',
		'>': '>',	
	},
	2: {//right
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	},
}
func (cart *Cart) move(c byte){
	if c == '-' || c == '|' { //nothing
		return 
	} else if c == '/' || c == '\\' {//normal turns
		
		key := string(cart.char) + string(c)
		cart.char = char_moves[key]

	} else if c == '+' { //advanced turns
	
		cart.char = turns[cart.turn][cart.char]
		cart.turn++
		cart.turn %= 3
	} else if c == ' '{
		panic("no rails!")
	}
	cart.dx = char_dx[cart.char]
	cart.dy = char_dy[cart.char]
}

//get field as is, remove all carts from it => we won't need to update it anymore
//in my input there was no cart in a corner at the start
func parseField(s string) (*Field, *map[int]map[int]*Cart){

	lines := strings.Split(s, "\n")

	field := make(Field, len(lines))
	carts := make(map[int]map[int]*Cart)

	for y, line := range lines {
		field[y] = []byte(line)

		carts[y] = map[int]*Cart{}
		for x, c := range line {
			if c == 'v' {
			 	carts[y][x] = &Cart{
			 		y,x,
			 		1,0,
			 		0,
			 		byte(c),
			 	}
			 	field[y][x] = '|'
			
			} else if c == '^' {
				carts[y][x] = &Cart{
			 		y,x,
			 		-1,0,
			 		0,
			 		byte(c),
			 	}
			 	field[y][x] = '|'
			} else if c == '<' {
				carts[y][x] = &Cart{
			 		y,x,
			 		0,-1,
			 		0,
			 		byte(c),
			 	}
			 	field[y][x] = '-'
			} else if c == '>' {
				carts[y][x] = &Cart{
			 		y,x,
			 		0,1,
			 		0,
			 		byte(c),
			 	}
			 	field[y][x] = '-'
			}
		}
	}
	return &field, &carts
}
func printField(field *Field, carts *map[int]map[int]*Cart){
	_field := *field
	

	if carts == nil {
		for _, line := range _field{
			fmt.Println(string(line))
		}
	} else {
		_carts := *carts
		for y, line := range _field {
			for x, c := range line {

				if cart, ex := _carts[y][x] ; ex {
					fmt.Print(string(cart.char))
				} else {
					fmt.Print(string(c))
				}
			}
			fmt.Println()
		}
	}
}


//get next state
func next(field *Field, carts *map[int]map[int]*Cart) *map[int]map[int]*Cart{
	_field := *field
	_carts := *carts

	//init next state
	nextCarts := map[int]map[int]*Cart{}
	for y, _ := range _field { nextCarts[y] = map[int]*Cart{} }

	//move all carts
	for y, line := range _field{
		for x, _ := range line {

			cart, ex := _carts[y][x]
			if !ex { continue }

			//move
			cart.x += cart.dx
			cart.y += cart.dy

			//collision
			_, ex1 := nextCarts[cart.y][cart.x]
			_, ex2 := _carts[cart.y][cart.x]
			if ex1 || ex2 {
				fmt.Println("Part1: ", cart.x, cart.y)
				return nil
				//panic("")
			}

			//move to next position and update
			next_pos := _field[cart.y][cart.x]
			cart.move(next_pos)

			nextCarts[cart.y][cart.x] = cart
		}
	}

	return &nextCarts
}
func part1(){
	/*
example := `/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `

/*	
example := `|
v
|
|
^
|`
*/

//example = "->--<-"

	start := readFile("input.txt")
	field, carts := parseField(start)
	//printField(field, nil)
	//printField(field, carts)
	for i := 0 ; i < 1000 && carts != nil ; i++{
		carts = next(field, carts)
		//printField(field, carts)
		//fmt.Println("")
	}
}


//get next state
func next2(field *Field, carts *map[int]map[int]*Cart) (*map[int]map[int]*Cart, bool){
	_field := *field
	_carts := *carts

	//init next state
	nextCarts := map[int]map[int]*Cart{}
	for y, _ := range _field { nextCarts[y] = map[int]*Cart{} }

	//move all carts
	for y, line := range _field{
		for x, _ := range line {

			cart, ex := _carts[y][x]
			if !ex { 
				continue 
			} else {
				//remove to avoid wrong removals later on (->--<< => --><<- => ---<--)
				delete(_carts[y], x)
			}

			//move
			cart.x += cart.dx
			cart.y += cart.dy

			//collision
			_, ex1 := nextCarts[cart.y][cart.x]
			_, ex2 := _carts[cart.y][cart.x]
			if ex1 || ex2 {
				fmt.Println("Crash: ", cart.x, cart.y)
				if ex1 { delete(nextCarts[cart.y], cart.x) }
				if ex2 { delete(_carts[cart.y], cart.x)}
				continue // => remove carts
			}

			//move to next position and update
			next_pos := _field[cart.y][cart.x]
			cart.move(next_pos)

			nextCarts[cart.y][cart.x] = cart
		}
	}

	var last_cart *Cart
	cart_count := 0
	for _, line := range nextCarts{
		for _, cart := range line {
			cart_count++
			last_cart = cart
		}
	}

	if cart_count == 1 {
		fmt.Println("Part2:", last_cart.x, last_cart.y)
	} 

	return &nextCarts, cart_count <= 1
}
func part2(){
	example := `/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/`

//example = "->--<<"

	input := example
	input = readFile("input.txt")
	field, carts := parseField(input)
	//printField(field, carts)

	var done bool
	for i := 0; /*i < 1000000 &&*/ !done ; i++ {
		carts, done = next2(field, carts)
		//printField(field, carts)
	}
}


func main(){

	part1()
	part2()

}