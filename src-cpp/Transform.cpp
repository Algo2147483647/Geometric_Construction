#include "Transform.h"

Mat<float> Graphics::TransformMat(4, 4);

/*#############################################################################

*                    Transformation 

##############################################################################*/

// 平移
Mat<float>& Graphics::translate(Mat<float>& delta) {
	static Mat<float> translateMat;
	Matrix::E(translateMat.alloc(TransformMat.rows, TransformMat.cols));

	for (int i = 0; i < delta.rows; i++) 
		translateMat(i + 1, 0) = delta[i];

	Matrix::mul(TransformMat, translateMat, TransformMat);
	return TransformMat;
}

// 旋转
Mat<float>& Graphics::rotate(Mat<float>& theta, Mat<float>& center) {
	static Mat<float> tmp, rotateMat(TransformMat.rows - 1, TransformMat.cols - 1);
	translate(Matrix::negative(tmp, center));

	Matrix::rotate(theta, rotateMat);
	Matrix::E(tmp.alloc(TransformMat.rows, TransformMat.cols));
	Matrix::setBlock(tmp, rotateMat, 1, 1);

	Matrix::mul(TransformMat, tmp, TransformMat);
	translate(center);
	return TransformMat;
}

// 3D·四元数
Mat<float>& Graphics::rotate(Mat<float>& rotateAxis, double theta, Mat<float>& center) {
	if (TransformMat.rows != 4 || TransformMat.cols != 4)
		exit(-1);

	static Mat<float> tmp, rotateMat;
	translate(Matrix::negative(tmp, center));

	Matrix::rotate(rotateAxis, theta, rotateMat);

	Matrix::mul(TransformMat, rotateMat, TransformMat);
	translate(center);
	return TransformMat;
}

// 缩放
Mat<float>& Graphics::scale(Mat<float>& ratio, Mat<float>& center) {
	static Mat<float> tmp, scaleMat;
	translate(Matrix::negative(tmp, center));

	Matrix::scale(ratio, scaleMat);
	Matrix::E(tmp.alloc(TransformMat.rows, TransformMat.cols));
	Matrix::setBlock(tmp, scaleMat, 1, 1);

	Matrix::mul(TransformMat, tmp, TransformMat);
	translate(center);
	return TransformMat;
}

// 镜像
Mat<float>& Graphics::reflect(Mat<float>& e, Mat<float>& center) {
	Mat<float> tmp, reflectMat;
	translate(Matrix::negative(tmp, center));

	Matrix::reflect(e, reflectMat);
	Matrix::E(tmp.alloc(TransformMat.rows, TransformMat.cols));
	Matrix::setBlock(tmp, reflectMat, 1, 1);

	Matrix::mul(TransformMat, tmp, TransformMat);
	translate(center);
	return TransformMat;
}

// 透视投影
/*
Mat<float>& Graphics::perspect(Mat<float>& e, Mat<float>& center) {

}*/