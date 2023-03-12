package main

import (
	"fmt"
	"os"
)

func main() {
	const image_width = 256
	const image_height = 256

	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)
	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", j)
		for i := 0; i < image_width; i++ {
			r := float32(i) / (image_width - 1)
			g := float32(j) / (image_height - 1)
			b := 0.25

			ir := int32(255.999 * r)
			ig := int32(255.999 * g)
			ib := int32(255.999 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
	fmt.Fprint(os.Stderr, "\nDone\n")

	v := new(vec3)
	fmt.Fprintln(os.Stderr, v.X())
}
