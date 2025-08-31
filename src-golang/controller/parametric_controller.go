package controller

import (
	"Geometric_Construction/math_lib"
	"fmt"
)

type ParametricController struct{}

func NewParametricController() *ParametricController {
	return &ParametricController{}
}

// ProcessParametricFunction 处理参数方程函数，生成STL文件
// f: 参数方程函数，接受u,v参数，返回x,y,z坐标
// uRange: u参数范围 [umin, umax]
// vRange: v参数范围 [vmin, vmax]
// divisions: 分割数 [uDiv, vDiv]
// outputFile: 输出STL文件名
func (c *ParametricController) ProcessParametricFunction(f func(u, v float64) (x, y, z float64), uRange, vRange []float64, divisions []int, outputFile string) error {
	triangles := math_lib.TriangulateParametricSurface(f, uRange, vRange, divisions)
	if err := math_lib.SaveBinarySTL(triangles, outputFile); err != nil {
		return fmt.Errorf("保存STL文件失败: %v", err)
	}

	return nil
}