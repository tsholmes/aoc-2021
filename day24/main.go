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

func evalInput(inputs []int, start int, startz int, count int) int {
	var vars [4]int
	vars[3] = startz

	lval := func(op operation) int {
		return vars[op.lvar]
	}
	rval := func(op operation) int {
		if op.rconst {
			return op.rconstval
		}
		return vars[op.rvar]
	}
	for _, op := range ops[18*start : 18*(start+count)] {
		vl, vr := lval(op), rval(op)
		switch op.op {
		case OpInp:
			vars[op.lvar] = inputs[0]
			inputs = inputs[1:]
		case OpAdd:
			vars[op.lvar] = vl + vr
		case OpMul:
			vars[op.lvar] = vl * vr
		case OpDiv:
			if vr == 0 {
				panic("WAT")
			}
			vars[op.lvar] = vl / vr
		case OpMod:
			if vl < 0 || vr <= 0 {
				panic("WAT")
			}
			vars[op.lvar] = vl % vr
		case OpEql:
			if vl == vr {
				vars[op.lvar] = 1
			} else {
				vars[op.lvar] = 0
			}
		}
	}
	return vars[3]
}

func tostr(ns []int) string {
	res := make([]byte, len(ns))
	for i, n := range ns {
		res[i] = byte(n + '0')
	}
	return string(res)
}

func part1() {
	const valCap = 26 * 26 * 26
	curVals := map[int]string{}
	curVals[0] = ""

	doSegment := func(start, end int) {
		nextVals := map[int]string{}
		for sz, cs := range curVals {
			each(end-start, func(vals []int) {
				v := evalInput(vals, start, sz, end-start)
				if v < valCap {
					s := cs + tostr(vals)
					if es, ok := nextVals[v]; !ok || s > es {
						nextVals[v] = s
					}
				}
			})
		}
		curVals = nextVals
	}

	for i := 0; i < 14; i++ {
		doSegment(i, i+1)
	}
	fmt.Println(curVals[0])
}

func part2() {
	const valCap = 26 * 26 * 26
	curVals := map[int]string{}
	curVals[0] = ""

	doSegment := func(start, end int) {
		nextVals := map[int]string{}
		for sz, cs := range curVals {
			each(end-start, func(vals []int) {
				v := evalInput(vals, start, sz, end-start)
				if v < valCap {
					s := cs + tostr(vals)
					if es, ok := nextVals[v]; !ok || s < es {
						nextVals[v] = s
					}
				}
			})
		}
		curVals = nextVals
	}

	for i := 0; i < 14; i++ {
		doSegment(i, i+1)
	}
	fmt.Println(curVals[0])
}

func each(n int, f func([]int)) {
	vals := make([]int, n)
	var search func(i int)
	search = func(i int) {
		if i == n {
			f(vals)
			return
		}
		for j := 1; j <= 9; j++ {
			vals[i] = j
			search(i + 1)
		}
	}
	search(0)
}

const (
	OpInp = iota
	OpAdd
	OpMul
	OpDiv
	OpMod
	OpEql
)

type operation struct {
	op int

	lvar int

	rconst    bool
	rconstval int

	rvar int
}

var ops []operation

func init() {
	varNums := map[string]int{
		"w": 0,
		"x": 1,
		"y": 2,
		"z": 3,
	}
	opNames := map[string]int{
		"inp": OpInp,
		"add": OpAdd,
		"mul": OpMul,
		"div": OpDiv,
		"mod": OpMod,
		"eql": OpEql,
	}
	for _, line := range strings.Split(input, "\n") {
		ps := strings.Split(line, " ")
		var op operation
		op.op = opNames[ps[0]]
		if ps[0] == "inp" {
			op.lvar = varNums[ps[1]]
		} else {
			op.lvar = varNums[ps[1]]
			if n, ok := varNums[ps[2]]; ok {
				op.rvar = n
			} else {
				op.rconst = true
				op.rconstval, _ = strconv.Atoi(ps[2])
			}
		}
		ops = append(ops, op)
	}
}

const input = `inp w
mul x 0
add x z
mod x 26
div z 1
add x 12
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 15
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 14
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 12
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 15
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -9
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 12
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -7
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 15
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 2
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -1
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 11
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -16
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 15
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 11
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 10
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -15
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 2
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 10
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 0
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 1
add x 12
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 0
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x -4
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 15
mul y x
add z y
inp w
mul x 0
add x z
mod x 26
div z 26
add x 0
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y 15
mul y x
add z y`
