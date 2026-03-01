class Solution:
    def __init__(self, max_order: int = 50):
        self.max_order = max_order
        self.N = [1] * (max_order + 1)
        for i in range(2, max_order + 1):
            self.N[i] = 1 + self.N[i - 1] + self.N[i - 2]

    def _ensure_capacity(self, order: int) -> None:
        if order <= self.max_order:
            return
        old_max = self.max_order
        self.N.extend([1] * (order - old_max))
        for i in range(old_max + 1, order + 1):
            self.N[i] = 1 + self.N[i - 1] + self.N[i - 2]
        self.max_order = order

    def get_path_from_root(self, order: int, x: int) -> str:
        path: str = ""
        # `root_label` is the global preorder index of the current subtree root.  
        # I update it while descending left or right 
        # so I can compare `x` against subtree boundaries in the same global labeling system.
        root_label = 0

        # Global labels use preorder numbering.
        # For order=n:
        # - left subtree order is n-2, its root label is root_label + 1
        # - right subtree order is n-1, its root label is root_label + 1 + size(left)
        while order > 1 and root_label != x:
            left_size = self.N[order - 2]
            right_root = root_label + 1 + left_size

            if x < right_root:
                path += "L"
                root_label += 1
                order -= 2
            else:
                path += "R"
                root_label = right_root
                order -= 1

        return path

    def findPath(self, order: int, source: int, dest: int) -> str:
        self._ensure_capacity(order)
        a = self.get_path_from_root(order, source)
        b = self.get_path_from_root(order, dest)
        i = 0
        while i < min(len(a), len(b)) and a[i] == b[i]:
            i += 1
        up = "U" * (len(a) - i)
        down = b[i:]
        return up + down


if __name__ == "__main__":
    sol = Solution()

    # Example 1
    assert sol.findPath(5, 5, 7) == "UUURL"
    # Example 2
    assert sol.findPath(4, 8, 3) == "UUULR"
    # Example 3
    assert sol.findPath(5, 4, 13) == "UUURRRL"

    # source == dest
    assert sol.findPath(5, 3, 3) == ""
    assert sol.findPath(3, 0, 0) == ""

    # Root to leaf
    assert sol.findPath(5, 0, 14) == "RRRR"
    # Leaf to root
    assert sol.findPath(5, 14, 0) == "UUUU"

    # order=2 (3 nodes: root=0, left=1, right=2)
    assert sol.findPath(2, 0, 2) == "R"
    assert sol.findPath(2, 0, 1) == "L"
    assert sol.findPath(2, 1, 2) == "UR"
    assert sol.findPath(2, 2, 1) == "UL"

    print("All tests passed!")
