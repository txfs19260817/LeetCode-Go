"""
The goal is to check if any of the car is in a sqaure or not to a given car
Returns true if any of the vehicles are within the queried square region.

def mock_squaredar_scan(lat_min: float, lon_min: float, side_length: float) → bool:

imagine we have an interface that we can get the lat min, lon min, side_length...
come up with an algorithm to call this mock_squaredar_scan function to give the locations of all
"""

from typing import List, Tuple

# ==========================================
# 1. Mock Interface provided by the question
# ==========================================
TRUE_CAR_LOCATIONS = [
    (10.1, 20.1),  # car 1
    (10.5, 20.5),  # car 2
    (30.0, 40.0),  # car 3
]


def mock_squaredar_scan(lat_min: float, lon_min: float, side_length: float) -> bool:
    """
    Returns True if any car is within the square defined by
    [lat_min, lat_min + side_length) and [lon_min, lon_min + side_length).
    """
    lat_max, lon_max = lat_min + side_length, lon_min + side_length
    return any(
        (lat_min <= car_lat < lat_max) and (lon_min <= car_lon < lon_max)
        for car_lat, car_lon in TRUE_CAR_LOCATIONS
    )


# ==========================================
# 2. Solution: Quadtree / Divide & Conquer
# ==========================================


def find_all_cars(
    lat_start: float, lon_start: float, size: float, eps: float = 0.001
) -> List[Tuple[float, float]]:
    output: List[Tuple[float, float]] = []
    stack: List[Tuple[float, float, float]] = [(lat_start, lon_start, size)] # (lat, lon, size)
    while stack:
        lat, lon, size = stack.pop()
        if not (has_car := mock_squaredar_scan(lat, lon, size)):
            continue
        if size <= eps:
            output.append((lat, lon))
            continue
        half_size = size / 2.0
        stack.append((lat,             lon,             half_size))
        stack.append((lat + half_size, lon,             half_size))
        stack.append((lat,             lon + half_size, half_size))
        stack.append((lat + half_size, lon + half_size, half_size))
    return output


# ==========================================
# 3. Execution
# ==========================================

if __name__ == "__main__":
    start_lat, start_lon = 0.0, 0.0
    area_size = 100.0
    eps = 0.001

    print(f"Scanning area {area_size}x{area_size}...")
    cars = find_all_cars(start_lat, start_lon, area_size, eps)

    print(f"Found {len(cars)} locations:")
    for i, car in enumerate(sorted(cars, key=lambda x: (x[0], x[1]))):
        print(f" - Approximate Location: ({car[0]:.2f}, {car[1]:.2f}), actual location: {TRUE_CAR_LOCATIONS[i]}, delta: {abs(car[0] - TRUE_CAR_LOCATIONS[i][0]):.2f}, {abs(car[1] - TRUE_CAR_LOCATIONS[i][1]):.2f}")
        assert abs(car[0] - TRUE_CAR_LOCATIONS[i][0]) < eps
        assert abs(car[1] - TRUE_CAR_LOCATIONS[i][1]) < eps
