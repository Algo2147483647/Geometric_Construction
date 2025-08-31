package controller

import (
	"Geometric_Construction/math_lib"
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type ImplicitFunctionController struct{}

func NewImplicitFunctionController() *ImplicitFunctionController {
	return &ImplicitFunctionController{}
}

func (c *ImplicitFunctionController) ProcessImplicitFunction(f func(*mat.VecDense) float64, st, ed []float64, N []int, outputFile string) error {
	triangles := math_lib.MarchingCubes(f, st, ed, N)
	if err := math_lib.SaveBinarySTL(triangles, outputFile); err != nil {
		return fmt.Errorf("保存STL文件失败: %v", err)
	}

	return nil
}
