package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

const WIDTH = 255
const HEIGHT = 255

type Ray struct {
	point, direction Vector
}

// https://stats.stackexchange.com/questions/281162/scale-a-number-between-a-range
func fitToRange(v float64, min float64, max float64, span float64) float64 {
	return (((v - min) / (max - min)) * span) + min
}

func render(rays *[WIDTH * HEIGHT]Ray, min_x, max_x, min_y, max_y float64) {
	out, err := os.Create("./output.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

	for i, ray := range *rays {
		x := i % HEIGHT
		y := int(math.Floor(float64(i) / HEIGHT))
		direction := ray.direction

		red := fitToRange(direction.x, min_x, max_x, WIDTH)
		green := fitToRange(direction.y, min_y, max_y, HEIGHT)

		img.Set(x, y, color.RGBA{uint8(red), uint8(green), 100, 255})
	}

	var opt jpeg.Options
	opt.Quality = 100

	err = jpeg.Encode(out, img, &opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generated image to output.jpg")
}

func main() {
	x1 := Vector{1, 0.75, 0}
	x2 := Vector{-1, 0.75, 0}
	x3 := Vector{1, -0.75, 0}
	x4 := Vector{-1, -0.75, 0}

	camera := Vector{0, 0, -1}

	rays := [WIDTH * HEIGHT]Ray{}

	var min_x, max_x float64 = math.MaxFloat64, math.SmallestNonzeroFloat64
	var min_y, max_y float64 = math.MaxFloat64, math.SmallestNonzeroFloat64

	for y := 0; y < HEIGHT; y++ {
		beta := float64(y) / HEIGHT
		neg_beta := 1 - beta

		for x := 0; x < WIDTH; x++ {
			alpha := float64(x) / WIDTH
			neg_alpha := 1 - alpha

			x1_scaled := x1.scale(neg_alpha)
			x3_scaled := x3.scale(neg_alpha)

			t := x1_scaled.add(x2.scale(alpha))
			t_scaled := t.scale(neg_beta)

			b := x3_scaled.add(x4.scale(alpha))
			b_scaled := b.scale(beta)

			p := t_scaled.add(b_scaled)

			ray := Ray{p, p.subtract(camera)}
			rays[y*HEIGHT+x] = ray
			direction := ray.direction

			if direction.x < min_x {
				min_x = direction.x
			} else if max_x < direction.x {
				max_x = direction.x
			}

			if direction.y < min_y {
				min_y = direction.y
			} else if max_y < direction.y {
				max_y = direction.y
			}
		}
	}

  render(&rays, min_x, max_x, min_y, max_y)
}
