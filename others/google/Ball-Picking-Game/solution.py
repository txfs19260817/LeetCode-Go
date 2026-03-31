from functools import lru_cache
from time import perf_counter


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

    def maxDifferenceDfsMemo(self, nums: list[int]) -> int:
        n = len(nums)

        @lru_cache(maxsize=None)
        def dfs(i: int) -> int:
            if i >= n:
                return 0

            # Take one ball now, then the opponent gets dfs(i + 1).
            best = nums[i] - dfs(i + 1)

            # Take two balls now if possible.
            if i + 1 < n:
                best = max(best, nums[i] + nums[i + 1] - dfs(i + 2))

            return best

        return dfs(0)


if __name__ == "__main__":
    solver = Solution()
    assert solver.maxDifference([1, -1, -3, 1, 2, 4]) == 2
    assert solver.maxDifferenceDfsMemo([1, -1, -3, 1, 2, 4]) == 2
    assert solver.maxDifference([1, 2, 3, 4]) == 0
    assert solver.maxDifferenceDfsMemo([1, 2, 3, 4]) == 0
    assert solver.maxDifference([4, -1, 2, -3, 5]) == 5
    assert solver.maxDifferenceDfsMemo([4, -1, 2, -3, 5]) == 5
    assert solver.maxDifference([7]) == 7
    assert solver.maxDifferenceDfsMemo([7]) == 7

    benchmark_nums = [((i * 37) % 21) - 10 for i in range(500)]
    rounds = 200

    start = perf_counter()
    for _ in range(rounds):
        answer_dp = solver.maxDifference(benchmark_nums)
    dp_seconds = perf_counter() - start

    start = perf_counter()
    for _ in range(rounds):
        answer_dfs = solver.maxDifferenceDfsMemo(benchmark_nums)
    dfs_seconds = perf_counter() - start

    assert answer_dp == answer_dfs
    print(f"benchmark n={len(benchmark_nums)}, rounds={rounds}")
    print(f"bottom-up dp: {dp_seconds:.6f}s")
    print(f"dfs + memo:   {dfs_seconds:.6f}s")
