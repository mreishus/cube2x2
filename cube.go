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

func doTurn(cube [24]CubeColor, turn CubeTurn) [24]CubeColor {
	var turns [][4]int

	switch turn {
	case U:
		turns = [][4]int{
			{8, 12, 16, 4}, // I M Q E
			{9, 13, 17, 5}, // J N R F
			{0, 3, 2, 1},   // A D C B
		}
	case UP:
		turns = [][4]int{
			{8, 4, 16, 12}, // I E Q M
			{9, 5, 17, 13}, // J F R N
			{0, 1, 2, 3},   // A B C D
		}
	case U2:
		// 4 style turns don't work, I need to do direct swaps...
		// Swap I,Q
		// Swap J,R
		// Swap E,M
		// Swap F,N
		// Also rotate the U face
	case R:
		turns = [][4]int{
			{9, 21, 19, 1},   // J V T B
			{10, 22, 16, 2},  // K W Q C
			{12, 15, 14, 13}, // M P O N
		}
	case RP:
		turns = [][4]int{
			{9, 1, 19, 21},   // J B T V
			{10, 2, 16, 22},  // K C Q W
			{12, 13, 14, 15}, // M N O P
		}
	case R2:
		// 4 style turns don't work, I need to do direct swaps...
		// Also rotate the R face
	case F:
	case FP:
	case F2:
		// 4 style turns don't work, I need to do direct swaps...
		// Also rotate the F face
	}

	for _, turn := range turns {
		cube = doRotate(cube, turn)
	}
	return cube
}

func doRotate(cube [24]CubeColor, is [4]int) [24]CubeColor {
	var temp CubeColor
	a, b, c, d := is[0], is[1], is[2], is[3]
	temp = cube[a]
	cube[a] = cube[b]
	cube[b] = cube[c]
	cube[c] = cube[d]
	cube[d] = temp
	return cube
}

func display(cube [24]CubeColor) {
	// Print U
	fmt.Print("    ")
	displaySquares(cube, []int{0, 1})
	fmt.Println()
	fmt.Print("    ")
	displaySquares(cube, []int{3, 2})
	fmt.Println()
	// Print L, F, R, B
	displaySquares(cube, []int{4, 5, 8, 9, 12, 13, 16, 17})
	fmt.Println()
	displaySquares(cube, []int{7, 6, 11, 10, 15, 14, 19, 18})
	fmt.Println()
	// Print D
	fmt.Print("    ")
	displaySquares(cube, []int{20, 21})
	fmt.Println()
	fmt.Print("    ")
	displaySquares(cube, []int{23, 22})
	fmt.Println()
	fmt.Println()
}

func displaySquares(cube [24]CubeColor, indexes []int) {
	for _, x := range indexes {
		displaySquare(cube[x])
		displaySquare(cube[x])
	}
}

func displaySquare(cc CubeColor) {
	block := "█"
	switch cc {
	case White:
		fmt.Print(gchalk.White(block))
	case Orange:
		fmt.Print(gchalk.Yellow(block))
	case Green:
		fmt.Print(gchalk.Green(block))
	case Red:
		fmt.Print(gchalk.Red(block))
	case Blue:
		fmt.Print(gchalk.Blue(block))
	case Yellow:
		fmt.Print(gchalk.BrightYellow(block))
	}
}

func GetSolvedCube() [24]CubeColor {
	cube := [24]CubeColor{White, White, White, White, Orange, Orange, Orange, Orange, Green, Green, Green, Green, Red, Red, Red, Red, Blue, Blue, Blue, Blue, Yellow, Yellow, Yellow, Yellow}
	return cube
}

func SMove(cube [24]CubeColor) [24]CubeColor {
	cube = doTurn(cube, R)
	cube = doTurn(cube, U)
	cube = doTurn(cube, RP)
	cube = doTurn(cube, UP)
	// display(cube)
	return cube
}

func main() {
	cube := GetSolvedCube()
	display(cube)
	cube = SMove(cube)
	display(cube)
	cube = SMove(cube)
	display(cube)
	cube = SMove(cube)
	display(cube)
	cube = SMove(cube)
	display(cube)
	cube = SMove(cube)
	display(cube)
	cube = SMove(cube)
	display(cube)
}
