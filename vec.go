package main

import (
	"fmt"
	"math"
)

type vec3 struct {
	e0 float64
	e1 float64
	e2 float64
}

func (v vec3) X() float64 {
	return v.e0
}

func (v vec3) Y() float64 {
	return v.e1
}

func (v vec3) Z() float64 {
	return v.e2
}

func (v *vec3) Neg() {
	v.e0 *= -1
	v.e1 *= -1
	v.e2 *= -1
}

func Add(v vec3, vv vec3) vec3 {
	return vec3{v.e0 + vv.e0, v.e1 + vv.e1, v.e2 + vv.e2}
}

func Sub(v vec3, vv vec3) vec3 {
	return vec3{v.e0 - vv.e0, v.e1 - vv.e1, v.e2 - vv.e2}
}

func (v vec3) Multi(t float64) vec3 {
	return vec3{v.e0 * t, v.e1 * t, v.e2 * t}
}

func Multi(v vec3, vv vec3) vec3 {
	return vec3{v.e0 * vv.e0, v.e1 * vv.e1, v.e2 * vv.e2}
}

func (v vec3) Div(t float64) vec3 {
	return v.Multi(1 / t)
}

func (v vec3) LenSquared() float64 {
	return v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2
}

func (v vec3) Len() float64 {
	return float64(math.Sqrt(float64(v.LenSquared())))
}

func Dot(v vec3, vv vec3) float64 {
	return v.e0*vv.e0 + v.e1*vv.e1 + v.e2*vv.e2
}

func Cross(v vec3, vv vec3) vec3 {
	return vec3{v.e1*vv.e2 - v.e2*vv.e1,
		v.e2*vv.e0 - v.e0*vv.e2,
		v.e0*vv.e1 - v.e1*vv.e0}
}

func UnitVector(v vec3) vec3 {
	return v.Div(v.Len())
}

// Implement Stringer interface
func (v vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.e0, v.e1, v.e2)
}

type point3 = vec3
type color = vec3
