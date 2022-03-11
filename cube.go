// Package main provides ...
package main

import "fmt"
import "github.com/jwalton/gchalk"

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

type searchState struct {
	cube [24]CubeColor
	path []CubeTurn
}

// Ignore L, D, and B based turns, since it's a 2x2
// Also ignore cube rotations

// Cube in different formats:
// {W, W, W, W, O, O, O, O, G, G, G,   G,  R,  R,  R,  R,  B,  B,  B,  B,  Y,  Y,  Y,  Y} // color
// {U, U, U, U, L, L, L, L, F, F, F,   F,  R,  R,  R,  R,  B,  B,  B,  B,  D,  D,  D,  D} // face
// {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23} // index
// {A, B, C, D, E, F, G, H, I, J,  K,  L,  M,  N,  O,  P,  Q,  R,  S,  T,  U,  V,  W,  X} // blindfold

// Given a Cube, and a Turn, apply the turn to the cube and return a new cube
func DoTurn(cube [24]CubeColor, turn CubeTurn) [24]CubeColor {
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
		// Hack, do U turns twice
		turns = [][4]int{
			{8, 12, 16, 4}, // I M Q E
			{9, 13, 17, 5}, // J N R F
			{0, 3, 2, 1},   // A D C B
			{8, 12, 16, 4}, // I M Q E
			{9, 13, 17, 5}, // J N R F
			{0, 3, 2, 1},   // A D C B
		}
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
		// Hack, do R turns twice
		turns = [][4]int{
			{9, 21, 19, 1},   // J V T B
			{10, 22, 16, 2},  // K W Q C
			{12, 15, 14, 13}, // M P O N
			{9, 21, 19, 1},   // J V T B
			{10, 22, 16, 2},  // K W Q C
			{12, 15, 14, 13}, // M P O N
		}
	case F:
		turns = [][4]int{
			{3, 6, 21, 12}, // D G V M
			{2, 5, 20, 15}, // C F U P
			{8, 11, 10, 9}, // I L K J
		}
	case FP:
		turns = [][4]int{
			{3, 12, 21, 6}, // D M V G
			{2, 15, 20, 5}, // C P U F
			{8, 9, 10, 11}, // I J K L
		}
	case F2:
		// 4 style turns don't work, I need to do direct swaps...
		// Hack, do F turns twice :)
		turns = [][4]int{
			{3, 6, 21, 12}, // D G V M
			{2, 5, 20, 15}, // C F U P
			{8, 11, 10, 9}, // I L K J
			{3, 6, 21, 12}, // D G V M
			{2, 5, 20, 15}, // C F U P
			{8, 11, 10, 9}, // I L K J
		}
	}

	for _, turn := range turns {
		cube = doRotate(cube, turn)
	}
	return cube
}

// Given a cube, and a list of 4 indexes, do an "array rotation" of the
// colors belonging to those indexes, and return the new cube
func doRotate(cube [24]CubeColor, indexes [4]int) [24]CubeColor {
	var temp CubeColor
	a, b, c, d := indexes[0], indexes[1], indexes[2], indexes[3]
	temp = cube[a]
	cube[a] = cube[b]
	cube[b] = cube[c]
	cube[c] = cube[d]
	cube[d] = temp
	return cube
}

func PrintMoves(turns []CubeTurn) {
	for _, turn := range turns {
		name := ""
		switch turn {
		case F:
			name = "F "
		case FP:
			name = "F' "
		case F2:
			name = "F2 "
		case R:
			name = "R "
		case RP:
			name = "R' "
		case R2:
			name = "R2 "
		case U:
			name = "U "
		case UP:
			name = "U' "
		case U2:
			name = "U2 "
		}
		fmt.Print(name)
	}
	fmt.Println("")
}

