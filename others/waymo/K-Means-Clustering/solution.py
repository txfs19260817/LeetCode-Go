import random


def k_means(points: list[list[float]], k: int, max_iter: int) -> tuple[list[list[float]], list[int]]:
    # 随机选 k 个不同点作为初始质心（浅拷贝避免后续共享引用）。
    centroids = [p[:] for p in random.sample(points, k)]
    assignments = [0] * len(points)

    # 经典 Lloyd 迭代：分配 -> 更新；质心不再变化则提前收敛。
    for _ in range(max_iter):
        assignments = assign(points, centroids)
        next_centroids = recompute(points, centroids, assignments, k)
        if same_centroids(centroids, next_centroids):
            centroids = next_centroids
            break
        centroids = next_centroids

    assignments = assign(points, centroids)
    return centroids, assignments


def assign(points: list[list[float]], centroids: list[list[float]]) -> list[int]:
    # 为每个点选最近质心；若距离相等，保留更小索引（由严格 < 控制）。
    out: list[int] = []
    for p in points:
        best_idx = 0
        best_dist = distance_squared(p, centroids[0])
        for i in range(1, len(centroids)):
            d = distance_squared(p, centroids[i])
            if d < best_dist:
                best_dist = d
                best_idx = i
        out.append(best_idx)
    return out


def recompute(
    points: list[list[float]],
    centroids: list[list[float]],
    assignments: list[int],
    k: int,
) -> list[list[float]]:
    # 按簇累加向量和，最后除以计数得到均值质心。
    dim = len(points[0])
    next_centroids = [[0.0] * dim for _ in range(k)]
    counts = [0] * k

    for p, cluster in zip(points, assignments):
        for d in range(dim):
            next_centroids[cluster][d] += p[d]
        counts[cluster] += 1

    for i in range(k):
        # 空簇保留旧质心，避免除零并保持迭代稳定。
        if counts[i] == 0:
            next_centroids[i] = centroids[i][:]
            continue
        inv = 1.0 / counts[i]
        for d in range(dim):
            next_centroids[i][d] *= inv

    return next_centroids


def same_centroids(a: list[list[float]], b: list[list[float]]) -> bool:
    # 用小阈值比较浮点向量是否“相等”。
    for i in range(len(a)):
        for d in range(len(a[i])):
            if abs(a[i][d] - b[i][d]) > 1e-9:
                return False
    return True


def distance_squared(a: list[float], b: list[float]) -> float:
    # 比较大小时不需要开方，平方距离更快且等价。
    return sum((x - y) * (x - y) for x, y in zip(a, b))


def assert_valid_solution(
    points: list[list[float]], centroids: list[list[float]], assignments: list[int], k: int
) -> None:
    # 基础结构约束：数量与索引范围。
    assert len(centroids) == k
    assert len(assignments) == len(points)

    for a in assignments:
        assert 0 <= a < k

    # Every assignment must be the nearest-centroid choice (tie -> smaller index).
    for p, assigned in zip(points, assignments):
        best_idx = 0
        best_dist = distance_squared(p, centroids[0])
        for i in range(1, k):
            d = distance_squared(p, centroids[i])
            if d < best_dist:
                best_dist = d
                best_idx = i
        assert assigned == best_idx


def close(a: float, b: float, eps: float = 1e-3) -> bool:
    return abs(a - b) <= eps


if __name__ == "__main__":
    # 固定随机种子，保证“随机初始化 + 测试”仍可稳定复现。
    random.seed(42)

    points2d = [[1, 1], [1.5, 2], [3, 4], [5, 7], [3.5, 5], [4.5, 5], [3.5, 4.5]]
    expected2d = [[1.25, 1.5], [3.9, 5.1]]
    for _ in range(20):
        centroids, assignments = k_means(points2d, 2, 30)
        assert_valid_solution(points2d, centroids, assignments, 2)
        got = sorted(centroids)
        exp = sorted(expected2d)
        assert close(got[0][0], exp[0][0])
        assert close(got[0][1], exp[0][1])
        assert close(got[1][0], exp[1][0])
        assert close(got[1][1], exp[1][1])

    # 再测一组 3D，确认实现与维度无关。
    points3d = [[0, 0, 0], [0, 1, 0], [9, 9, 9], [10, 9, 9]]
    expected3d = [[0.0, 0.5, 0.0], [9.5, 9.0, 9.0]]
    for _ in range(20):
        centroids3d, assignments3d = k_means(points3d, 2, 30)
        assert_valid_solution(points3d, centroids3d, assignments3d, 2)
        got3d = sorted(centroids3d)
        exp3d = sorted(expected3d)
        for i in range(2):
            for d in range(3):
                assert close(got3d[i][d], exp3d[i][d])
