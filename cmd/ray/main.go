// Ray tracer.

package main

import (
	"fmt"
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

	f, err := os.Create("/mnt/d/WSL/ray-out/sample.ppm") // TODO: read from command-line
	check(err)

	defer f.Close()

	fmt.Fprintf(f, "P3\n%d %d\n255\n", nx, ny)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2

			ir := uint8(255.99 * r)
			ig := uint8(255.99 * g)
			ib := uint8(255.99 * b)

			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}
}
