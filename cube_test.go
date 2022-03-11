package main

import "testing"

func TestSMoves(t *testing.T) {
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
			Red, White, Blue, Blue, // B
			Yellow, Red, Yellow, Yellow} // D
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestFTurns(t *testing.T) {
	t.Run("RFRF", func(t *testing.T) {
		cube1 := GetSolvedCube()
		cube1 = DoTurn(cube1, R)
		cube1 = DoTurn(cube1, F)
		cube1 = DoTurn(cube1, R)
		cube1 = DoTurn(cube1, F)
		got := cube1
		want := [24]CubeColor{White, Green, Yellow, Blue, // U
			Orange, Red, White, Orange, // L
			Yellow, Green, Red, Blue, // F
			Orange, White, Red, Yellow, // R
			Orange, Blue, Blue, Green, // B
			Red, Green, White, Yellow} // D
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("RF'RF'", func(t *testing.T) {
		cube1 := GetSolvedCube()
		cube1 = DoTurn(cube1, R)
		cube1 = DoTurn(cube1, FP)
		cube1 = DoTurn(cube1, R)
		cube1 = DoTurn(cube1, FP)
		got := cube1
		want := [24]CubeColor{White, Yellow, Red, Yellow, // U
			Orange, Green, Red, Orange, // L
			Orange, Blue, Green, Yellow, // F
			White, Blue, Red, Orange, // R
			Red, Blue, Blue, Green, // B
			Green, White, White, Yellow} // D
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
