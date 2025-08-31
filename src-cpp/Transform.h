#ifndef GRAPHICS_TRANSFORM_H
#define GRAPHICS_TRANSFORM_H

#include "Transform.h"
#include "Submatrix.h"
#include "RGB.h"
#include "Graphics2D.h"

namespace Graphics {
	extern Mat<float> TransformMat;

	/*---------------- 几何变换 Transformation ----------------*/
	Mat<float>& translate(Mat<float>& delta);						//平移
	Mat<float>& rotate	(Mat<float>& theta, Mat<float>& center);		//旋转
	Mat<float>& rotate	(Mat<float>& rotateAxis, double theta, Mat<float>& center);	//旋转 3D
	Mat<float>& scale	(Mat<float>& ratio,Mat<float>& center);		//缩放
	Mat<float>& reflect	(Mat<float>& e, Mat<float>& center);			//镜像
}

#endif