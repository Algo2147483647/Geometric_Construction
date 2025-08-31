package controller

import (
	"Geometric_Construction/math_lib"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

// ImplicitFunctionController 隐函数控制器
type ImplicitFunctionController struct{}

// NewImplicitFunctionController 创建新的隐函数控制器
func NewImplicitFunctionController() *ImplicitFunctionController {
	return &ImplicitFunctionController{}
}

// SphereFunction 球体隐函数示例: x² + y² + z² - r² = 0
func (c *ImplicitFunctionController) SphereFunction(radius float64) func(*mat.VecDense) float64 {
	return func(point *mat.VecDense) float64 {
		x, y, z := point.AtVec(0), point.AtVec(1), point.AtVec(2)
		return x*x + y*y + z*z - radius*radius
	}
}

// TorusFunction 环面隐函数示例
func (c *ImplicitFunctionController) TorusFunction(R, r float64) func(*mat.VecDense) float64 {
	return func(point *mat.VecDense) float64 {
		x, y, z := point.AtVec(0), point.AtVec(1), point.AtVec(2)
		return math.Pow(R-math.Sqrt(x*x+y*y), 2) + z*z - r*r
	}
}

// ProcessImplicitFunction 处理隐函数并生成STL文件
func (c *ImplicitFunctionController) ProcessImplicitFunction(
	f func(*mat.VecDense) float64,
	st []float64, // 起始坐标 [x, y, z]
	ed []float64, // 结束坐标 [x, y, z]
	N []int, // 网格划分 [nx, ny, nz]
	outputFile string) error {

	// 使用Marching Cubes算法计算三角形
	triangles := math_lib.MarchingCubes(f, st, ed, N)

	// 转换为Triangle结构体切片
	triangleList := make([]math_lib.Triangle, len(triangles))
	for i, t := range triangles {
		triangleList[i] = *t
	}

	// 保存为STL文件
	if err := math_lib.SaveBinarySTL(triangleList, outputFile); err != nil {
		return fmt.Errorf("保存STL文件失败: %v", err)
	}

	fmt.Printf("成功生成STL文件: %s，共%d个三角形\n", outputFile, len(triangleList))
	return nil
}
