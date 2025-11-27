package question14

import "testing"

func TestSolution1(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []string{
			"flower",
			"flow",
			"flight",
		}
		got := longestCommonPrefix1(data)
		want := "fl"
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := []string{
			"dog",
			"racecar",
			"car",
		}
		got := longestCommonPrefix1(data)
		want := ""
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}

func TestSolution2(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []string{
			"flower",
			"flow",
			"flight",
		}
		got := longestCommonPrefix2(data)
		want := "fl"
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := []string{
			"dog",
			"racecar",
			"car",
		}
		got := longestCommonPrefix2(data)
		want := ""
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}

func TestSolution3(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []string{
			"flower",
			"flow",
			"flight",
		}
		got := longestCommonPrefix3(data)
		want := "fl"
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := []string{
			"dog",
			"racecar",
			"car",
		}
		got := longestCommonPrefix3(data)
		want := ""
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
