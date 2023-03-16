package main

import (
	"fmt"
	"os"
)

func ray_color(r ray) color {
	if (hit_sphere(point3{0, 0, -1}, 0.5, r)) {
		return color{1, 0, 0}
	}
	unit_direction := UnitVector(r.direction())
	t := 0.5 * (unit_direction.Y() + 1.0)
	return Add(color{1.0, 1.0, 1.0}.Multi(1.0-t), color{0.5, 0.7, 1.0}.Multi(t))
}

func hit_sphere(center point3, radius float64, r ray) bool {
	oc := Sub(r.origin(), center)
	a := Dot(r.direction(), r.direction())
	b := 2.0 * Dot(oc, r.direction())
	c := Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}

func routine() {
	// Image
	const aspect_ratio float64 = 16.0 / 9.0
	const image_width int = 400
	const image_height int = int(float64(image_width) / aspect_ratio)

	// Camera
	viewport_height := 2.0
	viewport_width := aspect_ratio * viewport_height
	focal_length := 1.0

	origin := point3{0, 0, 0}
	horizontal := vec3{viewport_width, 0, 0}
	vertical := vec3{0, viewport_height, 0}
	// lower_left_corner = origin - horizontal/2 - vertical/2 - vec3(0, 0, focal_length);
	lower_left_corner := Sub(origin, horizontal.Div(2))
	lower_left_corner = Sub(lower_left_corner, vertical.Div(2))
	lower_left_corner = Sub(lower_left_corner, vec3{0, 0, focal_length})

	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)
	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", j)
		for i := 0; i < image_width; i++ {
			u := float64(i) / (float64(image_width - 1))
			v := float64(j) / (float64(image_height - 1))
			ray_dir := Add(lower_left_corner, horizontal.Multi(u))
			ray_dir = Add(ray_dir, vertical.Multi(v))
			ray_dir = Sub(ray_dir, origin)
			r := ray{origin, ray_dir}
			pixel_color := ray_color(r)
			write_color(pixel_color)
		}
	}
	fmt.Fprint(os.Stderr, "\nDone\n")

	//v := new(vec3)
	//fmt.Fprintln(os.Stderr, v.X())

}

func main() {
	routine()
}
