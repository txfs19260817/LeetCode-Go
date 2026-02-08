class Solution:
    def findPath(self, order: int, source: int, dest: int) -> str:
        if source == dest:
            return ""

        def fib_nodes(n: int) -> int:
            if n <= 1:
                return 1
            a, b = 1, 1
            for _ in range(2, n + 1):
                a, b = b, 1 + a + b
            return b

        def path_from_root(order: int, root_label: int, target: int) -> str:
            path = []
            while order > 1 and root_label != target:
                left_size = fib_nodes(order - 2)
                right_start = root_label + 1 + left_size
                if target < right_start:
                    path.append("L")
                    root_label += 1
                    order -= 2
                else:
                    path.append("R")
                    root_label = right_start
                    order -= 1
            return "".join(path)

        path_to_source = path_from_root(order, 0, source)
        path_to_dest = path_from_root(order, 0, dest)

        # Longest common prefix = LCA depth
        common = 0
        while (
            common < len(path_to_source)
            and common < len(path_to_dest)
            and path_to_source[common] == path_to_dest[common]
        ):
            common += 1

        return "U" * (len(path_to_source) - common) + path_to_dest[common:]


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
