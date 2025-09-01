import numpy as np
import matplotlib
import matplotlib.pyplot as plt
from scipy.special import sph_harm
import torch
matplotlib.use('TkAgg')
# 设置Matplotlib的3D显示
from mpl_toolkits.mplot3d import Axes3D

plt.rcParams['figure.figsize'] = (12, 8)


def load_bunny_point_cloud(gz_file_path="C:\\Users\\29753\\Downloads\\bunny.tar.gz"):
    """加载斯坦福兔子点云数据"""
    # 如果提供了 gz 文件路径，直接使用该文件
    if gz_file_path:
        print(f"使用提供的 gz 文件: {gz_file_path}")
        try:
            # 解压并读取PLY文件
            import tarfile
            with tarfile.open(gz_file_path, 'r:gz') as tar:
                for member in tar.getmembers():
                    if member.name.endswith('bun_zipper.ply'):
                        ply_file = tar.extractfile(member)
                        import plyfile
                        ply_data = plyfile.PlyData.read(ply_file)
                        vertices = ply_data['vertex']
                        points = np.vstack([vertices['x'], vertices['y'], vertices['z']]).T
                        return points
        except Exception as e:
            print(f"解析提供的 gz 文件失败: {e}")
            print("尝试从网络下载...")

    # 如果没有提供 gz 文件路径或解析失败，尝试从网络下载
    try:
        import urllib.request
        import gzip
        import tempfile
        import tarfile
        import plyfile

        # 下载兔子模型
        url = "http://graphics.stanford.edu/pub/3Dscanrep/bunny.tar.gz"
        print("从网络下载兔子模型...")
        with urllib.request.urlopen(url) as response:
            with tempfile.NamedTemporaryFile(suffix='.tar.gz', delete=False) as tmp_gz:
                tmp_gz.write(response.read())
                gz_path = tmp_gz.name
                print(f"下载完成，临时文件: {gz_path}")

        # 解压并读取PLY文件
        with tarfile.open(gz_path, 'r:gz') as tar:
            for member in tar.getmembers():
                if 'bunny.ply' in member.name.lower():
                    ply_file = tar.extractfile(member)
                    ply_data = plyfile.PlyData.read(ply_file)
                    vertices = ply_data['vertex']
                    points = np.vstack([vertices['x'], vertices['y'], vertices['z']]).T
                    return points

    except Exception as e:
        print(f"下载或解析兔子模型失败: {e}")

    # 如果上述方法都失败，使用内置的简单兔子模型
    print("使用内置简单兔子模型")
    return create_simple_bunny()

def create_simple_bunny():
    """创建简单的兔子点云（备选方案）"""
    # 生成参数化表面
    u = np.linspace(0, 2 * np.pi, 100)
    v = np.linspace(0, np.pi, 50)
    u, v = np.meshgrid(u, v)

    # 兔子形状参数方程
    x = 0.8 * np.cos(u) * (1 + 0.3 * np.cos(v))
    y = 0.6 * np.sin(u) * (1 + 0.3 * np.cos(v))
    z = 0.5 * (np.sin(v) - 0.5) + 0.2 * np.cos(3 * u) * np.sin(5 * v)

    # 添加耳朵
    ear_u = np.linspace(np.pi / 4, 3 * np.pi / 4, 20)
    ear_v = np.linspace(np.pi / 3, np.pi / 2, 10)
    ear_u, ear_v = np.meshgrid(ear_u, ear_v)

    ear_z = 0.4 + 0.3 * ear_v
    ear_x = 0.3 * np.cos(ear_u)
    ear_y = 0.3 * np.sin(ear_u)

    points = np.vstack([
        np.stack([x.ravel(), y.ravel(), z.ravel()], axis=1),
        np.stack([ear_x.ravel(), ear_y.ravel(), ear_z.ravel()], axis=1),
        np.stack([ear_x.ravel(), ear_y.ravel(), -ear_z.ravel()], axis=1)
    ])

    return points


def normalize_point_cloud(points):
    """点云归一化处理：中心化并缩放到单位球内"""
    centroid = np.mean(points, axis=0)
    points_centered = points - centroid
    max_norm = np.max(np.linalg.norm(points_centered, axis=1))
    points_normalized = points_centered / max_norm
    return points_normalized, centroid, max_norm


