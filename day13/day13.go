package day13

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
)

type Packet struct {
	Elements []*Packet
	Value    int
}

func (e *Packet) IsList() bool {
	// It is a list of Elements are not nil
	return e.Elements != nil
}

// Used only for testing. Check out day13_test.go
func (e *Packet) Print(stringAccum *string) {
	// Prints the Packet into its original string representation
	if !e.IsList() {
		value := strconv.Itoa(e.Value)
		*stringAccum += value
		return
	}

	*stringAccum += "["
	for i, item := range e.Elements {
		item.Print(stringAccum)

		if i+1 < len(e.Elements) {
			*stringAccum += ","
		}
	}
	*stringAccum += "]"
}

func convertToPacket(packetFragment string) (*Packet, string) {
	// Converts a packet string to a Packet object
	var packet *Packet = &Packet{make([]*Packet, 0), -1}
	number := ""

	// Helper function to convert a string to a integer
	toNumber := func(number string) {
		if number == "" {
			return
		}
		value, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal("Could not convert ", number, " to integer.")
		}
		packet.Elements = append(packet.Elements, &Packet{nil, value})
		number = ""
	}

	for i := 0; i < len(packetFragment); i++ {
		ch := packetFragment[i]

		if ch == '[' {
			if i == 0 {
				// Ignore first bracket
				continue
			}
			// Recursively call converToPacket function until we are find numbers
			var subPacket *Packet
			subPacket, packetFragment = convertToPacket(packetFragment[i:])
			packet.Elements = append(packet.Elements, subPacket)
			i = -1
		} else if ch == ']' {
			// Convert to number if possible and return the list
			toNumber(number)
			return packet, packetFragment[i+1:]
		} else if ch == ',' {
			// Convert to number if possible and return the list
			toNumber(number)
		} else {
			// Add digit to number
			number += string(ch)
		}
	}

	return packet, ""
}

func parsePackets(packetOneStr string, packetTwoStr string) (*Packet, *Packet) {
	// convert both strings to Packet's
	packetOne, _ := convertToPacket(packetOneStr)
	packetTwo, _ := convertToPacket(packetTwoStr)
	return packetOne, packetTwo
}

const (
	InvalidPacket   = 0
	ValidPacket     = 1
	UndecidedPacket = 2
)

func comparePackets(packetOne *Packet, packetTwo *Packet) int {
	if !packetOne.IsList() && !packetTwo.IsList() {
		// Do number vs number comparison
		diff := packetOne.Value - packetTwo.Value

		if diff < 0 {
			return ValidPacket
		} else if diff == 0 {
			return UndecidedPacket
		}
		return InvalidPacket
	} else if packetOne.IsList() && packetTwo.IsList() {
		// Do list vs list comparison
		result := UndecidedPacket

		var idx int = -1
		for idx = range packetTwo.Elements {
			if idx > len(packetOne.Elements)-1 {
				// Left-side ran out of items -> valid packet
				return ValidPacket
			}
			result = comparePackets(packetOne.Elements[idx], packetTwo.Elements[idx])
			if result == ValidPacket || result == InvalidPacket {
				return result
			}
		}

		if idx < len(packetOne.Elements)-1 {
			// Right-side ran out of items -> invalid packet
			return InvalidPacket
		}

		return result
	} else {
		// Do list vs number comparison
		if packetOne.IsList() {
			// Left-side ran out of items -> valid oacket
			if len(packetOne.Elements) == 0 {
				return ValidPacket
			}
			var newPacket *Packet = &Packet{make([]*Packet, 0), -1}
			newPacket.Elements = append(newPacket.Elements, packetTwo)
			return comparePackets(packetOne, newPacket)
		} else {
			// Right-side ran out of items -> invalid packet
			if len(packetTwo.Elements) == 0 {
				return InvalidPacket
			}
			var newPacket *Packet = &Packet{make([]*Packet, 0), -1}
			newPacket.Elements = append(newPacket.Elements, packetOne)
			return comparePackets(newPacket, packetTwo)
		}
	}
}

func Day13() {
	contents := utils.GetFileContents("day13/input")
	validIdxs := make([]int, 0)

	packets := make([]*Packet, 0)

	for i := 0; i < len(contents); i += 2 {
		if contents[i] == "" {
			i--
			continue
		}

		packetOne, packetTwo := parsePackets(contents[i], contents[i+1])

		packets = append(packets, packetOne)
		packets = append(packets, packetTwo)

		validity := comparePackets(packetOne, packetTwo)
		if validity == ValidPacket {
			validIdxs = append(validIdxs, i/3+1)
		}
	}

	// Sum indices of valid packets
	sum := 0
	for _, val := range validIdxs {
		sum += val
	}

	fmt.Println("Result part 1:", sum)

	// Part 2 - find indices of divider packets and multiple them
	// Initialize keys with 1 and 2 because key2 needs to include key1
	key1, key2 := 1, 2

	// Build packets for [[2]] and [[6]]
	dividerPackets := make([]*Packet, 0)
	dividerPackets = append(dividerPackets, &Packet{nil, 2})
	dividerPacket1 := &Packet{dividerPackets, -1}

	dividerPackets = make([]*Packet, 0)
	dividerPackets = append(dividerPackets, &Packet{nil, 6})
	dividerPacket2 := &Packet{dividerPackets, -1}

	// Iterate all packets and count all valid packets for each divider key
	for _, packet := range packets {
		result1 := comparePackets(packet, dividerPacket1)
		result2 := comparePackets(packet, dividerPacket2)

		if result1 == ValidPacket {
			key1 += 1
		}

		if result2 == ValidPacket {
			key2 += 1
		}
	}

	fmt.Println("Result part 2:", key1*key2)
}
