package math_lib

import (
	"gonum.org/v1/gonum/mat"
)

type Triangle struct {
	P [3]*mat.VecDense `json:"p"`
}

// GetNormal 计算三角形的法向量
func (f *Triangle) GetNormal() *mat.VecDense {
	edge1 := mat.NewVecDense(3, nil)
	edge2 := mat.NewVecDense(3, nil)
	edge1.SubVec(f.P[1], f.P[0])
	edge2.SubVec(f.P[2], f.P[0])
	return Normalize(Cross(mat.NewVecDense(3, nil), edge1, edge2))
}
