class Solution:
    def findPath(self, order: int, source: int, dest: int) -> str:
        if source == dest:
            return ""

        # Precompute subtree sizes once so each lookup inside path_from_root is O(1).
        size = [1] * (order + 1)
        for n in range(2, order + 1):
            size[n] = 1 + size[n - 1] + size[n - 2]

        def path_from_root(cur_order: int, root_label: int, target: int) -> str:
            """
            A helper that returns the L/R path from a subtree root to a target label.
            Under **preorder** labeling, the left child root is always root_label + 1,
            and the right child root starts at root_label + 1 + left_size,
            where left_size is the number of nodes in the left subtree.

            :param order: order describes that subtree’s structure
            :param root_label: root_label is always the root label of the current subtree
            :param target:
            :return:
            """
            path = []
            while cur_order > 1 and root_label != target:
                left_size = size[cur_order - 2]  # left order = root order - 2
                right_start = root_label + 1 + left_size  # root node, skip root, skip left subtree

                # If our target is strictly less than right_start,
                # it must be in the left subtree.
                # I append 'L', update the root to the left child, and drop the order by 2.
                # Otherwise, it's in the right subtree:
                # I append 'R', jump the root label straight to right_start,
                # and drop the order by 1.
                if target < right_start:
                    path.append("L")
                    root_label += 1
                    cur_order -= 2
                else:
                    path.append("R")
                    root_label = right_start
                    cur_order -= 1
            return "".join(path)

        path_to_source = path_from_root(order, 0, source)
        path_to_dest = path_from_root(order, 0, dest)

        # Longest common prefix = LCA depth
        # generate the root-to-node paths for both the source and the destination.
        # I'll iterate through both strings to find the length of their longest common prefix.
        # This prefix represents the exact path to their Lowest Common Ancestor.
        common = 0
        while (
            common < len(path_to_source)
            and common < len(path_to_dest)
            and path_to_source[common] == path_to_dest[common]
        ):
            common += 1

        # To construct the final answer, we need to walk 'Up' from the source to the LCA.
        # I do this by adding a 'U' for every remaining character in the source path.
        # Then, I just append the remaining characters of the destination path to walk down to our target
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
