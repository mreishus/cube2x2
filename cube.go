// Package main provides ...
package main

import "fmt"
import "github.com/jwalton/gchalk"

func AddOne(x int) int {
	return x + 1
}

type CubeColor int8

const (
	White CubeColor = iota
	Orange
	Green
	Red
	Blue
	Yellow
)

type CubeTurn int8

const (
	F CubeTurn = iota
	FP
	F2
	R
	RP
	R2
	U
	UP
	U2
)

// Cube in different formats:
// {W, W, W, W, O, O, O, O, G, G, G,   G,  R,  R,  R,  R,  B,  B,  B,  B,  Y,  Y,  Y,  Y} // color
// {U, U, U, U, L, L, L, L, F, F, F,   F,  R,  R,  R,  R,  B,  B,  B,  B,  D,  D,  D,  D} // face
// {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23} // index
// {A, B, C, D, E, F, G, H, I, J,  K,  L,  M,  N,  O,  P,  Q,  R,  S,  T,  U,  V,  W,  X} // blindfold

func do_turn(cube [24]CubeColor, turn CubeTurn) [24]CubeColor {
	var temp1 CubeColor
	var temp2 CubeColor

	switch turn {
	case U:
		temp1, temp2 = cube[8], cube[9]
		cube[8], cube[9] = cube[12], cube[13]
		cube[12], cube[13] = cube[16], cube[17]
		cube[16], cube[17] = cube[4], cube[5]
		cube[4], cube[5] = temp1, temp2
	case UP:
		temp1, temp2 = cube[8], cube[9]
		cube[8], cube[9] = cube[4], cube[5]
		cube[4], cube[5] = cube[16], cube[17]
		cube[16], cube[17] = cube[12], cube[13]
		cube[12], cube[13] = temp1, temp2

	}
	return cube
}

func display(cube [24]CubeColor) {
	// Print U
	fmt.Print("  ")
	display_squares(cube, []int{0, 1})
	fmt.Println()
	fmt.Print("  ")
	display_squares(cube, []int{3, 2})
	fmt.Println()
	// Print L, F, R, B
	display_squares(cube, []int{4, 5, 8, 9, 12, 13, 16, 17})
	fmt.Println()
	display_squares(cube, []int{7, 6, 11, 10, 15, 14, 19, 18})
	fmt.Println()
	// Print D
	fmt.Print("  ")
	display_squares(cube, []int{20, 21})
	fmt.Println()
	fmt.Print("  ")
	display_squares(cube, []int{23, 22})
	fmt.Println()
	fmt.Println()
}

func display_squares(cube [24]CubeColor, indexes []int) {
	for _, x := range indexes {
		display_square(cube[x])
	}
}

func display_square(cc CubeColor) {
	block := "█"
	switch cc {
	case White:
		fmt.Print(gchalk.White(block))
	case Orange:
		fmt.Print(gchalk.BrightYellow(block))
	case Green:
		fmt.Print(gchalk.Green(block))
	case Red:
		fmt.Print(gchalk.Red(block))
	case Blue:
		fmt.Print(gchalk.Blue(block))
	case Yellow:
		fmt.Print(gchalk.Yellow(block))
	}
}

func get_solved_cube() [24]CubeColor {
	cube := [24]CubeColor{White, White, White, White, Orange, Orange, Orange, Orange, Green, Green, Green, Green, Red, Red, Red, Red, Blue, Blue, Blue, Blue, Yellow, Yellow, Yellow, Yellow}
	return cube
}


func main() {
	cube := get_solved_cube()
	display(cube)
	cube = do_turn(cube, U)
	display(cube)
	cube = do_turn(cube, UP)
	display(cube)
	cube = do_turn(cube, UP)
	display(cube)
}
