package day24

import (
	"fmt"
	"goated-aoc-2024/year2024"
	"sort"
	"strconv"
	"strings"
)

type Gate struct {
	LHS string
	RHS string
	Op  string
}

func Evaluate(input string) int64 {
	lines := strings.Split(input, "\n")
	initialValues := make(map[string]int)
	outputWires := make(map[string]Gate)
	zCount := 0
	for _, line := range lines {
		if strings.Contains(line, ": ") {
			split := strings.Split(line, ": ")
			num, _ := strconv.Atoi(split[1])
			initialValues[split[0]] = num
		} else if strings.Contains(line, " -> ") {
			split := strings.Split(line, " -> ")
			output := split[len(split)-1]
			if strings.HasPrefix(output, "z") {
				zCount++
			}
			equation := strings.Split(split[0], " ")
			gate := Gate{LHS: equation[0], RHS: equation[2], Op: equation[1]}
			outputWires[output] = gate
		}
	}
	zs := make([]string, zCount)
	for i := zCount - 1; i >= 0; i-- {
		suffix := fmt.Sprintf("%02d", i)
		zWire := "z" + suffix
		zs[(zCount-1)-i] = strconv.Itoa(evaluate(zWire, outputWires, initialValues))
	}
	num, _ := strconv.ParseInt(strings.Join(zs, ""), 2, 64)
	return num
}

type Candidate struct {
	Last    string
	Visited *year2024.HashSet[string]
}

func FindBadWires(input string) string {
	lines := strings.Split(input, "\n")
	outputWires := make(map[string]Gate)
	zCount := 0
	for _, line := range lines {
		if strings.Contains(line, " -> ") {
			split := strings.Split(line, " -> ")
			output := split[len(split)-1]
			if strings.HasPrefix(output, "z") {
				zCount++
			}
			equation := strings.Split(split[0], " ")
			gate := Gate{LHS: equation[0], RHS: equation[2], Op: equation[1]}
			outputWires[output] = gate
		}
	}
	badNodes := year2024.NewHashSet[string]()
	for i := 0; i < zCount; i++ {
		suffix := fmt.Sprintf("%02d", i)
		zWire := "z" + suffix
		gate := outputWires[zWire]
		wire, _ := validateSumWire(zWire, gate, outputWires, i)
		if wire != "" {
			badNodes.Add(wire)
		}
	}
	badNodesSlice := make([]string, badNodes.Size())
	i := 0
	for n := range badNodes.Iterator() {
		badNodesSlice[i] = n
		i++
	}
	sort.Strings(badNodesSlice)
	return strings.Join(badNodesSlice, ",")
}

func hasXorForGate(gate Gate, i int) bool {
	suffix := fmt.Sprintf("%02d", i)
	x := "x" + suffix
	y := "y" + suffix
	return ((gate.LHS == x && gate.RHS == y) || (gate.LHS == y && gate.RHS == x)) && gate.Op == "XOR"
}

func hasAndForGate(gate Gate, i int) bool {
	suffix := fmt.Sprintf("%02d", i)
	x := "x" + suffix
	y := "y" + suffix
	return ((gate.LHS == x && gate.RHS == y) || (gate.LHS == y && gate.RHS == x)) && gate.Op == "AND"
}

func evaluate(wire string, wires map[string]Gate, values map[string]int) int {
	gate := wires[wire]
	lhs, exists := values[gate.LHS]
	if !exists {
		lhs = evaluate(gate.LHS, wires, values)
	}
	rhs, rhsExists := values[gate.RHS]
	if !rhsExists {
		rhs = evaluate(gate.RHS, wires, values)
	}
	switch gate.Op {
	case "OR":
		return lhs | rhs
	case "AND":
		return lhs & rhs
	case "XOR":
		return lhs ^ rhs
	}
	panic("unknown op")
}