def cartesian_to_spherical(points):
    """笛卡尔坐标转换为球坐标"""
    x, y, z = points[:, 0], points[:, 1], points[:, 2]
    r = np.sqrt(x ** 2 + y ** 2 + z ** 2)
    theta = np.arccos(z / r)  # 天顶角 [0, pi]
    phi = np.arctan2(y, x)  # 方位角 [0, 2pi)
    return r, theta, phi


def compute_real_spherical_harmonics(theta, phi, l_max):
    """计算实值球谐基函数矩阵"""
    num_points = len(theta)
    num_coeffs = (l_max + 1) ** 2
    Y = np.zeros((num_points, num_coeffs))

    idx = 0
    for l in range(l_max + 1):
        for m in range(-l, l + 1):
            # 计算复数球谐函数
            Y_complex = sph_harm(abs(m), l, phi, theta)

            # 转换为实值球谐函数
            if m < 0:
                Y[:, idx] = np.sqrt(2) * (-1) ** m * Y_complex.imag
            elif m == 0:
                Y[:, idx] = Y_complex.real
            else:
                Y[:, idx] = np.sqrt(2) * (-1) ** m * Y_complex.real
            idx += 1

    return Y


def reconstruct_spherical_surface(l_max, theta_grid, phi_grid, coefficients):
    """使用球谐系数重建表面"""
    Y_grid = compute_real_spherical_harmonics(
        theta_grid.flatten(),
        phi_grid.flatten(),
        l_max
    )

    r_recon = Y_grid @ coefficients
    r_recon = r_recon.reshape(theta_grid.shape)

    # 转换为笛卡尔坐标
    sin_theta = np.sin(theta_grid)
    x = r_recon * sin_theta * np.cos(phi_grid)
    y = r_recon * sin_theta * np.sin(phi_grid)
    z = r_recon * np.cos(theta_grid)

    return x, y, z


def visualize_reconstruction(original, reconstructed):
    """可视化原始点云和重建结果（使用Matplotlib）"""
    fig = plt.figure(figsize=(14, 7))

    # 原始点云
    ax1 = fig.add_subplot(121, projection='3d')
    ax1.scatter(original[:, 0], original[:, 1], original[:, 2],
                s=1, c='r', alpha=0.6)
    ax1.set_title('Original Point Cloud')
    ax1.set_xlabel('X')
    ax1.set_ylabel('Y')
    ax1.set_zlabel('Z')
    ax1.view_init(elev=20, azim=-35)

    # 重建结果
    ax2 = fig.add_subplot(122, projection='3d')
    ax2.plot_surface(reconstructed[0], reconstructed[1], reconstructed[2],
                     rstride=1, cstride=1, cmap='viridis', alpha=0.8,
                     edgecolor='none', antialiased=True)
    ax2.set_title('Spherical Harmonic Reconstruction')
    ax2.set_xlabel('X')
    ax2.set_ylabel('Y')
    ax2.set_zlabel('Z')
    ax2.view_init(elev=20, azim=-35)

    plt.tight_layout()
    plt.show()


def main():
    # 参数设置
    L_MAX = 50  # 球谐函数最大阶数
    GRID_RES = 250  # 球面网格分辨率

    # 1. 加载并预处理点云
    points = load_bunny_point_cloud()
    points_norm, centroid, max_norm = normalize_point_cloud(points)

    # 2. 转换为球坐标
    r, theta, phi = cartesian_to_spherical(points_norm)

    # 3. 计算球谐基函数矩阵
    Y = compute_real_spherical_harmonics(theta, phi, L_MAX)

    # 4. 最小二乘拟合球谐系数
    coefficients = np.linalg.lstsq(Y, r, rcond=None)[0]

    # 5. 在球面网格上重建表面
    theta_grid, phi_grid = np.meshgrid(
        np.linspace(0, np.pi, GRID_RES),
        np.linspace(0, 2 * np.pi, GRID_RES),
        indexing='ij'
    )

    x_recon, y_recon, z_recon = reconstruct_spherical_surface(
        L_MAX, theta_grid, phi_grid, coefficients
    )

    # 6. 可视化结果
    visualize_reconstruction(points_norm, (x_recon, y_recon, z_recon))

    print("重建完成！使用球谐阶数:", L_MAX)


if __name__ == "__main__":
    main()