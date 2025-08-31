package math_lib

import (
	"gonum.org/v1/gonum/mat"
)

// TriangulateParametricSurface 对参数曲面进行三角化
func TriangulateParametricSurface(f func(u, v float64) (x, y, z float64), uRange, vRange []float64, divisions []int) []*Triangle {
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

// crossProduct 计算二维叉积
func crossProduct(a, b, c *mat.VecDense) float64 {
	return (b.At(0, 0)-a.At(0, 0))*(c.At(1, 0)-b.At(1, 0)) - (b.At(1, 0)-a.At(1, 0))*(c.At(0, 0)-b.At(0, 0))
}

// isEar 检查三角形是否是"耳朵"
func isEar(a, b, c *mat.VecDense, polygon []*mat.VecDense) bool {
	// 如果不是凸顶点
	if crossProduct(a, b, c) < 0 {
		return false
	}

	// 检查是否有其他顶点在三角形内
	for _, p := range polygon {
		if (p.At(0, 0) == a.At(0, 0) && p.At(1, 0) == a.At(1, 0)) ||
			(p.At(0, 0) == b.At(0, 0) && p.At(1, 0) == b.At(1, 0)) ||
			(p.At(0, 0) == c.At(0, 0) && p.At(1, 0) == c.At(1, 0)) {
			continue
		}

		// 如果在三角形内
		if crossProduct(a, b, p) > 0 &&
			crossProduct(b, c, p) > 0 &&
			crossProduct(c, a, p) > 0 {
			return false
		}
	}
	return true
}

// EarClippingTriangulation 执行耳切法三角剖分
func EarClippingTriangulation(polygon []*mat.VecDense) []Triangle {
	triangleSet := make([]Triangle, 0)
	poly := make([]*mat.VecDense, len(polygon))
	copy(poly, polygon)

	for len(poly) >= 3 {
		n := len(poly)

		for i := 0; i < n; i++ {
			a := (i + n - 1) % n
			c := (i + 1) % n

			if len(poly) == 3 || isEar(poly[a], poly[i], poly[c], poly) {
				// 创建三角形
				tri := Triangle{
					[3]*mat.VecDense{poly[a], poly[i], poly[c]},
				}
				triangleSet = append(triangleSet, tri)

				// 移除当前顶点
				poly = append(poly[:i], poly[i+1:]...)
				n--
				i-- // 调整索引
			}
		}
	}
	return triangleSet
}
