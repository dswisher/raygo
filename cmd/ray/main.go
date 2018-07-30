// Ray tracer.

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/dswisher/raygo/pkg/raygo"
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

	lowerLeft := raygo.Vector{-2.0, -1.0, -1.0}
	horizontal := raygo.Vector{4.0, 0.0, 0.0}
	vertical := raygo.Vector{0.0, 2.0, 0.0}
	origin := raygo.Vector{0.0, 0.0, 0.0}

	for j := 0; j < ny; j++ {
		for i := 0; i < nx; i++ {

			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			r := raygo.Ray{origin, lowerLeft.Add(horizontal.Mult(u)).Add(vertical.Mult(v))}

			cv := r.Color()

			ir := uint8(255.99 * cv.X)
			ig := uint8(255.99 * cv.Y)
			ib := uint8(255.99 * cv.Z)

			c := color.RGBA{ir, ig, ib, 255}

			img.Set(i, j, c)
		}
	}

	png.Encode(f, img)
}
