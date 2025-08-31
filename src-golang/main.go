package main

import (
	"Geometric_Construction/controller"
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// 创建隐函数控制器
	ctrl := controller.NewImplicitFunctionController()

	// 示例1: 生成球体STL文件
	fmt.Println("生成球体STL文件...")
	err := ctrl.ProcessImplicitFunction(
		ctrl.SphereFunction(1.0),    // 半径为1的球体
		[]float64{-1.5, -1.5, -1.5}, // 起始坐标
		[]float64{1.5, 1.5, 1.5},    // 结束坐标
		[]int{50, 50, 50},           // 网格划分
		"sphere.stl")                // 输出文件名
	if err != nil {
		fmt.Printf("生成球体STL文件时出错: %v\n", err)
	}

	// 示例2: 生成环面STL文件
	fmt.Println("生成环面STL文件...")
	err = ctrl.ProcessImplicitFunction(
		ctrl.TorusFunction(1.0, 0.3), // 大半径1.0，小半径0.3的环面
		[]float64{-1.5, -1.5, -1.5},  // 起始坐标
		[]float64{1.5, 1.5, 1.5},     // 结束坐标
		[]int{50, 50, 50},            // 网格划分
		"torus.stl")                  // 输出文件名
	if err != nil {
		fmt.Printf("生成环面STL文件时出错: %v\n", err)
	}

	// 示例3: 自定义隐函数 - 两个球体的并集
	fmt.Println("生成双球体STL文件...")
	doubleSphere := func(point *mat.VecDense) float64 {
		x, y, z := point.AtVec(0), point.AtVec(1), point.AtVec(2)
		// 两个球体的并集
		d1 := math.Sqrt((x-0.5)*(x-0.5)+y*y+z*z) - 0.6
		d2 := math.Sqrt((x+0.5)*(x+0.5)+y*y+z*z) - 0.6
		return math.Min(d1, d2)
	}

	err = ctrl.ProcessImplicitFunction(
		doubleSphere,                // 双球体隐函数
		[]float64{-2.0, -2.0, -2.0}, // 起始坐标
		[]float64{2.0, 2.0, 2.0},    // 结束坐标
		[]int{60, 60, 60},           // 网格划分
		"double_sphere.stl")         // 输出文件名
	if err != nil {
		fmt.Printf("生成双球体STL文件时出错: %v\n", err)
	}

	fmt.Println("所有STL文件生成完成!")
}
