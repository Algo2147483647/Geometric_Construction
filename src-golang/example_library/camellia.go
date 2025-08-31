package example_library

import (
	"gonum.org/v1/gonum/mat"
	"math"
)

func Camellia(u, v float64) (x, y, z float64) {
	theta := 4*math.Pi + u*20*math.Pi
	r := v
	modVal := math.Mod(3.6*theta, 2*math.Pi)
	disturb := math.Sin(15*theta) / 150
	edge := 1 - 0.5*math.Pow(1-modVal/math.Pi, 4) + disturb
	f2 := 2 * math.Pow(r*r-r, 2)
	alpha := (math.Pi / 2) * math.Exp(-theta/(8*math.Pi))
	h := f2 * math.Sin(alpha)
	sinAlpha := math.Sin(alpha)
	cosAlpha := math.Cos(alpha)
	R := sinAlpha*r + cosAlpha*h
	H := cosAlpha*r - sinAlpha*h
	x = edge * R * math.Cos(theta)
	y = edge * R * math.Sin(theta)
	z = edge * H
	return
}

// Sphere x² + y² + z² - r² = 0
func Sphere(radius float64) func(*mat.VecDense) float64 {
	return func(point *mat.VecDense) float64 {
		x, y, z := point.AtVec(0), point.AtVec(1), point.AtVec(2)
		return x*x + y*y + z*z - radius*radius
	}
}

func Torus(R, r float64) func(*mat.VecDense) float64 {
	return func(point *mat.VecDense) float64 {
		x, y, z := point.AtVec(0), point.AtVec(1), point.AtVec(2)
		return math.Pow(R-math.Sqrt(x*x+y*y), 2) + z*z - r*r
	}
}
