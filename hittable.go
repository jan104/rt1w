package main

type hit_record struct {
	p      point3
	normal vec3
	t      float64
}

func (h hit_record) hit(r ray, t_min, t_max float64, rec hit_record) bool {
	return false
}
