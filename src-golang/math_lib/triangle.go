package math_lib

import (
	"encoding/binary"
	"gonum.org/v1/gonum/mat"
	"io"
	"math"
	"os"
)

type Triangle struct {
	P [3]*mat.VecDense `json:"p"`
}

// SaveBinarySTL 将三角形切片保存为二进制STL文件
func SaveBinarySTL(triangles []Triangle, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入80字节的头部信息（通常包含描述性信息）
	header := make([]byte, 80)
	copy(header, "Gonum STL Export")
	if _, err := file.Write(header); err != nil {
		return err
	}

	// 写入三角形数量（4字节小端序无符号整数）
	triangleCount := uint32(len(triangles))
	if err := binary.Write(file, binary.LittleEndian, triangleCount); err != nil {
		return err
	}

	// 写入每个三角形数据
	for _, tri := range triangles {
		// 计算法向量（使用右手定则，从第一个点开始按顺序）
		normal := calculateNormal(tri.P[0], tri.P[1], tri.P[2])

		// 写入法向量（3个float32）
		if err := writeVectorAsFloat32(file, normal); err != nil {
			return err
		}

		// 写入三个顶点（每个顶点3个float32）
		for i := 0; i < 3; i++ {
			if err := writeVectorAsFloat32(file, tri.P[i]); err != nil {
				return err
			}
		}

		// 写入属性字节计数（通常为0，2字节）
		attribute := uint16(0)
		if err := binary.Write(file, binary.LittleEndian, attribute); err != nil {
			return err
		}
	}

	return nil
}

// calculateNormal 计算三角形的法向量
func calculateNormal(v1, v2, v3 *mat.VecDense) *mat.VecDense {
	// 计算两个边向量
	edge1 := mat.NewVecDense(3, nil)
	edge1.SubVec(v2, v1)

	edge2 := mat.NewVecDense(3, nil)
	edge2.SubVec(v3, v1)

	// 计算叉积（法向量）
	normal := mat.NewVecDense(3, nil)
	normal.SetVec(0, edge1.At(1, 0)*edge2.At(2, 0)-edge1.At(2, 0)*edge2.At(1, 0))
	normal.SetVec(1, edge1.At(2, 0)*edge2.At(0, 0)-edge1.At(0, 0)*edge2.At(2, 0))
	normal.SetVec(2, edge1.At(0, 0)*edge2.At(1, 0)-edge1.At(1, 0)*edge2.At(0, 0))

	// 归一化法向量
	length := math.Sqrt(
		normal.At(0, 0)*normal.At(0, 0) +
			normal.At(1, 0)*normal.At(1, 0) +
			normal.At(2, 0)*normal.At(2, 0))

	if length > 0 {
		normal.ScaleVec(1/length, normal)
	}

	return normal
}

// writeVectorAsFloat32 将mat.VecDense向量以float32格式写入
func writeVectorAsFloat32(w io.Writer, vec *mat.VecDense) error {
	for i := 0; i < 3; i++ {
		value := float32(vec.At(i, 0))
		if err := binary.Write(w, binary.LittleEndian, value); err != nil {
			return err
		}
	}
	return nil
}
