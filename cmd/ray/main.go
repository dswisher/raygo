// Ray tracer.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const nx = 200
	const ny = 100

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "You must specify the name of the output file!\n")
		os.Exit(1)
	}

	fname := os.Args[1]

	f, err := os.Create(fname)
	check(err)

	defer f.Close()

	for j := 0; j < ny; j++ {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(ny-j) / float64(ny)
			b := 0.2

			ir := uint8(255.99 * r)
			ig := uint8(255.99 * g)
			ib := uint8(255.99 * b)

			c := color.RGBA{ir, ig, ib, 255}

			img.Set(i, j, c)
		}
	}

	png.Encode(f, img)
}