func validateSumWire(sumWire string, gate Gate, wires map[string]Gate, i int) (string, string) {
	if i == 0 {
		if !hasXorForGate(gate, i) {
			return sumWire, "Expected sum wire at level 0 to be an XOR of inputs x00 and y00"
		}
		return "", ""
	}
	if i == 45 {
		// should just be a carry wire, not a sum wire as it is the last z
		return validateCarryV2(sumWire, gate, wires, i)
	}
	lhsGate := wires[gate.LHS]
	rhsGate := wires[gate.RHS]
	if gate.Op != "XOR" {
		return sumWire, fmt.Sprintf("Expected sum wire to have an op XOR, instead found %s", gate.Op)
	}

	if i == 1 {
		expectedOps := year2024.NewHashSet[string]()
		expectedOps.Add("AND")
		expectedOps.Add("XOR")
		// half adder
		if !(expectedOps.Contains(rhsGate.Op)) {
			return gate.RHS, fmt.Sprintf("Expected input RHS wire to sum wire %s to have a gate that was an AND or XOR, instead found %s", sumWire, rhsGate.Op)
		}
		if !(expectedOps.Contains(lhsGate.Op)) {
			return gate.LHS, fmt.Sprintf("Expected input LHS wire to sum wire %s to have a gate that was an AND or XOR, instead found %s", sumWire, lhsGate.Op)
		}
		if lhsGate.Op == rhsGate.Op {
			panic(fmt.Sprintf("Need to find the correct sub wire to swap - LHS: %v, RHS: %v", lhsGate, rhsGate))
		}
		andGate := rhsGate
		xOrGate := lhsGate
		if lhsGate.Op == "AND" {
			andGate = lhsGate
			xOrGate = rhsGate
		}
		if !(hasAndForGate(andGate, i-1) && hasXorForGate(xOrGate, i)) {
			panic(fmt.Sprintf("Need to find the correct sub wire to swap - LHS: %v, RHS: %v", lhsGate, rhsGate))
		}
	} else {
		expectedOps := year2024.NewHashSet[string]()
		expectedOps.Add("OR")
		expectedOps.Add("XOR")
		if !(expectedOps.Contains(rhsGate.Op)) {
			return gate.RHS, fmt.Sprintf("Expected input RHS wire to sum wire %s to have a gate that was an OR or XOR, instead found %s", sumWire, rhsGate.Op)
		}
		if !(expectedOps.Contains(lhsGate.Op)) {
			return gate.LHS, fmt.Sprintf("Expected input LHS wire to sum wire %s to have a gate that was an OR or XOR, instead found %s", sumWire, lhsGate.Op)
		}
		if lhsGate.Op == rhsGate.Op {
			if strings.HasPrefix(lhsGate.LHS, "y") || strings.HasPrefix(lhsGate.RHS, "y") {
				return gate.RHS, fmt.Sprintf("RHS of input to sum wire has an XOR for elements not the inputs, %v", rhsGate)
			} else {
				return gate.LHS, fmt.Sprintf("LHS of input to sum wire has an XOR for elements not the inputs, %v", lhsGate)
			}
		}
		xOrGate := rhsGate
		xOrWire := gate.RHS
		orGate := lhsGate
		orWire := gate.LHS
		if lhsGate.Op == "XOR" {
			xOrGate = lhsGate
			xOrWire = gate.LHS
			orGate = rhsGate
			orWire = gate.RHS
		}
		if !(hasXorForGate(xOrGate, i)) {
			return xOrWire, fmt.Sprintf("Expected XOR wire %s to be an XOR of inputs, but was %v", xOrWire, xOrGate)
		}
		badWire, err := validateCarryV2(orWire, orGate, wires, i)
		if badWire != "" {
			return badWire, err
		}
	}
	return "", ""
}

