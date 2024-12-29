package day24

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput1 = `x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02`

var exampleInput2 = `x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj`

func TestEvaluateExampleOne(t *testing.T) {
	result := Evaluate(exampleInput1)
	if result != 4 {
		t.Errorf("Expected 4 after evaluation found %d", result)
	}
}

func TestEvaluateExampleTwo(t *testing.T) {
	result := Evaluate(exampleInput2)
	if result != 2024 {
		t.Errorf("Expected 4 after evaluation found %d", result)
	}
}

func TestEvaluate(t *testing.T) {
	result := Evaluate(year2024.ReadInput("input.txt"))
	if result != 61495910098126 {
		t.Errorf("Expected 61495910098126 after evaluation found %d", result)
	}
}

func TestFixWires(t *testing.T) {
	result := FindBadWires(year2024.ReadInput("input.txt"))
	if result == "cwt,gdd,jmv,pqt,z05,z09,z37,z45" {
		t.Errorf("Expected input to not be cwt,gdd,jmv,pqt,z05,z09,z37,z45, but it was")
	}
	if result != "css,cwt,gdd,jmv,pqt,z05,z09,z37" {
		t.Errorf("Expected bad wires css,cwt,gdd,jmv,pqt,z05,z09,z37, found, %s", result)
	}
}
