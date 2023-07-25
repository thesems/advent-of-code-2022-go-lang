package day13

import (
	"adventofcode/utils"
	"fmt"
	"log"
)

type Expression struct {
	nextValue *Expression
	value     int
}

func (e *Expression) IsList() bool {
	if e.nextValue == nil {
		return false
	}
	return true
}

func (e *Expression) Value() int {
	return e.value
}

func (e *Expression) ToList() {
	if e.IsList() {
		log.Fatal("Expression already a list.")
	}

	e.nextValue = &Expression{nil, e.value}
	e.value = 0
}

func convertToPacket(packet string) *Expression {
	// Parse list or number

	return &Expression{nil, 1}
}

func parsePackets(packetOneStr string, packetTwoStr string) (*Expression, *Expression) {
	// convert to Expression
	packetOne := convertToPacket(packetOneStr)
	packetTwo := convertToPacket(packetTwoStr)
	return packetOne, packetTwo
}

const (
	InvalidPacket   = 0
	ValidPacket     = 1
	UndecidedPacket = 2
)

func comparePackets(packetOne *Expression, packetTwo *Expression) int {
	for {
		// Perform comparison logic
		if !packetOne.IsList() && !packetTwo.IsList() {
			// Do number vs number comparison
			diff := packetOne.Value() - packetTwo.Value()

			if diff < 0 {
				return ValidPacket
			} else if diff == 0 {
				// TODO: what happens here?
				return UndecidedPacket
			}
			return InvalidPacket
		} else if packetOne.IsList() && packetTwo.IsList() {
			// Do list vs list comparison
			return comparePackets(packetOne.nextValue, packetTwo.nextValue)
		} else {
			// Do list vs number comparison
			if packetOne.IsList() {
				return comparePackets(packetOne.nextValue, packetTwo)
			} else {
				return comparePackets(packetOne, packetTwo.nextValue)
			}
		}
	}
}

func Day13() {
	contents := utils.GetFileContents("day13/example")
	validIdxs := make([]int, 0)
	for i := 0; i < len(contents); i += 2 {
		if contents[i] == "" {
			i--
			continue
		}

		packetOne, packetTwo := parsePackets(contents[i], contents[i+1])
		validity := comparePackets(packetOne, packetTwo)

		if validity == ValidPacket {
			validIdxs = append(validIdxs, i/2+1)
		}
	}

	sum := 0
	for _, val := range validIdxs {
		sum += val
	}

	fmt.Println("Result:", sum)
}
