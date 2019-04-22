package main


import (
	"strings"
	"fmt"
	"strconv"
	"regexp"
	"io/ioutil"
	"sort"
)

func toInt(s string) int{
	v, _ := strconv.Atoi(s)
	return v
}
var re = regexp.MustCompile("-?\\d+")
func getInts(line string) []int{
	parts := re.FindAllString(line, 4)
	pi := make([]int, 4)
	for i, p := range parts{ 
		pi[i] = toInt(p) 
	}
	return pi
}
func readInput(fname string) string {
	file, _ := ioutil.ReadFile(fname)
	return string(file)
}

type Group struct{
	units int
	hp int
	weak []string
	immune []string

	damage int
	damageType string
	initiative int

	isInfection bool
	idx int
}
func (g Group) name() string{
	name := "Immune"
	if g.isInfection { 
		name = "Infection"
	}
	return name + " group " + strconv.Itoa(g.idx)
}
func (attacker Group) possibleDamage(defender *Group) int {
	dmg := attacker.units * attacker.damage
	for _, t := range defender.weak {
		if attacker.damageType == t {
			dmg *= 2
			break
		}
	}

	for _, t := range defender.immune {
		if attacker.damageType == t {
			dmg = 0
			break
		}
	}
	
	//fmt.Println(attacker.name(), "would deal", dmg, " to ", defender.name())
	return dmg
}
func (g *Group) takeDamage(damage int) (bool, int){
	deadUnits := damage / g.hp
	g.units -= deadUnits
	return g.units <= 0, deadUnits
}
func (g Group) effectivePower() int{
	return g.units * g.damage
}
func parseGroup(line string) *Group {
	ints := getInts(line)
	units := ints[0]
	hp := ints[1]
	damage := ints[2]
	initiative := ints[3]

	//parse weak and immune
	start_sub, stop_sub := -1, -1
	for i := 0; i < len(line) && stop_sub == -1 ; i++ {
		if line[i] == '(' { 
			start_sub = i
		} else if line[i] == ')' {
			stop_sub = i
		}
	}

	weak := []string{}
	immune := []string{}
	if start_sub >= 0 && stop_sub >= 0 {
		substring := line[start_sub:stop_sub]
		parties := strings.Split(substring, ";")

		curr := &weak
		for _, party := range parties {
			split_parts := strings.Split(party, " ")
			parts := []string{}
			for _, p := range split_parts { 
				if p == "" { continue }
				parts = append(parts, strings.Trim(p, "();, "))
			}

			if parts[0] == "weak" {
				curr = &weak
			} else if parts[0] == "immune" {
				curr = &immune
			}
			for _, p := range parts[2:]{
				*curr = append(*curr, p)
			}
		}

	}
	//parse damagetype
	parts := strings.Split(line, " ")
	damageType := parts[len(parts)-5]

	return &Group{
		units, hp, weak, immune, damage, damageType, initiative, false, 0,
	}
}
func parseInput(input string) ([]*Group, []*Group){

	immune := []*Group{}
	infection := []*Group{}
	curr := &immune

	for _, line := range strings.Split(input, "\n") {	
		if line == "Immune System:"{
			curr = &immune
		} else if line == "Infection:"{
			curr = &infection
		} else if line == "" {
			continue//nop
		} else {
			*curr = append(*curr, parseGroup(line))
		}
	}

	for i, _ := range immune { 
		immune[i].idx = i+1
		immune[i].isInfection = false 
	}
	for i, _ := range infection { 
		infection[i].idx = i+1
		infection[i].isInfection = true 
	}
	return immune, infection
}


type Fight struct{
	attacker *Group
	defender *Group
}
func (f Fight) exec() (int, bool){

	dmg := f.attacker.possibleDamage(f.defender)
	isDead, units := f.defender.takeDamage(dmg)

	return units, isDead
}

