package main

import "fmt"

func main() {
	part1()
	part2()
}

func part1() {
	minx, maxx := 34, 67
	miny, maxy := -215, -186

	high := 0

	ixv := 10
	for iyv := 0; iyv < 10000; iyv++ {
		x, y := 0, 0
		xv, yv := ixv, iyv

		my := 0
		ok := false
		for y > miny {
			x += xv
			y += yv
			if y > my {
				my = y
			}
			if xv > 0 {
				xv--
			} else if xv < 0 {
				xv++
			}
			yv--
			if x >= minx && x <= maxx && y >= miny && y <= maxy {
				ok = true
			}
		}
		if ok && my > high {
			high = my
		}
	}
	fmt.Println(high)
}

func part2() {
	minx, maxx := 34, 67
	miny, maxy := -215, -186

	count := 0
	for ixv := 0; ixv <= 67; ixv++ {
		for iyv := -215; iyv < 10000; iyv++ {
			x, y := 0, 0
			xv, yv := ixv, iyv

			ok := false
			for y > miny {
				x += xv
				y += yv
				if xv > 0 {
					xv--
				} else if xv < 0 {
					xv++
				}
				yv--
				if x >= minx && x <= maxx && y >= miny && y <= maxy {
					ok = true
					break
				}
			}
			if ok {
				count++
			}
		}
	}
	fmt.Println(count)
}

const input = `target area: x=34..67, y=-215..-186`
