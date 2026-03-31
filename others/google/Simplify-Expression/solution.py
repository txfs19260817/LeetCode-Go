class Solution:
    def simplifyExpression(self, expr: str) -> str:
        ans = []
        sign, stack = 1, [1]

        for ch in expr:
            if ch == "+":
                sign = 1
            elif ch == "-":
                sign = -1
            elif ch == "(":
                stack.append(stack[-1] * sign)
                sign = 1
            elif ch == ")":
                stack.pop()
            else:  # var
                effective = stack[-1] * sign
                if not ans:
                    if effective == -1:
                        ans.append('-')
                else:
                    ans.append('+' if effective == 1 else '-')
                ans.append(ch)
                sign = 1
        return "".join(ans)


if __name__ == "__main__":
    solver = Solution()

    assert solver.simplifyExpression("a-(b+c)") == "a-b-c"
    assert solver.simplifyExpression("a-(-b-c)") == "a+b+c"
    assert solver.simplifyExpression("(x+y)-z") == "x+y-z"

    assert solver.simplifyExpression("a+b-c") == "a+b-c"
    assert solver.simplifyExpression("-(a-b+c)") == "-a+b-c"
    assert solver.simplifyExpression("a-(b-(c-d))") == "a-b+c-d"
    assert solver.simplifyExpression("a-((b+c)-d)") == "a-b-c+d"
    assert solver.simplifyExpression("((a))") == "a"
