package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	raw, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	dc := &bitDecoder{raw: raw}

	vsum := 0

	var parsePacket func()
	parsePacket = func() {
		version := dc.take(3)
		typ := dc.take(3)

		vsum += version

		switch typ {
		case 4:
			num := 0
			for {
				p := dc.take(5)
				num = (num << 4) | (p & 0xf)
				if p&0x10 == 0 {
					break
				}
			}
		default:
			tlenid := dc.take(1)
			if tlenid == 0 {
				blen := dc.take(15)
				start := dc.total
				for dc.total < start+blen {
					parsePacket()
				}
			} else {
				pcount := dc.take(11)
				for i := 0; i < pcount; i++ {
					parsePacket()
				}
			}
		}
	}

	parsePacket()
	fmt.Println(vsum)
}

func part2() {
	raw, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	dc := &bitDecoder{raw: raw}

	var vstack []int

	var parsePacket func()
	parsePacket = func() {
		version := dc.take(3)
		typ := dc.take(3)
		_ = version

		switch typ {
		case 4:
			num := 0
			for {
				p := dc.take(5)
				num = (num << 4) | (p & 0xf)
				if p&0x10 == 0 {
					break
				}
			}
			vstack = append(vstack, num)
		default:
			tlenid := dc.take(1)
			vcount := 0
			if tlenid == 0 {
				blen := dc.take(15)
				start := dc.total
				for dc.total < start+blen {
					parsePacket()
					vcount++
				}
			} else {
				pcount := dc.take(11)
				for i := 0; i < pcount; i++ {
					parsePacket()
				}
				vcount = pcount
			}
			vs := vstack[len(vstack)-vcount:]
			vstack = vstack[:len(vstack)-vcount]
			val := 0
			switch typ {
			case 0: // sum
				for _, v := range vs {
					val += v
				}
			case 1: // product
				val = 1
				for _, v := range vs {
					val *= v
				}
			case 2: // min
				val = vs[0]
				for _, v := range vs[1:] {
					if v < val {
						val = v
					}
				}
			case 3: // max
				val = vs[0]
				for _, v := range vs[1:] {
					if v > val {
						val = v
					}
				}
			case 5: // greater
				if vs[0] > vs[1] {
					val = 1
				}
			case 6: // less
				if vs[0] < vs[1] {
					val = 1
				}
			case 7: // equal
				if vs[0] == vs[1] {
					val = 1
				}
			}
			vstack = append(vstack, val)
		}
	}

	parsePacket()
	fmt.Println(vstack)
}

type bitDecoder struct {
	raw   []byte
	bit   int
	total int
}

func (d *bitDecoder) take(bits int) int {
	v := 0
	for i := 0; i < bits; i++ {
		v = (v << 1) | d.takeBit()
	}
	return v
}

func (d *bitDecoder) takeBit() int {
	if d.bit == 8 {
		d.bit = 0
		d.raw = d.raw[1:]
	}
	shift := 7 - d.bit
	v := d.raw[0] & (1 << shift)
	d.bit++
	d.total++
	return int(v >> shift)
}

const input = `E20D7880532D4E551A5791BD7B8C964C1548CB3EC1FCA41CC00C6D50024400C202A65C00C20257C008AF70024C00810039C00C3002D400A300258040F200D6040093002CC0084003FA52DB8134DE620EC01DECC4C8A5B55E204B6610189F87BDD3B30052C01493E2DC9F1724B3C1F8DC801E249E8D66C564715589BCCF08B23CA1A00039D35FD6AC5727801500260B8801F253D467BFF99C40182004223B4458D2600E42C82D07CC01D83F0521C180273D5C8EE802B29F7C9DA1DCACD1D802469FF57558D6A65372113005E4DB25CF8C0209B329D0D996C92605009A637D299AEF06622CE4F1D7560141A52BC6D91C73CD732153BF862F39BA49E6BA8C438C010E009AA6B75EF7EE53BBAC244933A48600B025AD7C074FEB901599A49808008398142013426BD06FA00D540010C87F0CA29880370E21D42294A6E3BCF0A080324A006824E3FCBE4A782E7F356A5006A587A56D3699CF2F4FD6DF60862600BF802F25B4E96BDD26049802333EB7DDB401795FC36BD26A860094E176006A0200FC4B8790B4001098A50A61748D2DEDDF4C6200F4B6FE1F1665BED44015ACC055802B23BD87C8EF61E600B4D6BAD5800AA4E5C8672E4E401D0CC89F802D298F6A317894C7B518BE4772013C2803710004261EC318B800084C7288509E56FD6430052482340128FB37286F9194EE3D31FA43BACAF2802B12A7B83E4017E4E755E801A2942A9FCE757093005A6D1F803561007A17C3B8EE0008442085D1E8C0109E3BC00CDE4BFED737A90DC97FDAE6F521B97B4619BE17CC01D94489E1C9623000F924A7C8C77EA61E6679F7398159DE7D84C015A0040670765D5A52D060200C92801CA8A531194E98DA3CCF8C8C017C00416703665A2141008CF34EF8019A080390962841C1007217C5587E60164F81C9A5CE0E4AA549223002E32BDCEA36B2E100A160008747D8B705C001098DB13A388803F1AE304600`
