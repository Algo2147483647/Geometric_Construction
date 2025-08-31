package math_lib

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"sort"
)

// Edge 表示边，包含两个点
type Edge [2]*mat.VecDense

// Delaunay 执行Delaunay三角剖分
func Delaunay(points []*mat.VecDense) []Triangle {
	// 第一步：按x坐标排序点集
	sort.Slice(points, func(i, j int) bool {
		if points[i].At(0, 0) != points[j].At(0, 0) {
			return points[i].At(0, 0) < points[j].At(0, 0)
		}
		return points[i].At(1, 0) < points[j].At(1, 0)
	})

	// 第二步：计算超级三角形
	maxPoint := points[0]
	minPoint := points[0]

	for i := 1; i < len(points); i++ {
		if points[i].At(0, 0) > maxPoint.At(0, 0) || (points[i].At(0, 0) == maxPoint.At(0, 0) && points[i].At(1, 0) > maxPoint.At(1, 0)) {
			maxPoint = points[i]
		}
		if points[i].At(0, 0) < minPoint.At(0, 0) || (points[i].At(0, 0) == minPoint.At(0, 0) && points[i].At(1, 0) < minPoint.At(1, 0)) {
			minPoint = points[i]
		}
	}

	lengthX := maxPoint.At(0, 0) - minPoint.At(0, 0)
	lengthY := maxPoint.At(1, 0) - minPoint.At(1, 0)

	supertriangle := Triangle{
		[3]*mat.VecDense{
			mat.NewVecDense(3, []float64{minPoint.At(0, 0) - lengthX - 2, minPoint.At(1, 0) - 2, 0}),
			mat.NewVecDense(3, []float64{maxPoint.At(0, 0) + lengthX + 2, minPoint.At(1, 0) - 2, 0}),
			mat.NewVecDense(3, []float64{(maxPoint.At(0, 0) + minPoint.At(0, 0)) / 2, maxPoint.At(1, 0) + lengthY + 2, 0}),
		},
	}

	triTemp := []Triangle{supertriangle}
	triAns := []Triangle{}
	edgeBuffer := []Edge{}

	// 第三步：逐点插入
	for _, point := range points {
		edgeBuffer = []Edge{}

		for j := 0; j < len(triTemp); j++ {
			triangle := triTemp[j]
			center, R := circumcircle(triangle.P[0], triangle.P[1], triangle.P[2])

			// 检查点是否在外接圆右侧
			if point.At(0, 0) > center.At(0, 0)+R {
				triAns = append(triAns, triangle)
				triTemp = append(triTemp[:j], triTemp[j+1:]...)
				j--
				continue
			}

			// 计算点到圆心的距离
			dx := point.At(0, 0) - center.At(0, 0)
			dy := point.At(1, 0) - center.At(1, 0)
			distance := math.Sqrt(dx*dx + dy*dy)

			// 如果点在圆内
			if distance < R {
				// 添加三角形的三条边到边缘缓冲区
				edgeBuffer = append(edgeBuffer, Edge{triangle.P[0], triangle.P[1]})
				edgeBuffer = append(edgeBuffer, Edge{triangle.P[1], triangle.P[2]})
				edgeBuffer = append(edgeBuffer, Edge{triangle.P[2], triangle.P[0]})

				triTemp = append(triTemp[:j], triTemp[j+1:]...)
				j--
			}
		}

		// 移除重复的边
		edgeBuffer = removeDuplicateEdges(edgeBuffer)

		// 为每条边创建新的三角形
		for _, edge := range edgeBuffer {
			triTemp = append(triTemp, Triangle{[3]*mat.VecDense{edge[0], edge[1], point}})
		}
	}

	// 第四步：添加剩余的三角形
	triAns = append(triAns, triTemp...)

	// 第五步：移除包含超级三角形顶点的三角形
	finalTriangles := []Triangle{}
	for _, triangle := range triAns {
		keep := true
		for _, vertex := range triangle.P {
			if vertex.At(0, 0) < minPoint.At(0, 0) || vertex.At(1, 0) < minPoint.At(1, 0) ||
				vertex.At(0, 0) > maxPoint.At(0, 0) || vertex.At(1, 0) > maxPoint.At(1, 0) {
				keep = false
				break
			}
		}
		if keep {
			finalTriangles = append(finalTriangles, triangle)
		}
	}

	return finalTriangles
}

// circumcircle 计算三点外接圆的圆心和半径
func circumcircle(a, b, c *mat.VecDense) (center *mat.VecDense, radius float64) {
	d := 2 * (a.At(0, 0)*(b.At(1, 0)-c.At(1, 0)) + b.At(0, 0)*(c.At(1, 0)-a.At(1, 0)) + c.At(0, 0)*(a.At(1, 0)-b.At(1, 0)))

	ux := (a.At(0, 0)*a.At(0, 0)+a.At(1, 0)*a.At(1, 0))*(b.At(1, 0)-c.At(1, 0)) +
		(b.At(0, 0)*b.At(0, 0)+b.At(1, 0)*b.At(1, 0))*(c.At(1, 0)-a.At(1, 0)) +
		(c.At(0, 0)*c.At(0, 0)+c.At(1, 0)*c.At(1, 0))*(a.At(1, 0)-b.At(1, 0))
	ux /= d

	uy := (a.At(0, 0)*a.At(0, 0)+a.At(1, 0)*a.At(1, 0))*(c.At(0, 0)-b.At(0, 0)) +
		(b.At(0, 0)*b.At(0, 0)+b.At(1, 0)*b.At(1, 0))*(a.At(0, 0)-c.At(0, 0)) +
		(c.At(0, 0)*c.At(0, 0)+c.At(1, 0)*c.At(1, 0))*(b.At(0, 0)-a.At(0, 0))
	uy /= d

	center = mat.NewVecDense(3, []float64{ux, uy, 0})

	dx := a.At(0, 0) - ux
	dy := a.At(1, 0) - uy
	radius = math.Sqrt(dx*dx + dy*dy)

	return center, radius
}

// removeDuplicateEdges 移除重复的边
func removeDuplicateEdges(edges []Edge) []Edge {
	// 排序边
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][0].At(0, 0) != edges[j][0].At(0, 0) {
			return edges[i][0].At(0, 0) < edges[j][0].At(0, 0)
		}
		if edges[i][0].At(1, 0) != edges[j][0].At(1, 0) {
			return edges[i][0].At(1, 0) < edges[j][0].At(1, 0)
		}
		if edges[i][1].At(0, 0) != edges[j][1].At(0, 0) {
			return edges[i][1].At(0, 0) < edges[j][1].At(0, 0)
		}
		return edges[i][1].At(1, 0) < edges[j][1].At(1, 0)
	})

	// 移除重复的边
	result := []Edge{}
	for i := 0; i < len(edges); i++ {
		if i < len(edges)-1 && edgesEqual(edges[i], edges[i+1]) {
			i++ // 跳过下一个边（重复的）
			continue
		}
		result = append(result, edges[i])
	}

	return result
}

// edgesEqual 检查两条边是否相等
func edgesEqual(e1, e2 Edge) bool {
	return (pointsEqual(e1[0], e2[0]) && pointsEqual(e1[1], e2[1])) ||
		(pointsEqual(e1[0], e2[1]) && pointsEqual(e1[1], e2[0]))
}

// pointsEqual 检查两个点是否相等
func pointsEqual(p1, p2 *mat.VecDense) bool {
	return p1.At(0, 0) == p2.At(0, 0) && p1.At(1, 0) == p2.At(1, 0)
}
