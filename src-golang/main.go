package main

import (
	"Geometric_Construction/controller"
	"fmt"
)

func main() {
	//// 创建隐函数控制器
	//implicitCtrl := controller.NewImplicitFunctionController()
	//
	//f := func(point *mat.VecDense) float64 {
	//	x, y, z := point.AtVec(0), point.AtVec(1), point.AtVec(2)
	//	r := math.Sqrt(x*x + y*y + z*z)
	//	return math.Pow(x, 4) + math.Pow(y, 4) + math.Pow(z, 4) -
	//		2.5*(x*x*y*y+y*y*z*z+z*z*x*x) +
	//		0.5*math.Cos(3*x)*math.Sin(4*y)*math.Exp(-0.1*(z*z)) -
	//		math.Tanh(r)*math.Abs(math.Sin(2*x)*math.Cos(2*y)*math.Sin(2*z))
	//}
	//
	//// 示例2: 生成环面STL文件
	//fmt.Println("生成环面STL文件...")
	//a := 2.0
	//b := 100
	//err := implicitCtrl.ProcessImplicitFunction(
	//	f,                     // 大半径1.0，小半径0.3的环面
	//	[]float64{-a, -a, -a}, // 起始坐标
	//	[]float64{a, a, a},    // 结束坐标
	//	[]int{b, b, b},        // 网格划分
	//	"torus.stl")           // 输出文件名
	//if err != nil {
	//	fmt.Printf("生成环面STL文件时出错: %v\n", err)
	//}

	// 创建参数方程控制器
	parametricCtrl := controller.NewParametricController()
	a := 2.0
	b := 100

	// 示例：生成球面的参数方程
	Func := func(u, v float64) (x, y, z float64) {
		x = u - u*u*u/3 + u*v*v
		y = v - v*v*v/3 + v*u*u
		z = u*u - v*v
		return
	}

	fmt.Println("生成球面STL文件...")
	err := parametricCtrl.ProcessParametricFunction(
		Func,             // 球面参数方程
		[]float64{-a, a}, // u参数范围 [0, 2π]
		[]float64{-a, a}, // v参数范围 [0, π]
		[]int{b, b},      // u和v方向的分割数
		"sphere.stl")     // 输出文件名
	if err != nil {
		fmt.Printf("生成球面STL文件时出错: %v\n", err)
	}

	fmt.Println("所有STL文件生成完成!")
}
