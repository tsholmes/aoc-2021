package main

import (
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Println(
		6 +
			60 +
			3 +
			6000 +
			500 +
			7000 +
			30 +
			40 +
			700 +
			3 +
			8,
	)
}

func part2() {
	initState := [16][2]int{
		{3, 0}, {3, 2}, {4, 1}, {4, 3},
		{1, 0}, {2, 2}, {3, 1}, {4, 0},
		{1, 3}, {2, 0}, {2, 1}, {4, 2},
		{1, 1}, {1, 2}, {2, 3}, {3, 3},
	}

	minScore := 100000

	costs := [4]int{1, 10, 100, 1000}

	ents := [][2]int{
		{0, 2}, {0, 4}, {0, 6}, {0, 8},
	}

	seenState := map[[16][2]int]bool{}

	var search func([16][2]int, int)
	search = func(state [16][2]int, score int) {
		if seenState[state] {
			return
		}
		seenState[state] = true
		if score >= minScore {
			return
		}

		{
			done := true
			for i := 0; i < 16; i++ {
				if state[i][0]-1 != i/4 {
					done = false
					break
				}
			}
			if done {
				minScore = score
				return
			}
		}

		open := [4]bool{true, true, true, true}
		for i := 0; i < 16; i++ {
			exp := i / 4
			if state[i][0] > 0 && state[i][0] != exp+1 {
				open[state[i][0]-1] = false
			}
		}

		seen := map[[2]int]bool{}
		for i := 0; i < 16; i++ {
			seen[state[i]] = true
		}

		for i := 0; i < 16; i++ {
			exp := i / 4
			if open[exp] && state[i][0] == exp+1 {
				// in place in the right room
				continue
			}

			if state[i][0] == 0 {
				// current in hallway, try move to room
				if !open[exp] {
					// if we can't move it into the right room, we have to leave it
					continue
				}
				path := true
				steps := 0

				hp := state[i][1]
				hdir := 1
				if hp > ents[exp][1] {
					hdir = -1
				}
				for hp != ents[exp][1] {
					hp += hdir
					steps++
					if seen[[2]int{0, hp}] {
						path = false
						break
					}
				}
				if !path {
					continue
				}
				steps++
				hp = 0
				for hp < 3 && !seen[[2]int{exp + 1, hp + 1}] {
					hp++
					steps++
				}
				nstate := state
				nstate[i] = [2]int{exp + 1, hp}
				search(nstate, score+steps*costs[exp])
			} else {
				// currently in wrong room (or on top of bad), move to all storage spots
				canup := true
				for p := state[i][1] - 1; p >= 0; p-- {
					if seen[[2]int{state[i][0], p}] {
						canup = false
						break
					}
				}
				if !canup {
					continue
				}
				usteps := 1 + state[i][1]
				shp := ents[state[i][0]-1][1]
				// move left
				{
					hp := shp
					steps := 0
					for {
						hp--
						steps++
						if hp < 0 {
							break
						}
						if seen[[2]int{0, hp}] {
							break
						}
						if hp < 2 || hp&1 == 1 {
							nstate := state
							nstate[i] = [2]int{0, hp}
							search(nstate, score+(usteps+steps)*costs[exp])
						}
					}
				}
				// move right
				{
					hp := shp
					steps := 0
					for {
						hp++
						steps++
						if hp > 10 {
							break
						}
						if seen[[2]int{0, hp}] {
							break
						}
						if hp > 8 || hp&1 == 1 {
							nstate := state
							nstate[i] = [2]int{0, hp}
							search(nstate, score+(usteps+steps)*costs[exp])
						}
					}
				}
			}
		}
	}

	search(initState, 0)
	fmt.Println(minScore)
}

const input = `#############
#...........#
###B#C#A#B###
  #C#D#D#A#
  #########`

const input2 = `#############
#...........#
###B#C#A#B###
	#D#C#B#A#
	#D#B#A#C#
  #C#D#D#A#
	#########
`

var _, _ = input, input2
