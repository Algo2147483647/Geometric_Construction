package math

// Point 表示三维点
type Point [3]float64

// Triangle 表示三角形，包含9个浮点数（3个顶点，每个顶点3个坐标）
type Triangle [9]float64

// crossProduct 计算二维叉积
func crossProduct(a, b, c Point) float64 {
	return (b[0]-a[0])*(c[1]-b[1]) - (b[1]-a[1])*(c[0]-b[0])
}

// isEar 检查三角形是否是"耳朵"
func isEar(a, b, c Point, polygon []Point) bool {
	// 如果不是凸顶点
	if crossProduct(a, b, c) < 0 {
		return false
	}

	// 检查是否有其他顶点在三角形内
	for _, p := range polygon {
		if (p[0] == a[0] && p[1] == a[1]) ||
			(p[0] == b[0] && p[1] == b[1]) ||
			(p[0] == c[0] && p[1] == c[1]) {
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
func EarClippingTriangulation(polygon []Point) []Triangle {
	triangleSet := make([]Triangle, 0)
	poly := make([]Point, len(polygon))
	copy(poly, polygon)

	for len(poly) >= 3 {
		n := len(poly)

		for i := 0; i < n; i++ {
			a := (i + n - 1) % n
			c := (i + 1) % n

			if len(poly) == 3 || isEar(poly[a], poly[i], poly[c], poly) {
				// 创建三角形
				tri := Triangle{
					poly[a][0], poly[a][1], 0,
					poly[i][0], poly[i][1], 0,
					poly[c][0], poly[c][1], 0,
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
