# Geometric_Construction

[TOC]

## Introduction

***Implicit Equation | Camellia***: 

$$f: (r, \theta) \to \mathbb R^3$$

$$
\begin{align*}
(r, \theta) &\in ([0, 1], [4\pi, 24 \pi])  \tag{参数空间}\\
edge(\theta) &= 1-\frac{1}{2}\left(1-\frac{\mod(3.6 \theta, 2\pi)}{\pi}\right)^4 + disturb(\theta)   \tag{花边}\\
disturb(\theta) &= \frac{\sin(15 \theta)}{150}   \tag{花边扰动}\\
f_2(r)&=2(r^2-r)^2 \tag{花瓣弧面}\\
\alpha(\theta) &= \frac{\pi}{2} \cdot e^{-\frac{\theta}{8\pi}}  \tag{倾斜角衰减}\\
h(r,\theta)&=f_2(r) \sin(\alpha(\theta))  \tag{弧面衰减}\\
\left(\begin{matrix}R\\ H\end{matrix}\right) &= \left(\begin{matrix}\sin\alpha(\theta) & \cos\alpha(\theta) \\ \cos\alpha(\theta) & -\sin\alpha(\theta) \end{matrix}\right) \left(\begin{matrix}r\\  h(r,\theta)\end{matrix}\right) \tag{旋转矩阵}\\
\left(\begin{matrix}X\\ Y\\ Z\end{matrix}\right) &= \left(\begin{matrix}edge(\theta) \cdot R \cos\theta\\ edge(\theta) \cdot R \sin\theta\\  edge(\theta) \cdot H\end{matrix}\right)  \tag{3D欧式空间映射}
\end{align*}
$$
<img src="./docs/assets/db54bafc40ecf0743799a487eb9f812.jpg" alt="db54bafc40ecf0743799a487eb9f812" style="zoom:50%;" />

### System Overview



### System Objectives



## System Architecture


### Core Components Detailed

The system supports various geometric shapes:
| Shape                                        |                          Expression                          |
| :------------------------------------------- | :----------------------------------------------------------: |
| ***Implicit Equation***                      |                    $f(\boldsymbol x) = 0$                    |
| ***Parametric Equation***                    |                 $f: (u, v) \to \mathbb R^3$                  |


### Workflow



### Module Descriptions


### Technical Features



## Usage Instructions

### Environment Requirements

- Go 1.23.0
- Python 3.x (for visualization)

### Scene Script Format




### Building and Running

### Visualization Interface

### Development Guide



#### Performance Optimization




## TODO
