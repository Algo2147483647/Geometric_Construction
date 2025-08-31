package main

import (
	"Geometric_Construction/application"
	"Geometric_Construction/example_library"
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

	h := application.NewHandler()
	a := 2.0
	var err error

	h.TriangulateParametricEquation(
		example_library.Camellia, // 球面参数方程
		[]float64{0, a},          // u参数范围 [0, 2π]
		[]float64{0, 1},          // v参数范围 [0, π]
		[]int{10000, 100},        // u和v方向的分割数
	)
	if err != nil {
		panic(err)
	}

	err = application.SaveBinarySTL(h.Triangles, "sphere.stl")
	if err != nil {
		panic(err)
	}
}
