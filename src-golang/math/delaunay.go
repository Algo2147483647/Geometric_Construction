package math

import (
	"math"
	"sort"
)

// Point 表示二维点
type Point []float64

// Triangle 表示三角形，包含三个点
type Triangle []Point

// Delaunay 执行Delaunay三角剖分
func Delaunay(points []Point) []Triangle {
	// 第一步：按x坐标排序点集
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	// 第二步：计算超级三角形
	maxPoint := points[0]
	minPoint := points[0]

	for i := 1; i < len(points); i++ {
		if points[i][0] > maxPoint[0] || (points[i][0] == maxPoint[0] && points[i][1] > maxPoint[1]) {
			maxPoint = points[i]
		}
		if points[i][0] < minPoint[0] || (points[i][0] == minPoint[0] && points[i][1] < minPoint[1]) {
			minPoint = points[i]
		}
	}

	lengthX := maxPoint[0] - minPoint[0]
	lengthY := maxPoint[1] - minPoint[1]

	supertriangle := Triangle{
		{minPoint[0] - lengthX - 2, minPoint[1] - 2},
		{maxPoint[0] + lengthX + 2, minPoint[1] - 2},
		{(maxPoint[0] + minPoint[0]) / 2, maxPoint[1] + lengthY + 2},
	}

	triTemp := []Triangle{supertriangle}
	triAns := []Triangle{}
	edgeBuffer := []Edge{}

	// 第三步：逐点插入
	for _, point := range points {
		edgeBuffer = []Edge{}

		for j := 0; j < len(triTemp); j++ {
			triangle := triTemp[j]
			center, R := circumcircle(triangle[0], triangle[1], triangle[2])

			// 检查点是否在外接圆右侧
			if point[0] > center[0]+R {
				triAns = append(triAns, triangle)
				triTemp = append(triTemp[:j], triTemp[j+1:]...)
				j--
				continue
			}

			// 计算点到圆心的距离
			dx := point[0] - center[0]
			dy := point[1] - center[1]
			distance := math.Sqrt(dx*dx + dy*dy)

			// 如果点在圆内
			if distance < R {
				// 添加三角形的三条边到边缘缓冲区
				edgeBuffer = append(edgeBuffer, Edge{triangle[0], triangle[1]})
				edgeBuffer = append(edgeBuffer, Edge{triangle[1], triangle[2]})
				edgeBuffer = append(edgeBuffer, Edge{triangle[2], triangle[0]})

				triTemp = append(triTemp[:j], triTemp[j+1:]...)
				j--
			}
		}

		// 移除重复的边
		edgeBuffer = removeDuplicateEdges(edgeBuffer)

		// 为每条边创建新的三角形
		for _, edge := range edgeBuffer {
			triTemp = append(triTemp, Triangle{edge[0], edge[1], point})
		}
	}

	// 第四步：添加剩余的三角形
	triAns = append(triAns, triTemp...)

	// 第五步：移除包含超级三角形顶点的三角形
	finalTriangles := []Triangle{}
	for _, triangle := range triAns {
		keep := true
		for _, vertex := range triangle {
			if vertex[0] < minPoint[0] || vertex[1] < minPoint[1] ||
				vertex[0] > maxPoint[0] || vertex[1] > maxPoint[1] {
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

// Edge 表示边，包含两个点
type Edge [2]Point

// circumcircle 计算三点外接圆的圆心和半径
func circumcircle(a, b, c Point) (center Point, radius float64) {
	d := 2 * (a[0]*(b[1]-c[1]) + b[0]*(c[1]-a[1]) + c[0]*(a[1]-b[1]))

	ux := (a[0]*a[0]+a[1]*a[1])*(b[1]-c[1]) +
		(b[0]*b[0]+b[1]*b[1])*(c[1]-a[1]) +
		(c[0]*c[0]+c[1]*c[1])*(a[1]-b[1])
	ux /= d

	uy := (a[0]*a[0]+a[1]*a[1])*(c[0]-b[0]) +
		(b[0]*b[0]+b[1]*b[1])*(a[0]-c[0]) +
		(c[0]*c[0]+c[1]*c[1])*(b[0]-a[0])
	uy /= d

	center = Point{ux, uy}

	dx := a[0] - ux
	dy := a[1] - uy
	radius = math.Sqrt(dx*dx + dy*dy)

	return center, radius
}

// removeDuplicateEdges 移除重复的边
func removeDuplicateEdges(edges []Edge) []Edge {
	// 排序边
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][0][0] != edges[j][0][0] {
			return edges[i][0][0] < edges[j][0][0]
		}
		if edges[i][0][1] != edges[j][0][1] {
			return edges[i][0][1] < edges[j][0][1]
		}
		if edges[i][1][0] != edges[j][1][0] {
			return edges[i][1][0] < edges[j][1][0]
		}
		return edges[i][1][1] < edges[j][1][1]
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
func pointsEqual(p1, p2 Point) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}
