package main

type hit_record struct {
	p          point3
	normal     vec3
	t          float64
	front_face bool
}

func (h hit_record) hit(r ray, t_min, t_max float64, rec hit_record) bool {
	return false
}

func (h hit_record) set_face_normal(r ray, outward_normal vec3) {
	h.front_face = Dot(r.direction(), outward_normal) < 0
	if h.front_face {
		h.normal = outward_normal
	} else {
		h.normal = outward_normal.Multi(-1)
	}
}
