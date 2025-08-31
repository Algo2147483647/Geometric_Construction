package application

import (
	"Geometric_Construction/math_lib"
	"gonum.org/v1/gonum/mat"
)

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) TriangulateParametricEquation(f func(u, v float64) (x, y, z float64), uRange, vRange []float64, divisions []int) *Handler {
	if h.error != nil {
		return h
	}

	res := math_lib.TriangulateParametricEquation(f, uRange, vRange, divisions)
	h.Triangles = append(h.Triangles, res...)
	return h
}

func (h *Handler) TriangulateImplicitEquation(f func(*mat.VecDense) float64, st, ed []float64, N []int) *Handler {
	if h.error != nil {
		return h
	}

	res := math_lib.MarchingCubes(f, st, ed, N)
	h.Triangles = append(h.Triangles, res...)
	return h
}

func (h *Handler) Transform(f func(*mat.VecDense) float64, st, ed []float64, N []int) *Handler {
	if h.error != nil {
		return h
	}

	return h
}
