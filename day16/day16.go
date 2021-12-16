package main

import (
	// "adventofcodego/utils/inputs"
	"adventofcodego/utils/utils"
	"fmt"
)

var DAY int = 16

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func parseHex(input string) []bool {
	res := make([]bool, len(input)*4)
	count := 0
	var v int
	for _, r := range input {
		if '0' <= r && r <= '9' {
			v = int(r - '0')
		} else if 'A' <= r && r <= 'F' {
			v = int(r-'A') + 10
		}
		for i := 0; i < 4; i++ {
			res[count] = v>>(4-i-1)&1 == 1
			count++
		}
	}
	return res
}

func packetLiteral(bits []bool) (value int, read int) {
	i := 0
	value = 0
	for {
		for j := 1; j < 5; j++ {
			value = value<<1 + boolToInt(bits[i+j])
		}
		if bits[i] == false {
			i += 5
			break
		}
		i += 5
	}
	return value, i
}

func printBinary(bits []bool) {
	for i := range bits {
		fmt.Print(boolToInt(bits[i]))
	}
	fmt.Println()
}

func binToInt(bits []bool) int {
	value := 0
	for i := range bits {
		value = value<<1 + boolToInt(bits[i])
	}
	return value
}

func parsePacket(bits []bool) (int, int, int) {
	i := 0
	version := 0
	ptype := 0
	value := 0
	read := 0

	for ; i < 3; i++ {
		version = version<<1 + boolToInt(bits[i])
	}
	for ; i < 6; i++ {
		ptype = ptype<<1 + boolToInt(bits[i])
	}

	if ptype == 4 {
		value, read = packetLiteral(bits[i:])
		i += read
	} else {
		if !bits[i] {
			i++
			bitlen := binToInt(bits[i : i+15])
			i += 15
			for j := 0; j < bitlen; {
				vers, val, c := parsePacket(bits[i:])
				version += vers
				value += val
				j += c
				i += c
			}
		} else {
			i++
			packetcount := binToInt(bits[i : i+11])
			i += 11

			for j := 0; j < packetcount; j++ {
				vers, val, c := parsePacket(bits[i:])
				version += vers
				value += val
				i += c
			}
		}
	}

	return version, value, i
}

func part1(input string) interface{} {
	bits := parseHex(input)
	version, _, _ := parsePacket(bits)
	return version
}

func part2(input string) interface{} {
	return nil
}

func main() {
	utils.Solve(part1, part2, DAY)
}
