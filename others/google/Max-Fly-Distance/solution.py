class Solution:
    def maxFlyDistance(self, arr: list[int], k: int) -> int:
        n = len(arr)
        neg_inf = float("-inf")
        dp = [neg_inf] * (k + 1)
        dp[k] = 0

        for i in range(n):
            next_dp = [neg_inf] * (k + 1)
            for energy in range(k + 1):
                if dp[energy] == neg_inf:
                    continue

                rest_energy = min(k, energy + 1)
                next_dp[rest_energy] = max(next_dp[rest_energy], dp[energy])

                if energy > 0:
                    next_dp[energy - 1] = max(
                        next_dp[energy - 1],
                        dp[energy] + arr[i],
                    )
            dp = next_dp

        return int(max(dp))


if __name__ == "__main__":
    solver = Solution()

    assert solver.maxFlyDistance([5, 2, 8, 4, 3], 2) == 17
    assert solver.maxFlyDistance([1, 2, 3, 4, 5], 3) == 13
    assert solver.maxFlyDistance([10, 1, 1, 10], 2) == 21
    assert solver.maxFlyDistance([7], 1) == 7
