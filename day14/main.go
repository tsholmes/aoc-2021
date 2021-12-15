package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := strings.Split(input, "\n")
	pol := lines[0]
	rules := map[string]string{}
	for _, line := range lines[2:] {
		ps := strings.Split(line, " -> ")
		rules[ps[0]] = ps[1]
	}

	for step := 0; step < 10; step++ {
		ins := []string{}
		idxs := []int{}
		for i := 0; i < len(pol)-1; i++ {
			if r, ok := rules[pol[i:i+2]]; ok {
				ins = append(ins, r)
				idxs = append(idxs, i+1)
			}
		}
		next := make([]byte, 0, len(pol)+len(ins))
		for i := 0; i < len(pol); i++ {
			if len(idxs) > 0 && i == idxs[0] {
				next = append(next, ins[0]...)
				idxs = idxs[1:]
				ins = ins[1:]
			}
			next = append(next, pol[i])
		}
		pol = string(next)
	}

	counts := map[byte]int{}
	for _, c := range []byte(pol) {
		counts[c]++
	}
	cs := make([]int, 0, len(counts))
	for _, c := range counts {
		cs = append(cs, c)
	}
	sort.Ints(cs)
	fmt.Println(cs[len(cs)-1] - cs[0])
}

func part2() {
	lines := strings.Split(input, "\n")
	pol := lines[0]
	rules := map[[2]byte]byte{}
	for _, line := range lines[2:] {
		ps := strings.Split(line, " -> ")
		rules[[2]byte{ps[0][0], ps[0][1]}] = ps[1][0]
	}

	ccounts := map[byte]int{}
	pcounts := map[[2]byte]int{}

	for i := 0; i < len(pol); i++ {
		ccounts[pol[i]]++
		if i < len(pol)-1 {
			pcounts[[2]byte{pol[i], pol[(i + 1)]}]++
		}
	}

	for step := 0; step < 40; step++ {
		npcounts := map[[2]byte]int{}
		for k, v := range pcounts {
			if r, ok := rules[k]; ok {
				npcounts[[2]byte{k[0], r}] += v
				npcounts[[2]byte{r, k[1]}] += v
				ccounts[r] += v
			} else {
				npcounts[k] += v
			}
		}
		pcounts = npcounts
	}

	cs := make([]int, 0, len(ccounts))
	for _, c := range ccounts {
		cs = append(cs, c)
	}
	sort.Ints(cs)
	fmt.Println(cs[len(cs)-1] - cs[0])
}

const input = `SHHNCOPHONHFBVNKCFFC

HH -> K
PS -> P
BV -> H
HB -> H
CK -> F
FN -> B
PV -> S
KK -> F
OF -> C
SF -> B
KB -> S
HO -> O
NH -> N
ON -> V
VF -> K
VP -> K
PH -> K
FF -> V
OV -> N
BO -> K
PO -> S
CH -> N
FO -> V
FB -> H
FV -> N
FK -> S
VC -> V
CP -> K
CO -> K
SV -> N
PP -> B
BS -> P
VS -> C
HV -> H
NN -> F
NK -> C
PC -> V
HS -> S
FS -> S
OB -> S
VV -> N
VO -> P
KV -> F
SK -> O
BC -> P
BP -> F
NS -> P
SN -> S
NC -> N
FC -> V
CN -> S
OK -> B
SC -> N
HN -> B
HP -> B
KP -> B
CB -> K
KF -> C
OS -> B
BH -> O
PN -> K
VN -> O
KH -> F
BF -> H
HF -> K
HC -> P
NP -> H
KO -> K
CF -> H
BK -> O
OH -> P
SO -> F
BB -> F
VB -> K
SP -> O
SH -> O
PK -> O
HK -> P
CC -> V
NB -> O
NV -> F
OO -> F
VK -> V
FH -> H
SS -> C
NO -> P
CS -> H
BN -> V
FP -> N
OP -> N
PB -> P
OC -> O
SB -> V
VH -> O
KS -> B
PF -> N
KN -> H
NF -> N
CV -> K
KC -> B`
