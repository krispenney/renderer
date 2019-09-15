package main

import (
	"./vector"
	"fmt"
)

func main() {
	x1 := vector.Vector{1, 0.75, 0}
	x2 := vector{-1, 0.75, 0}
	x3 := vector{1, -0.75, 0}
	x4 := vector{-1, -0.75, 0}

	camera := vector{0, 0, -1}

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(camera)

	fmt.Println(x1.add(x2))
}
