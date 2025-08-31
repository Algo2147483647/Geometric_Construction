package math_lib

import "gonum.org/v1/gonum/mat"

// TriangulateParametricSurface 对参数曲面进行三角化
func TriangulateParametricEquation(f func(u, v float64) (x, y, z float64), uRange, vRange []float64, divisions []int) []*Triangle {
	// 计算步长
	uStep := (uRange[1] - uRange[0]) / float64(divisions[0])
	vStep := (vRange[1] - vRange[0]) / float64(divisions[1])

	// 存储顶点
	vertices := make([][]*mat.VecDense, divisions[0]+1)
	for i := range vertices {
		vertices[i] = make([]*mat.VecDense, divisions[1]+1)
	}

	// 生成顶点
	for i := 0; i <= divisions[0]; i++ {
		for j := 0; j <= divisions[1]; j++ {
			u := uRange[0] + float64(i)*uStep
			v := vRange[0] + float64(j)*vStep
			x, y, z := f(u, v)
			vertices[i][j] = mat.NewVecDense(3, []float64{x, y, z})
		}
	}

	// 生成三角形
	var triangles []*Triangle
	for i := 0; i < divisions[0]; i++ {
		for j := 0; j < divisions[1]; j++ {
			// 每个网格单元生成两个三角形

			// 第一个三角形
			tri1 := &Triangle{}
			tri1.P[0] = vertices[i][j]
			tri1.P[1] = vertices[i+1][j]
			tri1.P[2] = vertices[i][j+1]
			triangles = append(triangles, tri1)

			// 第二个三角形
			tri2 := &Triangle{}
			tri2.P[0] = vertices[i+1][j]
			tri2.P[1] = vertices[i+1][j+1]
			tri2.P[2] = vertices[i][j+1]
			triangles = append(triangles, tri2)
		}
	}

	return triangles
}
