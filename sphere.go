package main

import "math"

type sphere struct {
	center point3
	radius float64
}

func (s sphere) hit(r ray, t_min, t_max float64, rec hit_record) bool {
	oc := Sub(r.origin(), s.center)
	a := r.direction().LenSquared()
	half_b := Dot(oc, r.direction())
	c := oc.LenSquared() - s.radius*s.radius

	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	root := (-half_b - sqrtd) / a
	if root < t_min || t_max < root {
		root = (-half_b + sqrtd) / a
		if root < t_min || t_max < root {
			return false
		}
	}
	rec.t = root
	rec.p = r.at(rec.t)
	outward_normal := Sub(rec.p, s.center).Div(s.radius)
	rec.set_face_normal(r, outward_normal)
	return true
}
