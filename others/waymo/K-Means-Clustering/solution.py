import random


def k_means(points: list[list[float]], k: int, max_iter: int) -> tuple[list[list[float]], list[int]]:
    centroids = random.sample(points, k)
    assignments = [0] * len(points)

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
    dim = len(points[0])
    next_centroids = [[0.0] * dim for _ in range(k)]
    counts = [0] * k

    for p, cluster in zip(points, assignments):
        for d in range(dim):
            next_centroids[cluster][d] += p[d]
        counts[cluster] += 1

    for i in range(k):
        if counts[i] == 0:
            next_centroids[i] = centroids[i][:]
            continue
        inv = 1.0 / counts[i]
        for d in range(dim):
            next_centroids[i][d] *= inv

    return next_centroids


def same_centroids(a: list[list[float]], b: list[list[float]]) -> bool:
    for i in range(len(a)):
        for d in range(len(a[i])):
            if abs(a[i][d] - b[i][d]) > 1e-9:
                return False
    return True


def distance_squared(a: list[float], b: list[float]) -> float:
    return sum((x - y) * (x - y) for x, y in zip(a, b))


if __name__ == "__main__":
    random.seed(42)

    points2d = [[1, 1], [1.5, 2], [3, 4], [5, 7], [3.5, 5], [4.5, 5], [3.5, 4.5]]
    centroids, assignments = k_means(points2d, 2, 10)
    assert len(centroids) == 2
    assert abs(centroids[0][0] - 1.25) < 1e-6
    assert abs(centroids[0][1] - 1.5) < 1e-6
    assert abs(centroids[1][0] - 3.9) < 1e-6
    assert abs(centroids[1][1] - 5.1) < 1e-6
    assert assignments == [0, 0, 1, 1, 1, 1, 1]

    points3d = [[0, 0, 0], [0, 1, 0], [9, 9, 9], [10, 9, 9]]
    centroids3d, assignments3d = k_means(points3d, 2, 10)
    assert assignments3d == [0, 0, 1, 1]
    assert abs(centroids3d[0][0] - 0.0) < 1e-6
    assert abs(centroids3d[0][1] - 0.5) < 1e-6
    assert abs(centroids3d[0][2] - 0.0) < 1e-6
    assert abs(centroids3d[1][0] - 9.5) < 1e-6
    assert abs(centroids3d[1][1] - 9.0) < 1e-6
    assert abs(centroids3d[1][2] - 9.0) < 1e-6
