package main

import "fmt"

func main() {
	part1()
	part2()
}

func part1() {
	p1 := 8
	p2 := 3

	s1, s2 := 0, 0

	rolls := 0
	next := 1
	roll := func() int {
		v := next
		next++
		if next > 100 {
			next = 1
		}
		rolls++
		return v
	}

	for {
		p1 += roll() + roll() + roll()
		for p1 > 10 {
			p1 -= 10
		}
		s1 += p1
		if s1 >= 1000 {
			break
		}
		p2 += roll() + roll() + roll()
		for p2 > 10 {
			p2 -= 10
		}
		s2 += p2
		if s2 >= 1000 {
			break
		}
	}

	if s1 >= 1000 {
		fmt.Println(s2 * rolls)
	} else {
		fmt.Println(s1 * rolls)
	}
}

func part2() {
	dp := [22][22][11][11][2]int{}
	dp[0][0][8][3][0] = 1

	rolls := make([]int, 10)
	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				rolls[r1+r2+r3]++
			}
		}
	}

	w1, w2 := 0, 0

	for s1 := 0; s1 <= 20; s1++ {
		for s2 := 0; s2 <= 20; s2++ {
			for p1 := 1; p1 <= 10; p1++ {
				for p2 := 1; p2 <= 10; p2++ {
					for r := 1; r <= 9; r++ {
						if dp[s1][s2][p1][p2][0] > 0 {
							np1 := p1 + r
							for np1 > 10 {
								np1 -= 10
							}
							ns := s1 + np1
							nc := dp[s1][s2][p1][p2][0] * rolls[r]
							if ns > 20 {
								w1 += nc
							} else {
								dp[ns][s2][np1][p2][1] += nc
							}
						}
						if dp[s1][s2][p1][p2][1] > 0 {
							np2 := p2 + r
							for np2 > 10 {
								np2 -= 10
							}
							ns := s2 + np2
							nc := dp[s1][s2][p1][p2][1] * rolls[r]
							if ns > 20 {
								w2 += nc
							} else {
								dp[s1][ns][p1][np2][0] += nc
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(w1, w2)
}

const input = `Player 1 starting position: 8
Player 2 starting position: 3`
