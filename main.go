package main

import (
	"fmt"
	"math"
	"os"
)

func ray_color(r ray) color {
	t := hit_sphere(point3{0, 0, -1}, 0.5, r)
	if t > 0.0 {
		N := UnitVector(Sub(r.at(t), vec3{0, 0, -1}))
		return color{N.X() + 1, N.Y() + 1, N.Z() + 1}.Multi(0.5)
	}
	unit_direction := UnitVector(r.direction())
	t = 0.5 * (unit_direction.Y() + 1.0)
	return Add(color{1.0, 1.0, 1.0}.Multi(1.0-t), color{0.5, 0.7, 1.0}.Multi(t))
}

func hit_sphere(center point3, radius float64, r ray) float64 {
	oc := Sub(r.origin(), center)
	a := r.direction().LenSquared()
	half_b := Dot(oc, r.direction())
	c := oc.LenSquared() - radius*radius
	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-half_b - math.Sqrt(discriminant)) / a
	}
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
