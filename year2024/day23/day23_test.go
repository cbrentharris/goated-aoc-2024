package day23

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

func TestNumTConnectedExample(t *testing.T) {
	result := NumTConnected(exampleInput)
	if result != 7 {
		t.Errorf("expected 7 t connected 3 sets found %d", result)
	}
}

func TestNumTConnected(t *testing.T) {
	result := NumTConnected(year2024.ReadInput("input.txt"))
	if result != 1077 {
		t.Errorf("expected 7 t connected 3 sets found %d", result)
	}
}

func TestNumTConnectedV2Example(t *testing.T) {
	result := NumTConnectedV2(exampleInput)
	if result != "co,de,ka,ta" {
		t.Errorf("expected co,de,ka,ta max connection instead found %s", result)
	}
}
func TestNumTConnectedV2(t *testing.T) {
	result := NumTConnectedV2(year2024.ReadInput("input.txt"))
	if result != "bc,bf,do,dw,dx,ll,ol,qd,sc,ua,xc,yu,zt" {
		t.Errorf("expected bc,bf,do,dw,dx,ll,ol,qd,sc,ua,xc,yu,zt max connection instead found %s", result)
	}
}
