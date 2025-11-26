package question09

import "testing"

// TestSolution 单元测试
func TestSolution(t *testing.T) {
	// Test-1
	t.Run("Test-1", func(t *testing.T) {
		got := isPalindrome1(121121)
		want := true
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	// Test-2
	t.Run("Test-2", func(t *testing.T) {
		got := isPalindrome1(-121)
		want := false
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	// Test-3
	t.Run("Test-3", func(t *testing.T) {
		got := isPalindrome1(10)
		want := false
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	// Test-4
	t.Run("Test-1", func(t *testing.T) {
		got := isPalindrome2(121121)
		want := true
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	// Test-5
	t.Run("Test-2", func(t *testing.T) {
		got := isPalindrome2(-121)
		want := false
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	// Test-6
	t.Run("Test-3", func(t *testing.T) {
		got := isPalindrome2(10)
		want := false
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
