package main

import (
	"aoc2021/helper"
	"fmt"
	"math"
	"os"
	"strconv"
)

var lazyHexMap = map[rune]string {
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

var versionCount = int64(0)

func parsePacketLiteral(decoded string, pos int) (number int64, length int, err error) {
	i := pos+6
	done := false
	numbers := ""
	for !done {
		// read a packet
		packetData := decoded[i:i+5]
		if packetData[0] == '0' {
			done = true
		}
		numbers += packetData[1:]
		i += 5
	}
	// Determine packet length
	length = i-pos
	number, err = strconv.ParseInt(numbers, 2, 64)
	if err != nil {
		return 0, 0, err
	}
	return number, length, nil
}

func atEnd(s string) bool {
	nonzero := false
	for _, x := range s {
		if x != '0' {
			nonzero = true
			break
		}
	}
	return !nonzero
}

func parsePacket(pos int, decoded string) (int, int64, error) {

	fmt.Printf("Decoded length: %v\n", len(decoded))
	if atEnd(decoded[pos:]) {
		return 0, 0, nil
	}

	idx := pos
	packetVersion, err := strconv.ParseInt(decoded[idx:idx+3], 2, 64)
	if err != nil {
		return 0, 0, err
	}
	fmt.Printf("Found version: %v\n", packetVersion)
	versionCount += packetVersion

	packetId, err := strconv.ParseInt(decoded[idx+3:idx+6], 2, 64)
	if err != nil {
		return 0, 0, err
	}

	if packetId == 4 {
		fmt.Println("Processing literal")
		number, packetLen, err := parsePacketLiteral(decoded, idx)
		number = number
		fmt.Printf("Found literal: %v\n", number)
		if err != nil {
			return 0, 0, err
		}
		idx += packetLen
		return idx, number, nil
	}

	if packetId != 4 {
		lengthTypeId := decoded[idx+6]
		subpacketValues := make([]int64, 0)
		if lengthTypeId == '0' {
			// next 15 bits are total length in bits
			totalLengthInBits, err := strconv.ParseInt(decoded[idx+7:idx+7+15], 2, 64)
			fmt.Printf("Total length in bits: %v\n",totalLengthInBits)
			if err != nil {
				return 0, 0, err
			}

			for n := 0; n < int(totalLengthInBits); {
				pl, num, err := parsePacket(n, decoded[idx+7+15:idx+7+15+int(totalLengthInBits)])
				if err != nil {
					return 0, 0, err
				}
				n = pl
				subpacketValues = append(subpacketValues, num)
			}
			idx += 7 + 15 + int(totalLengthInBits)
		} else {
			// next 11 bits are number of sub-packets
			numSubPackets, err := strconv.ParseInt(decoded[idx+7:idx+7+11], 2, 64)
			if err != nil {
				return 0, 0, err
			}
			pcpos := idx+7+11
			fmt.Printf("Processing %v sub-packets\n", numSubPackets)
			for pc := 0; pc < int(numSubPackets); pc++ {
				nxtidx, num, err := parsePacket(pcpos, decoded)
				if err != nil {
					return 0, 0, err
				}
				pcpos = nxtidx
				subpacketValues = append(subpacketValues, num)
			}
			idx = pcpos
		}

		output := int64(0)
		switch packetId {
		case 0:
			for _, v := range subpacketValues {
				output += v
			}
		case 1:
			output = 1
			for _, v := range subpacketValues {
				output *= v
			}
		case 2:
			output = math.MaxInt64
			for _, v := range subpacketValues {
				if v < output {
					output = v
				}
			}
		case 3:
			output = math.MinInt64
			for _, v := range subpacketValues {
				if v > output {
					output = v
				}
			}
		case 5:
			if subpacketValues[0] > subpacketValues[1] {
				output = 1
			} else {
				output = 0
			}
		case 6:
			if subpacketValues[0] < subpacketValues[1] {
				output = 1
			} else {
				output = 0
			}
		case 7:
			if subpacketValues[0] == subpacketValues[1] {
				output = 1
			} else {
				output = 0
			}
		}
		return idx, output, nil
	}

	return idx, 0, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	lines, err := helper.ReadLines(fh, true)
	if err != nil {
		fmt.Printf("Unable to read input: %v\n", err)
		return
	}

	decoded := ""
	for _, r := range lines[0] {
		decoded += lazyHexMap[r]
	}

	fmt.Println(decoded)
	_, ans, err := parsePacket(0, decoded)
	fmt.Printf("Part one: %v\n", versionCount)
	fmt.Printf("Part two: %v\n", ans)
}
