class _Node:
    __slots__ = ("key", "height", "left", "right")

    def __init__(self, key: int) -> None:
        self.key = key
        self.height = 1
        self.left: "_Node | None" = None
        self.right: "_Node | None" = None


class AVLTree:
    def __init__(self) -> None:
        self._root: _Node | None = None

    def insert(self, key: int) -> None:
        self._root = self._insert(self._root, key)

    def delete(self, key: int) -> None:
        self._root = self._delete(self._root, key)

    def contains(self, key: int) -> bool:
        cur = self._root
        while cur is not None:
            if key < cur.key:
                cur = cur.left
            elif key > cur.key:
                cur = cur.right
            else:
                return True
        return False

    def inorder(self) -> list[int]:
        out: list[int] = []

        def walk(node: _Node | None) -> None:
            if node is None:
                return
            walk(node.left)
            out.append(node.key)
            walk(node.right)

        walk(self._root)
        return out

    def _insert(self, node: _Node | None, key: int) -> _Node:
        if node is None:
            return _Node(key)
        if key < node.key:
            node.left = self._insert(node.left, key)
        elif key > node.key:
            node.right = self._insert(node.right, key)
        else:
            return node
        return self._rebalance(node)

    def _delete(self, node: _Node | None, key: int) -> _Node | None:
        if node is None:
            return None
        if key < node.key:
            node.left = self._delete(node.left, key)
            return self._rebalance(node)
        if key > node.key:
            node.right = self._delete(node.right, key)
            return self._rebalance(node)

        if node.left is None:
            return node.right
        if node.right is None:
            return node.left

        successor = self._min_node(node.right)
        node.key = successor.key
        node.right = self._delete(node.right, successor.key)
        return self._rebalance(node)

    def _rebalance(self, node: _Node) -> _Node:
        self._update_height(node)
        balance = self._height(node.left) - self._height(node.right)

        if balance > 1:
            assert node.left is not None
            if self._height(node.left.left) < self._height(node.left.right):
                node.left = self._rotate_left(node.left)
            return self._rotate_right(node)
        if balance < -1:
            assert node.right is not None
            if self._height(node.right.right) < self._height(node.right.left):
                node.right = self._rotate_right(node.right)
            return self._rotate_left(node)
        return node

    @staticmethod
    def _min_node(node: _Node) -> _Node:
        cur = node
        while cur.left is not None:
            cur = cur.left
        return cur

    def _rotate_left(self, x: _Node) -> _Node:
        y = x.right
        assert y is not None
        t2 = y.left
        y.left = x
        x.right = t2
        self._update_height(x)
        self._update_height(y)
        return y

    def _rotate_right(self, y: _Node) -> _Node:
        x = y.left
        assert x is not None
        t2 = x.right
        x.right = y
        y.left = t2
        self._update_height(y)
        self._update_height(x)
        return x

    def _update_height(self, node: _Node) -> None:
        node.height = max(self._height(node.left), self._height(node.right)) + 1

    @staticmethod
    def _height(node: _Node | None) -> int:
        return 0 if node is None else node.height


if __name__ == "__main__":
    tree = AVLTree()
    for key in [10, 20, 30, 40, 50, 25]:
        tree.insert(key)
    assert tree.inorder() == [10, 20, 25, 30, 40, 50]
    tree.delete(40)
    assert tree.contains(40) is False
    assert tree.contains(25) is True
    assert tree.inorder() == [10, 20, 25, 30, 50]

    tree2 = AVLTree()
    for i in range(1, 101):
        tree2.insert(i)
    assert tree2.inorder() == list(range(1, 101))
    for i in range(1, 101, 2):
        tree2.delete(i)
    assert tree2.inorder() == list(range(2, 101, 2))

    tree3 = AVLTree()
    for key in [5, 5, 5, 3, 7, 7]:
        tree3.insert(key)
    assert tree3.inorder() == [3, 5, 7]
