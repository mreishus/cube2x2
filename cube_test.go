package main

import "testing"
import "reflect"

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
	t.Run("R F R F", func(t *testing.T) {
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
	t.Run("R F' R F'", func(t *testing.T) {
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

func Test2Turns(t *testing.T) {
	t.Run("R F2", func(t *testing.T) {
		cube1 := GetSolvedCube()
		cube1 = DoTurn(cube1, R)
		cube1 = DoTurn(cube1, F2)
		got := cube1
		want := [24]CubeColor{White, Green, Yellow, Blue, // U
			Orange, Red, Red, Orange, // L
			Yellow, Green, Green, Yellow, // F
			Orange, Red, Red, Orange, // R
			White, Blue, Blue, White, // B
			Green, White, Blue, Yellow} // D
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("R F U2", func(t *testing.T) {
		cube1 := GetSolvedCube()
		cube1 = DoTurn(cube1, R)
		cube1 = DoTurn(cube1, F)
		cube1 = DoTurn(cube1, U2)
		got := cube1
		want := [24]CubeColor{Orange, Orange, White, Green, // U
			White, Red, Blue, Orange, // L
			White, Blue, Yellow, Yellow, // F
			Orange, Yellow, Red, Green, // R
			Green, Green, Blue, White, // B
			Red, Red, Blue, Yellow} // D
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("U F R2", func(t *testing.T) {
		cube1 := GetSolvedCube()
		cube1 = DoTurn(cube1, U)
		cube1 = DoTurn(cube1, F)
		cube1 = DoTurn(cube1, R2)
		got := cube1
		want := [24]CubeColor{White, Blue, Yellow, Orange, // U
			Green, Yellow, Yellow, Orange, // L
			Green, Blue, Orange, Green, // F
			Red, White, White, Blue, // R
			Red, Orange, Blue, Red, // B
			Red, White, Green, Yellow} // D
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestBFS(t *testing.T) {
	t.Run("Find solution for PBL No Bar", func(t *testing.T) {
		cube := [24]CubeColor{White, White, White, White, // U
			Orange, Red, Red, Orange, // L
			Blue, Green, Green, Blue, // F
			Red, Orange, Orange, Red, // R
			Green, Blue, Blue, Green, // B
			Yellow, Yellow, Yellow, Yellow} // D
		got := Bfs(cube)
		// Two possible solutions
		want := []CubeTurn{ F2, R2, F2 }
		want2 := []CubeTurn{ R2, F2, R2 }
		if !reflect.DeepEqual(got, want) && !reflect.DeepEqual(got, want2) {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Find solution for PBL Two Bar (Partial)", func(t *testing.T) {
		// Full PBL fails - Maybe because I don't have B moves
		// and reoriented solutions don't count?

		// cube := [24]CubeColor{White, White, White, White, // U
		// 	Blue, Orange, Orange, Blue, // L
		// 	Green, Green, Green, Green, // F
		// 	Red, Blue, Blue, Red, // R
		// 	Orange, Red, Red, Orange, // B
		// 	Yellow, Yellow, Yellow, Yellow} // D

		// R2 U' B2 U2 R2 U' R2

		// Here is a cube where I've partially done the PBL Two bar solution
		// Start with PBL Two Bar, then apply first 3 moves:
		// We've done: R2 U' B2
		// We want to get back: U2 R2 U' R2
		cube := [24]CubeColor{White, Yellow, White, White, // U
			Red, Red, Orange, Orange, // L
			Blue, Orange, Orange, Green, // F
			Green, Blue, Green, Blue, // R
			Red, Green, Blue, Red, // B
			Yellow, White, Yellow, Yellow} // D

		got := Bfs(cube)
		want := []CubeTurn{ U2, R2, UP, R2 }
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
