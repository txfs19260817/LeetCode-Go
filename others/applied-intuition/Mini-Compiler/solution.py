"""
Input: A list of strings, where each string is an assignment expression.
       Format: "var_name = expression"
       Expression contains integers, variables, +, -, *, /.

Output: A dictionary mapping all variables to their evaluated integer values.

Example:
Input:
[
  "foo = bar + 5",
  "bar = 2 * 3",
  "baz = foo + bar"
]

Output:
{
  "bar": 6,
  "foo": 11,   # 6 + 5
  "baz": 17    # 11 + 6
}
"""


class MiniCompiler:
    def __init__(self):
        self.expression_map: dict[str, str] = (
            {}
        )  # variable name to expression (e.g. "foo" -> "bar + 5 * 2")
        self.memo: dict[str, int] = {}  # variable name to evaluated value
        self.visited: set[str] = set()  # cycle detection

    def compile(self, inputs: list[str]) -> dict[str, int]:
        for line in inputs:
            var_name, expression = line.split("=")
            self.expression_map[var_name.strip()] = expression.strip()

        for var_name in self.expression_map.keys():
            self.evaluate(var_name)
        return self.memo

    def evaluate(self, var_name: str) -> int:
        if var_name in self.memo:
            return self.memo[var_name]
        if var_name in self.visited:
            raise ValueError(f"Cycle detected: {var_name}")
        self.visited.add(var_name)
        expression = self.expression_map[var_name]
        result = self.calculate_expression(expression)
        self.memo[var_name] = result
        self.visited.remove(var_name)
        return result

    def calculate_expression(self, expr: str) -> int:
        stack, current_num, last_op = [], 0, "+"
        expr = expr.replace(" ", "")
        i, n = 0, len(expr)
        while i < n:
            char = expr[i]
            if char.isdigit():
                num = 0
                while i < n and expr[i].isdigit():
                    num = num * 10 + int(expr[i])
                    i += 1
                current_num = num
            elif char.isalpha():
                var_name = ""
                while i < n and expr[i].isalpha():
                    var_name += expr[i]
                    i += 1
                current_num = self.evaluate(var_name)
            if i == n or expr[i] in "+-*/":
                if last_op == "+":
                    stack.append(current_num)
                elif last_op == "-":
                    stack.append(-current_num)
                elif last_op == "*":
                    stack.append(stack.pop() * current_num)
                elif last_op == "/":
                    stack.append(int(stack.pop() / current_num))
                current_num, last_op = 0, expr[i] if i < n else "+"
                i += 1
        return sum(stack)


# ==========================================
if __name__ == "__main__":
    compiler = MiniCompiler()

    input_lines = ["A = B + 10", "B = 5 * C", "C = 2", "D = A + B + C"]

    # Logic Trace:
    # C = 2
    # B = 5 * 2 = 10
    # A = 10 + 10 = 20
    # D = 20 + 10 + 2 = 32

    result = compiler.compile(input_lines)
    print("Result:", result)
    # Expected: {'A': 20, 'B': 10, 'C': 2, 'D': 32}
    assert result["A"] == 20
    assert result["B"] == 10
    assert result["C"] == 2
    assert result["D"] == 32

    # Test with division and precedence
    compiler2 = MiniCompiler()
    lines2 = [
        "x = 10 + 2 * y",
        "y = 4 / 2",  # y should be 2, so x should be 10 + 4 = 14
    ]
    result2 = compiler2.compile(lines2)
    print("Result 2:", result2)
    assert result2["x"] == 14
    assert result2["y"] == 2

    # More complex: multiple ops, multi-digit constants, deep dependencies
    compiler3 = MiniCompiler()
    lines3 = [
        "p = 12 + 3 * 4 - 6 / 2",  # 12 + 12 - 3 = 21
        "q = p * 2 + 7",
        "r = q / 5 + p * 3",
        "s = r + q - p / 2",
    ]
    result3 = compiler3.compile(lines3)
    print("Result 3:", result3)
    assert result3["p"] == 21
    assert result3["q"] == 49
    assert result3["r"] == 72
    assert result3["s"] == 111

    # Negative results via subtraction and integer truncation
    compiler4 = MiniCompiler()
    lines4 = [
        "a = 5 - 10 * 3",  # 5 - 30 = -25
        "b = a / 4",  # int(-25 / 4) = -6
        "c = b * 7 + 1",  # -42 + 1 = -41
    ]
    result4 = compiler4.compile(lines4)
    print("Result 4:", result4)
    assert result4["a"] == -25
    assert result4["b"] == -6
    assert result4["c"] == -41

    # Cycle detection
    compiler5 = MiniCompiler()
    lines5 = [
        "m = n + 1",
        "n = m + 2",
    ]
    try:
        compiler5.compile(lines5)
        raise AssertionError("Expected cycle detection to raise ValueError")
    except ValueError:
        print("Cycle detected as expected.")
