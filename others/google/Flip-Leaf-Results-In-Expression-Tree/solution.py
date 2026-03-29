from __future__ import annotations

from dataclasses import dataclass
from typing import Dict, List, Optional


@dataclass(eq=False)
class Node:
    op: str
    val: Optional[bool] = None
    left: Optional["Node"] = None
    right: Optional["Node"] = None


class Solution:
    def flipLeafResults(self, root: Node) -> List[bool]:
        values: Dict[Node, bool] = {}

        def evaluate(node: Node) -> bool:
            if node.op == "leaf":
                if node.val is None:
                    raise ValueError("leaf nodes must have a boolean value")
                values[node] = node.val
            elif node.op == "not":
                if node.left is None:
                    raise ValueError("not nodes must have a left child")
                values[node] = not evaluate(node.left)
            else:
                if node.left is None or node.right is None:
                    raise ValueError(f"{node.op} nodes must have two children")
                left_value = evaluate(node.left)
                right_value = evaluate(node.right)
                if node.op == "and":
                    values[node] = left_value and right_value
                elif node.op == "or":
                    values[node] = left_value or right_value
                elif node.op == "xor":
                    values[node] = left_value ^ right_value
                else:
                    raise ValueError(f"unsupported operator: {node.op}")
            return values[node]

        root_value = evaluate(root)
        result: List[bool] = []

        def collect(node: Node, impact: bool) -> None:
            if node.op == "leaf":
                result.append(root_value ^ impact)
                return

            if node.op == "not":
                collect(node.left, impact)  # type: ignore[arg-type]
            elif node.op == "xor":
                collect(node.left, impact)  # type: ignore[arg-type]
                collect(node.right, impact)  # type: ignore[arg-type]
            elif node.op == "and":
                collect(node.left, impact and values[node.right])  # type: ignore[arg-type]
                collect(node.right, impact and values[node.left])  # type: ignore[arg-type]
            elif node.op == "or":
                collect(node.left, impact and (not values[node.right]))  # type: ignore[arg-type]
                collect(node.right, impact and (not values[node.left]))  # type: ignore[arg-type]
            else:
                raise ValueError(f"unsupported operator: {node.op}")

        collect(root, True)
        return result


if __name__ == "__main__":
    solver = Solution()

    root = Node(
        "or",
        left=Node(
            "and",
            left=Node("leaf", True),
            right=Node("leaf", False),
        ),
        right=Node(
            "not",
            left=Node("leaf", False),
        ),
    )
    assert solver.flipLeafResults(root) == [True, True, False]

    root = Node(
        "xor",
        left=Node("leaf", True),
        right=Node("leaf", False),
    )
    assert solver.flipLeafResults(root) == [False, False]

    root = Node(
        "and",
        left=Node("leaf", True),
        right=Node(
            "or",
            left=Node("leaf", False),
            right=Node("leaf", True),
        ),
    )
    assert solver.flipLeafResults(root) == [False, True, False]