func match(attackers, defenders []*Group) []Fight{
	//sort(decrease effective power, decreasing initiative)
	sort.Slice(attackers, func(i,j int)bool{
		if attackers[i].effectivePower() > attackers[j].effectivePower() {
			return true
		} else if attackers[i].effectivePower() == attackers[j].effectivePower() {
			return attackers[i].initiative > attackers[j].initiative
		} else {
			return false
		}
	})

	//add candidates
	toMatch := make([]*Group, len(defenders))
	for i, _ := range defenders { toMatch[i] = defenders[i] }

	//for each attacker get defender
	fights := make([]Fight, 0, len(attackers))
	for _, attacker := range attackers{

		//find best opponent
		dmg, def := -1, -1
		for i, defender := range toMatch{
			expected_damage := attacker.possibleDamage(defender)

			//skip immune
			immune := false
			for _, o := range defender.immune {
				if attacker.damageType == o {
					immune = true
					break
				}
			}

			if immune {
				//skip immune
				continue
			} else if expected_damage > dmg { //most damage
				dmg = expected_damage
				def = i
			} else if expected_damage == dmg{

				if defender.effectivePower() > toMatch[def].effectivePower() { //power
					dmg = expected_damage
					def = i	
				} else if defender.effectivePower() == toMatch[def].effectivePower() && 
						defender.initiative > toMatch[def].initiative { //initiative
					dmg = expected_damage
					def = i	
				}
			}
		}

		//if we have someone to attack
		if def >= 0 {
			fights = append(fights, Fight{attacker, toMatch[def]})
			//remove from pool
			toMatch = append(toMatch[:def], toMatch[def+1:]...)			
		}
	}

	return fights
}
func targetSelection(immune, infection []*Group) []Fight{
	fights := make([]Fight, 0, len(immune) + len(infection))
	for _, f := range match(immune, infection){
		fights = append(fights, f)
	}
	for _, f := range match(infection, immune){
		fights = append(fights, f)
	}

	return fights
}
func fight(immune, infection []*Group, fights []Fight) ([]*Group, []*Group){
	//sort in attack order
	sort.Slice(fights, func(i,j int)bool{
		return fights[i].attacker.initiative > fights[j].attacker.initiative
	})

	for i := 0; i < len(fights); i++ {

		fight := fights[i]

		_, isDead := fight.exec()
		//fmt.Println(fight.attacker.name(), " attacks ", fight.defender.name(), " killing ", units, isDead)

		if isDead {
			//remove from attackers
			for j := i+1 ; j < len(fights); j++{
				if fights[j].attacker == fight.defender {
					fights = append(fights[:j], fights[j+1:]...)
					break
				}
			}

			//remove from immune/infection
			if fight.defender.isInfection{
				for j := 0; j < len(infection); j++ {
					if infection[j].idx == fight.defender.idx{
						infection = append(infection[:j], infection[j+1:]...)
						break
					}
				}
			} else {
				for j := 0; j < len(immune); j++ {
					if immune[j].idx == fight.defender.idx{
						immune = append(immune[:j], immune[j+1:]...)
						break
					}
				}
			}
		}	
	}

	return immune, infection
}
func round(immune, infection []*Group) ([]*Group, []*Group){

	//fmt.Println("fights")
	fights := targetSelection(immune, infection)
	immune, infection = fight(immune, infection, fights)

	/*
	fmt.Println()
	fmt.Println("immune")
	for _, o := range immune{ fmt.Println(o.name(), " => ", o) }
	fmt.Println("infection")
	for _, o := range infection{ fmt.Println(o.name(), " => ", o) }
	*/
	return immune, infection
}


func part1(immune, infection []*Group) (int,int){
	/*
	fmt.Println("immune")
	for _, o := range immune{ fmt.Println(o.name(), " => ", o) }
	fmt.Println("infection")
	for _, o := range infection{ fmt.Println(o.name(), " => ", o) }
	fmt.Println()
	*/

	last_total_units := 0
	for _, o := range immune { last_total_units += o.units}
	for _, o := range infection { last_total_units += o.units}

	for i := 0 ; len(immune) > 0 && len(infection) > 0; i++ {
		//fmt.Println("round ", i)
		immune, infection = round(immune, infection)

		total_units := 0
		for _, o := range immune { total_units += o.units}
		for _, o := range infection { total_units += o.units}

		if total_units == last_total_units { return 0,0 } //loop
		last_total_units = total_units
		//fmt.Println()
		//fmt.Println()
	}

	sum_immune := 0
	for _, o := range immune {
		sum_immune += o.units
	}

	sum_infection := 0
	for _, o := range infection {
		sum_infection += o.units
	}
	return sum_immune, sum_infection
}

func part2(immune, infection []*Group, boost int) (int,int) {

	for i := 0 ; i < len(immune); i++ {
		immune[i].damage += boost
	}
	return part1(immune, infection)
}

func main(){
	input := readInput("input.txt")
	
	immune, infection := parseInput(input)
	_, part1 := part1(immune, infection)
	fmt.Println("Part 1:", part1)
	
	p2 := 0
	for boost := 0; p2 <= 0 ; boost++{
		immune, infection := parseInput(input)
		p2, _ = part2(immune, infection, boost)

		//fmt.Println(boost, " => ", p2)
	}
	fmt.Println("Part 2:", p2)
	
}