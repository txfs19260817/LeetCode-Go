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
    self.expression_map: dict[str, str] = {} # variable name to expression (e.g. "foo" -> "bar + 5 * 2")
    self.memo: dict[str, int] = {} # variable name to evaluated value
    self.visited: set[str] = set() # cycle detection

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
    n = len(expr)
    i = 0
    
    while i < n:
        char = expr[i]
        
        if char == ' ':
            i += 1
            continue
        
        # --- 解析由数字或变量构成的 "Operand" ---
        is_variable = False
        token_start = i
        if char.isdigit():
            # 读取完整数字
            while i < n and expr[i].isdigit():
                i += 1
            current_num = int(expr[token_start:i])
        elif char.isalpha():
            # 读取完整变量名
            is_variable = True
            while i < n and expr[i].isalnum(): # 假设变量名为 alphanumeric
                i += 1
            var_name = expr[token_start:i]
            # 关键点：递归获取变量的值
            current_num = self.evaluate(var_name)
        else:
            # 是操作符，稍后处理
            pass
        
        # --- 应用上一个操作符 (或者如果在字符串末尾) ---
        # 如果当前字符是操作符，或者已经到了末尾，我们需要处理之前的数字
        # 注意：这里的逻辑稍微有点绕，因为我们需要在读取到下一个 OP 时，处理上一个 OP 和当前的 NUM
        
        # 简化逻辑：每次循环我们都尝试解析一个数。
        # 如果上面解析到了数(或变量)，我们现在将其压栈。
        
        if is_variable or (i > token_start and expr[token_start].isdigit()):
            if last_op == '+':
                stack.append(current_num)
            elif last_op == '-':
                stack.append(-current_num)
            elif last_op == '*':
                top = stack.pop()
                stack.append(top * current_num)
            elif last_op == '/':
                top = stack.pop()
                # Python int division behaves differently for negative numbers, 
                # use int() truncation for C++ style behavior usually expected in interviews
                stack.append(int(top / current_num))
        
        # 检查是否是操作符
        if i < n and expr[i] in "+-*/":
            last_op = expr[i]
            i += 1
        
    return sum(stack)

# ==========================================
if __name__ == "__main__":
    compiler = MiniCompiler()
    
    input_lines = [
        "A = B + 10",
        "B = 5 * C",
        "C = 2",
        "D = A + B + C" 
    ]
    
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
        "y = 4 / 2" # y should be 2, so x should be 10 + 4 = 14
    ]
    result2 = compiler2.compile(lines2)
    print("Result 2:", result2)
    assert result2["x"] == 14
    assert result2["y"] == 2