package main


import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
	"sort"
)

var regex = regexp.MustCompile(
	`\[(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2})\] (.*)`,
)

func readLines(fname string) []string{
	file, _ := ioutil.ReadFile(fname)
	return strings.Split(
		string(file),
		"\n",
	)
}
func parseLine(line string) (int, string){
	parsed := regex.FindStringSubmatch(line)
	minutes, _ := strconv.Atoi(parsed[5])
	text := parsed[6]

	return minutes, text
}
func handleLines(_times *map[int]map[int]int, lines []string){

	times := *_times
	start, current_guard := -1, -1
	

	for _, line := range lines {
	
		minutes, text := parseLine(line)
		parts := strings.Split(text, " ")
	
		switch parts[0]{
		case "Guard":
			current_guard, _ = strconv.Atoi(parts[1][1:])
		case "falls":
			start = minutes
		case "wakes":

			if _, ok := times[current_guard] ; !ok{
				times[current_guard] = make(map[int]int)
			}

			//store details
			for i := start; i < minutes; i++{
				times[current_guard][i]++
			}
		}
	}
}
func getStats(workHours *map[int]int) (int, int, int){
	total_time_asleep := 0 //total time sleps
	favorite_minute, favorite_max := -1, -1 //get best minute
	
	for min, cnt := range *workHours {
	
		total_time_asleep += cnt

		if cnt > favorite_max {
			favorite_minute = min
			favorite_max = cnt
		}
	}

	return total_time_asleep, favorite_minute, favorite_max
}

func main(){
	
	lines := readLines("input.txt")
	sort.Strings(lines)
	
	//get data struct
	times := make(map[int]map[int]int)
	handleLines(&times, lines)

	//loop for max guard
	max_guard, max_time, max_min := -1, -1, -1
	for guard, workHours := range times {

		total_time_asleep, favorite_minute, _ := getStats(&workHours)
		if total_time_asleep > max_time{
			max_time = total_time_asleep
			max_guard = guard
			max_min = favorite_minute
		}
	}
	fmt.Println("Part 1:", max_guard * max_min)

	//loop for max guard
	max_guard, max_min, max_cnt := -1, -1, -1
	for guard, workHours := range times {

		_, minute, cnt := getStats(&workHours)

		if cnt > max_cnt{
			max_guard = guard
			max_cnt = cnt
			max_min = minute
		}
	}
	fmt.Println("Part 2:", max_guard * max_min)

}

