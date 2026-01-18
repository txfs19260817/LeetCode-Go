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
# 假设真实场景下这是调用外部硬件或API
# 为了演示，我们需要模拟一些“真实”存在的车
TRUE_CAR_LOCATIONS = [
    (10.5, 20.5), # car 1
    (10.1, 20.1), # car 2
    (30.0, 40.0)  # car 3
]

def mock_squaredar_scan(lat_min: float, lon_min: float, side_length: float) -> bool:
    """
    Returns True if any car is within the square defined by 
    [lat_min, lat_min + side_length) and [lon_min, lon_min + side_length).
    """
    lat_max, lon_max = lat_min + side_length, lon_min + side_length
    return any((lat_min <= car_lat < lat_max) and (lon_min <= car_lon < lon_max) for car_lat, car_lon in TRUE_CAR_LOCATIONS)

# ==========================================
# 2. Solution: Quadtree / Divide & Conquer
# ==========================================

class CarLocator:
    def __init__(self, resolution: float = 0.01):
        # resolution (EPSILON): 决定了定位的精度
        # 当正方形小于这个尺寸时，我们认为已经找到具体位置了
        self.resolution = resolution
        self.found_locations: List[Tuple[float, float]] = []

    def find_all_cars(self, lat_start: float, lon_start: float, size: float) -> List[Tuple[float, float]]:
        """
        Main entry point. 
        Start searching within a massive bounding box defined by start coords and size.
        """
        self.found_locations.clear() # Reset results
        self._recursive_scan(lat_start, lon_start, size)
        return self.found_locations

    def _recursive_scan(self, lat: float, lon: float, size: float) -> None:
        # 1. Check: Is there anything here?
        # 这里的关键优化在于：如果返回 False，直接剪枝，不再向下递归
        
        if not (has_car := mock_squaredar_scan(lat, lon, size)):
            return

        # 2. Base Case: The square is small enough to be considered a point
        half_size = size / 2.0
        if size <= self.resolution:
            # 记录中心点作为车的位置
            center_lat, center_lon = lat + half_size, lon + half_size
            self.found_locations.append((center_lat, center_lon)) # Add the center point as a car location
            return

        # 3. Recursive Step: Split into 4 quadrants
        
        # 为了清晰，定义四个子区域的坐标
        # Quadrant 1: Bottom-Left
        self._recursive_scan(lat, lon, half_size)
        
        # Quadrant 2: Bottom-Right
        self._recursive_scan(lat, lon + half_size, half_size)
        
        # Quadrant 3: Top-Left
        self._recursive_scan(lat + half_size, lon, half_size)
        
        # Quadrant 4: Top-Right
        self._recursive_scan(lat + half_size, lon + half_size, half_size)

# ==========================================
# 3. Execution
# ==========================================

if __name__ == "__main__":
    locator = CarLocator(resolution=0.1) # 精度设为 0.1
    
    # 假设我们搜索一个 100x100 的大区域
    start_lat, start_lon = 0.0, 0.0
    area_size = 100.0
    
    print(f"Scanning area {area_size}x{area_size}...")
    cars = locator.find_all_cars(start_lat, start_lon, area_size)
    
    print(f"Found {len(cars)} locations:")
    for car in cars:
        print(f" - Approximate Location: {car}")