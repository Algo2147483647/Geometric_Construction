package math_lib

import "gonum.org/v1/gonum/mat"

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
