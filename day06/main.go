package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := strings.Split(input, ",")
	var counts [9]int
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		counts[i]++
	}

	for i := 0; i < 80; i++ {
		var next [9]int
		next[8] += counts[0]
		next[6] += counts[0]
		for i := 0; i < 8; i++ {
			next[i] += counts[i+1]
		}
		counts = next
	}
	res := 0
	for i := 0; i < 9; i++ {
		res += counts[i]
	}
	fmt.Println(res)
}

func part2() {
	lines := strings.Split(input, ",")
	var counts [9]int
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		counts[i]++
	}

	for i := 0; i < 256; i++ {
		var next [9]int
		next[8] += counts[0]
		next[6] += counts[0]
		for i := 0; i < 8; i++ {
			next[i] += counts[i+1]
		}
		counts = next
	}
	res := 0
	for i := 0; i < 9; i++ {
		res += counts[i]
	}
	fmt.Println(res)
}

const input = `1,4,1,1,1,1,1,1,1,4,3,1,1,3,5,1,5,3,2,1,1,2,3,1,1,5,3,1,5,1,1,2,1,2,1,1,3,1,5,1,1,1,3,1,1,1,1,1,1,4,5,3,1,1,1,1,1,1,2,1,1,1,1,4,4,4,1,1,1,1,5,1,2,4,1,1,4,1,2,1,1,1,2,1,5,1,1,1,3,4,1,1,1,3,2,1,1,1,4,1,1,1,5,1,1,4,1,1,2,1,4,1,1,1,3,1,1,1,1,1,3,1,3,1,1,2,1,4,1,1,1,1,3,1,1,1,1,1,1,2,1,3,1,1,1,1,4,1,1,1,1,1,1,1,1,1,1,1,1,2,1,1,5,1,1,1,2,2,1,1,3,5,1,1,1,1,3,1,3,3,1,1,1,1,3,5,2,1,1,1,1,5,1,1,1,1,1,1,1,2,1,2,1,1,1,2,1,1,1,1,1,2,1,1,1,1,1,5,1,4,3,3,1,3,4,1,1,1,1,1,1,1,1,1,1,4,3,5,1,1,1,1,1,1,1,1,1,1,1,1,1,5,2,1,4,1,1,1,1,1,1,1,1,1,1,1,1,1,5,1,1,1,1,1,1,1,1,2,1,4,4,1,1,1,1,1,1,1,5,1,1,2,5,1,1,4,1,3,1,1`