func validateCarryV2(wire string, gate Gate, wires map[string]Gate, i int) (string, string) {
	rhsGate := wires[gate.RHS]
	lhsGate := wires[gate.LHS]
	if i == 1 {
		if gate.Op != "AND" {
			return wire, fmt.Sprintf("Expected carry for level 1 to be just an AND as it is a half adder, but found %s", gate.Op)
		}
		if !hasAndForGate(gate, i-1) {
			return wire, fmt.Sprintf("Expected carry for level 1 to be an AND of previous input, but found %v", gate)
		}
	} else {
		if gate.Op != "OR" {
			return wire, fmt.Sprintf("Expected carry for level %d to have an OR as an op, but found %s", i, gate.Op)
		}
		expectedOps := year2024.NewHashSet[string]()
		expectedOps.Add("AND")
		if !expectedOps.Contains(lhsGate.Op) {
			return gate.LHS, fmt.Sprintf("Expected LHS of input to carry for %d to have and AND, but found %s", i, lhsGate.Op)
		}
		if !expectedOps.Contains(rhsGate.Op) {
			return gate.RHS, fmt.Sprintf("Expected RHS of input to carry for %d to have and AND, but found %s", i, rhsGate.Op)
		}
		var andOfInputGate Gate
		var andOfCarryGate Gate
		if strings.HasPrefix(lhsGate.LHS, "y") || strings.HasPrefix(lhsGate.RHS, "y") {
			andOfCarryGate = rhsGate
			andOfInputGate = lhsGate
		} else {
			andOfCarryGate = lhsGate
			andOfInputGate = rhsGate
		}
		if !hasAndForGate(andOfInputGate, i-1) {
			panic("need to find bad AND gate")
		}
		lhsOfCarryAndGate := wires[andOfCarryGate.LHS]
		rhsOfCarryAndGate := wires[andOfCarryGate.RHS]
		expectedSubOps := year2024.NewHashSet[string]()
		expectedSubOps.Add("OR")
		expectedSubOps.Add("XOR")
		if i == 2 {
			expectedSubOps.Remove("OR")
			expectedSubOps.Add("AND")
		}
		if !expectedSubOps.Contains(lhsOfCarryAndGate.Op) {
			return andOfCarryGate.LHS, fmt.Sprintf("Input to the AND of the carry gate for level %d that is not the AND of the input should have XOR and OR, instead found %s", i, lhsOfCarryAndGate.Op)
		}
		if !expectedSubOps.Contains(rhsOfCarryAndGate.Op) {
			return andOfCarryGate.RHS, fmt.Sprintf("Input to the AND of the carry gate for level %d that is not the AND of the input should have XOR and OR, instead found %s", i, rhsOfCarryAndGate.Op)
		}
		var xOrOfPreviousLevelGate Gate
		var xOrOfPreviousLevelWire string
		var carryOfPreviousLevelGate Gate
		var carryOfPreviousLevelWire string
		if lhsOfCarryAndGate.Op == "XOR" {
			xOrOfPreviousLevelGate = lhsOfCarryAndGate
			xOrOfPreviousLevelWire = andOfCarryGate.LHS
			carryOfPreviousLevelGate = rhsOfCarryAndGate
			carryOfPreviousLevelWire = andOfCarryGate.RHS
		} else {
			xOrOfPreviousLevelGate = rhsOfCarryAndGate
			xOrOfPreviousLevelWire = andOfCarryGate.RHS
			carryOfPreviousLevelGate = lhsOfCarryAndGate
			carryOfPreviousLevelWire = andOfCarryGate.LHS
		}
		if !hasXorForGate(xOrOfPreviousLevelGate, i-1) {
			return xOrOfPreviousLevelWire, fmt.Sprintf("The XOR gate for the AND of the carry gate that is not the AND of the previous level shoudl be an XOR of the previous level, instead found %v", xOrOfPreviousLevelGate)
		}
		return validateCarryV2(carryOfPreviousLevelWire, carryOfPreviousLevelGate, wires, i-1)
	}
	return "", ""
}
