package main

import (
	"fmt"
	"os"
)

func main() {
	routine()
}

func routine() {
	const image_width = 256
	const image_height = 256

	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)
	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", j)
		for i := 0; i < image_width; i++ {
			pixel_color := color{float64(i) / (image_width - 1), float64(j) / (image_height - 1), 0.25}
			write_color(pixel_color)
		}
	}
	fmt.Fprint(os.Stderr, "\nDone\n")

	//v := new(vec3)
	//fmt.Fprintln(os.Stderr, v.X())

}
