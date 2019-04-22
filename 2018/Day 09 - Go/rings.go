package main


import(
	"fmt"
	"container/ring" //https://golang.org/pkg/container/ring/
)

//Ring
func insert(r *ring.Ring, turn int) *ring.Ring{
	new_ring := ring.New(1)
	new_ring.Value = turn

	r = r.Next()
	r.Link(new_ring)
	r = r.Next()

	return r
}
func remove(r *ring.Ring) (*ring.Ring, int){
	//fmt.Println()
	r = r.Move(-7)
	points := r.Value.(int)
	//showRing(r)
	r = r.Prev()
	r.Unlink(1)
	r = r.Next()
	//showRing(r)
	//fmt.Println()
	return r, points
}
func toList(r *ring.Ring) *[]int{
	res := make([]int, r.Len())
	for i := 0; i < r.Len() ; i++ {
		res[i] = r.Value.(int)
		r = r.Next()
	}
	return &res
}
func showRing(r *ring.Ring){
	// Iterate through the ring and print its contents
	values := toList(r)
	fmt.Println("curr ", r.Value, " => ", *values)
}

//game
func finishRing(player_cnt, marble_cnt int) *[]int{
	//initialise scores
	scores := make([]int, player_cnt)

	//initialise ring
	r := ring.New(1)
	r.Value = 0
	//showRing(r)
	for turn := 1; turn <= marble_cnt; turn++ {

		if turn%23 == 0 {
			//get current marble as score
			player := (turn-1+player_cnt) %player_cnt
			scores[player] += turn

			//remove additional marble, update current & get points
			var points int
			r, points = remove(r)
			scores[player]+= points
		} else {
			r = insert(r, turn)	
		}
		//showRing(r)
	}

	return &scores
}
func getMaxScore(scores *[]int) (int, int){

	max_s, max_p := -1, -1
	for p, s := range *scores {
		if s > max_s{
			max_s = s
			max_p = p
		}
	}

	return max_s, max_p
}


//exercises
func part1(player_cnt, marble_cnt int) int{

	scores := finishRing(player_cnt, marble_cnt)
	winning_score, _ := getMaxScore(scores)
	return winning_score
}
func part2(player_cnt, marble_cnt int) int{
	return part1(player_cnt, marble_cnt*100)
}



func main(){
	player_cnt := 465
	marble_cnt := 71940
	fmt.Println(part1(player_cnt, marble_cnt))
	fmt.Println(part2(player_cnt, marble_cnt))
}