package application

import (
	"Geometric_Construction/math_lib"
	"encoding/binary"
	"gonum.org/v1/gonum/mat"
	"io"
	"os"
)

// SaveBinarySTL 将三角形切片保存为二进制STL文件
func SaveBinarySTL(f []*math_lib.Triangle, filename string) error {
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
	triangleCount := uint32(len(f))
	if err := binary.Write(file, binary.LittleEndian, triangleCount); err != nil {
		return err
	}

	// 写入每个三角形数据
	for _, tri := range f {
		normal := tri.GetNormal()                                 // 计算法向量（使用右手定则，从第一个点开始按顺序）
		if err = writeVectorAsFloat32(file, normal); err != nil { // 写入法向量（3个float32）
			return err
		}
		for i := 0; i < 3; i++ { // 写入三个顶点（每个顶点3个float32）
			if err = writeVectorAsFloat32(file, tri.P[i]); err != nil {
				return err
			}
		}
		attribute := uint16(0) // 写入属性字节计数（通常为0，2字节）
		if err = binary.Write(file, binary.LittleEndian, attribute); err != nil {
			return err
		}
	}

	return nil
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
