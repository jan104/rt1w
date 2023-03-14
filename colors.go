package main

import "fmt"

func write_color(pixel_color color) {
	fmt.Printf("%d %d %d\n", int32(255.999*pixel_color.X()), int32(255.999*pixel_color.Y()), int32(255.999*pixel_color.Z()))
}
