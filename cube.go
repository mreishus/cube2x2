// Package main provides ...
package main

import "fmt"

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

func main() {
   cube := [24]CubeColor{White, White, White, White, Orange, Orange, Orange, Orange, Green, Green, Green, Green, Red, Red, Red, Red, Blue, Blue, Blue, Blue, Yellow, Yellow, Yellow, Yellow}
   fmt.Println(cube[0])
}
