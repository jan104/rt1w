package main

type ray struct {
	orig point3
	dir  vec3
}

func (r ray) origin() point3 {
	return r.orig
}

func (r ray) direction() vec3 {
	return r.dir
}

func (r ray) at(t float64) point3 {
	// Horrible
	return Add(r.orig, r.dir.Multi(t))
}
