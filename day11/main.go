package main

import (
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := strings.Split(input, "\n")
	cs := make([][]int, len(lines))
	for i, line := range lines {
		cs[i] = make([]int, len(line))
		for j, c := range line {
			cs[i][j] = int(c) - int('0')
		}
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	fcount := 0
	for step := 0; step < 100; step++ {
		for i := 0; i < len(cs); i++ {
			for j := 0; j < len(cs[i]); j++ {
				if cs[i][j] == -1 {
					cs[i][j] = 0
				}
				cs[i][j]++
			}
		}
		var checkFlash func(y, x int, add int)
		checkFlash = func(y, x int, add int) {
			if y < 0 || x < 0 || y >= len(cs) || x >= len(cs[y]) {
				return
			}
			if cs[y][x] == -1 {
				return
			}
			cs[y][x] += add
			if cs[y][x] <= 9 {
				return
			}
			fcount++
			cs[y][x] = -1

			for _, d := range dirs {
				checkFlash(y+d[0], x+d[1], 1)
			}
		}
		for i := 0; i < len(cs); i++ {
			for j := 0; j < len(cs[i]); j++ {
				checkFlash(i, j, 0)
			}
		}
	}
	fmt.Println(fcount)
}

func part2() {
	lines := strings.Split(input, "\n")
	cs := make([][]int, len(lines))
	for i, line := range lines {
		cs[i] = make([]int, len(line))
		for j, c := range line {
			cs[i][j] = int(c) - int('0')
		}
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for step := 0; step < 1000000000; step++ {
		fcount := 0
		for i := 0; i < len(cs); i++ {
			for j := 0; j < len(cs[i]); j++ {
				if cs[i][j] == -1 {
					cs[i][j] = 0
				}
				cs[i][j]++
			}
		}
		var checkFlash func(y, x int, add int)
		checkFlash = func(y, x int, add int) {
			if y < 0 || x < 0 || y >= len(cs) || x >= len(cs[y]) {
				return
			}
			if cs[y][x] == -1 {
				return
			}
			cs[y][x] += add
			if cs[y][x] <= 9 {
				return
			}
			fcount++
			cs[y][x] = -1

			for _, d := range dirs {
				checkFlash(y+d[0], x+d[1], 1)
			}
		}
		for i := 0; i < len(cs); i++ {
			for j := 0; j < len(cs[i]); j++ {
				checkFlash(i, j, 0)
			}
		}
		if fcount == len(cs)*len(cs[0]) {
			fmt.Println(step + 1)
			return
		}
	}
}

const input = `4764745784
4643457176
8322628477
7617152546
6137518165
1556723176
2187861886
2553422625
4817584638
3754285662`
