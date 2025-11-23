package question07

import (
	"reflect"
	"testing"
)

// TestSolution1 反转整数单元测试1
func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs int
		expect int
	}{
		{"TestCacse 1", 123, 321},
		{"TestCacse 2", -123, -321},
		{"TestCacse 3", 120, 21},
		{"TestCacse 4", 1534236469, 0},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverse1(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs int
		expect int
	}{
		{"TestCacse 1", 123, 321},
		{"TestCacse 2", -123, -321},
		{"TestCacse 3", 120, 21},
		{"TestCacse 4", 1534236469, 0},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := reverse2(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
			}
		})
	}
}
