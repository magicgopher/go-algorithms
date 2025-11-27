package question13

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"TestCase", "III", 3},
		{"TestCase", "IV", 4},
		{"TestCase", "IX", 9},
		{"TestCase", "LVIII", 58},
		{"TestCase", "MCMXCIV", 1994},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := romanToInt1(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"TestCase", "III", 3},
		{"TestCase", "IV", 4},
		{"TestCase", "IX", 9},
		{"TestCase", "LVIII", 58},
		{"TestCase", "MCMXCIV", 1994},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := romanToInt2(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestSolution3(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"TestCase", "III", 3},
		{"TestCase", "IV", 4},
		{"TestCase", "IX", 9},
		{"TestCase", "LVIII", 58},
		{"TestCase", "MCMXCIV", 1994},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := romanToInt3(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

func TestSolution4(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs string
		expect int
	}{
		{"TestCase", "III", 3},
		{"TestCase", "IV", 4},
		{"TestCase", "IX", 9},
		{"TestCase", "LVIII", 58},
		{"TestCase", "MCMXCIV", 1994},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := romanToInt4(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}