// Given a cube, print it to screen
func Display(cube [24]CubeColor) {
	// Print U
	fmt.Print("    ")
	DisplaySquares(cube, []int{0, 1})
	fmt.Println()
	fmt.Print("    ")
	DisplaySquares(cube, []int{3, 2})
	fmt.Println()
	// Print L, F, R, B
	DisplaySquares(cube, []int{4, 5, 8, 9, 12, 13, 16, 17})
	fmt.Println()
	DisplaySquares(cube, []int{7, 6, 11, 10, 15, 14, 19, 18})
	fmt.Println()
	// Print D
	fmt.Print("    ")
	DisplaySquares(cube, []int{20, 21})
	fmt.Println()
	fmt.Print("    ")
	DisplaySquares(cube, []int{23, 22})
	fmt.Println()
	fmt.Println()
}

// Given a cube and a list of indexes, print the colors belonging to those
// indexes to the screen
func DisplaySquares(cube [24]CubeColor, indexes []int) {
	for _, x := range indexes {
		DisplaySquare(cube[x])
		DisplaySquare(cube[x])
	}
}

// Given a cube color, print it to screen
func DisplaySquare(cc CubeColor) {
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

// Return a fully solved cube
func GetSolvedCube() [24]CubeColor {
	cube := [24]CubeColor{White, White, White, White, Orange, Orange, Orange, Orange, Green, Green, Green, Green, Red, Red, Red, Red, Blue, Blue, Blue, Blue, Yellow, Yellow, Yellow, Yellow}
	return cube
}

// Given a cube, return a cube after an "SMove" has been applied, or R U R' U'
func SMove(cube [24]CubeColor) [24]CubeColor {
	cube = DoTurn(cube, R)
	cube = DoTurn(cube, U)
	cube = DoTurn(cube, RP)
	cube = DoTurn(cube, UP)
	return cube
}

// Given a cube, return a map of all possible next states
// key = CubeTurn
// value = New Cube
func nextStates(cube [24]CubeColor) map[CubeTurn][24]CubeColor {
	turns := []CubeTurn{F, FP, F2, R, RP, R2, U, UP, U2}
	m := make(map[CubeTurn][24]CubeColor)
	for _, turn := range turns {
		m[turn] = DoTurn(cube, turn)
	}
	return m
}

// Bfs. Given a cube, return a list of turns needed to make that
// cube solved
func Bfs(cube [24]CubeColor) []CubeTurn {
	var state searchState

	q := make([]searchState, 1)
	q[0] = searchState{cube: cube, path: []CubeTurn{}}

	maxDepth := 0

	for len(q) > 0 {

		if len(q) > maxDepth {
			maxDepth = len(q)
		}

		state, q = q[0], q[1:]

		if state.cube == GetSolvedCube() {
			fmt.Printf("bfs: maxDepth was %d\n", maxDepth)
			return state.path
		}

		for turn, newCube := range nextStates(state.cube) {
			newPath := make([]CubeTurn, len(state.path), len(state.path)+1)
			copy(newPath, state.path)
			newPath = append(newPath, turn)

			newState := searchState{cube: newCube, path: newPath}
			q = append(q, newState)
		}
		// Display(state.cube)
	}
	return []CubeTurn{}
}

func main() {
	cube := GetSolvedCube()
	Display(cube)
	// for i := 0; i < 6; i += 1 {
	// 	cube = SMove(cube)
	// 	Display(cube)
	// }
	cube = DoTurn(cube, R)
	Display(cube)
	cube = DoTurn(cube, F2)
	Display(cube)

	cube = [24]CubeColor{White, White, White, White, // U
		Orange, Red, Red, Orange, // L
		Blue, Green, Green, Blue, // F
		Red, Orange, Orange, Red, // R
		Green, Blue, Blue, Green, // B
		Yellow, Yellow, Yellow, Yellow} // D
	Display(cube)
	Bfs(cube)

	cube = [24]CubeColor{White, White, White, White, // U
		Orange, Red, Red, Orange, // L
		Blue, Green, Green, Blue, // F
		Red, Orange, Orange, Red, // R
		Green, Blue, Blue, Green, // B
		Yellow, Yellow, Yellow, Yellow} // D
	Display(cube)
	Bfs(cube)
	// cube = DoTurn(cube, R2)
	// cube = DoTurn(cube, F2)
	// cube = DoTurn(cube, R2)
	// Display(cube)
}
