package main

import "testing"

func TestSixSMoves(t *testing.T) {
	t.Run("6 SMoves Restores the Cube", func(t *testing.T) {
		cube1 := GetSolvedCube()
		for i := 0; i < 6; i += 1 {
			cube1 = SMove(cube1)
		}
		got := cube1
		want := GetSolvedCube()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestTwoSMoves(t *testing.T) {
	t.Run("2 SMoves", func(t *testing.T) {
		cube1 := GetSolvedCube()
		for i := 0; i < 2; i += 1 {
			cube1 = SMove(cube1)
		}
		got := cube1
		want := [24]CubeColor{Orange, Blue, Green, White, // U
			Blue, Orange, Orange, Orange, // L
			Green, Red, Yellow, Green, // F
			White, White, Red, Green, // R
			Red, White, Blue, Blue,  // B
			Yellow, Red, Yellow, Yellow} // D
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
