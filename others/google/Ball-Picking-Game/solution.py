class Solution:
    def maxDifference(self, nums: list[int]) -> int:
        n = len(nums)
        # dp[i] = maximum score difference the current player can guarantee
        # over the opponent starting from the suffix nums[i:].
        dp = [0] * (n + 2)

        for i in range(n - 1, -1, -1):
            # If we take one ball, we gain nums[i], then the opponent gets the
            # advantage represented by dp[i + 1], so the net difference is:
            take_one = nums[i] - dp[i + 1]

            # If we take two balls, we gain both values, then the opponent
            # plays optimally from i + 2.
            take_two = take_one
            if i + 1 < n:
                take_two = nums[i] + nums[i + 1] - dp[i + 2]

            # Choose the move that gives the larger final difference.
            dp[i] = max(take_one, take_two)

        return dp[0]


if __name__ == "__main__":
    solver = Solution()
    assert solver.maxDifference([1, -1, -3, 1, 2, 4]) == 2
    assert solver.maxDifference([1, 2, 3, 4]) == 0
    assert solver.maxDifference([4, -1, 2, -3, 5]) == 5
    assert solver.maxDifference([7]) == 7
