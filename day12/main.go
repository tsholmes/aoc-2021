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
	fwd := map[string][]string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		ps := strings.Split(line, "-")
		a, b := ps[0], ps[1]
		fwd[a] = append(fwd[a], b)
		fwd[b] = append(fwd[b], a)
	}

	res := 0
	seen := map[string]bool{}
	var search func(cur string)
	search = func(cur string) {
		if seen[cur] && cur[0] >= 'a' && cur[0] <= 'z' {
			return
		}
		if cur == "end" {
			res++
			return
		}
		seen[cur] = true
		for _, n := range fwd[cur] {
			search(n)
		}
		seen[cur] = false
	}

	search("start")

	fmt.Println(res)
}

func part2() {
	fwd := map[string][]string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		ps := strings.Split(line, "-")
		a, b := ps[0], ps[1]
		fwd[a] = append(fwd[a], b)
		fwd[b] = append(fwd[b], a)
	}

	res := 0
	seen := map[string]bool{}
	var search func(cur string, used bool)
	search = func(cur string, used bool) {
		scur := seen[cur]
		if cur[0] >= 'a' && cur[0] <= 'z' && scur {
			if used || cur == "start" {
				return
			}
			used = true
		}
		if cur == "end" {
			res++
			return
		}
		seen[cur] = true
		for _, n := range fwd[cur] {
			search(n, used)
		}
		seen[cur] = scur
	}

	search("start", false)

	fmt.Println(res)
}

const input = `EG-bj
LN-end
bj-LN
yv-start
iw-ch
ch-LN
EG-bn
OF-iw
LN-yv
iw-TQ
iw-start
TQ-ch
EG-end
bj-OF
OF-end
TQ-start
TQ-bj
iw-LN
EG-ch
yv-iw
KW-bj
OF-ch
bj-ch
yv-TQ`
