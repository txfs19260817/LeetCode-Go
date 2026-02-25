from typing import List


def min_partition_difference(nums: List[int]) -> int:
    if not nums:
        raise ValueError("nums must be non-empty")

    total = sum(nums)
    target = total // 2

    # dp[s] = whether subset sum s is achievable
    dp = [False] * (target + 1)
    dp[0] = True

    for num in nums:
        # Reverse traversal keeps each number used at most once (0/1 knapsack).
        for s in range(target, num - 1, -1):
            if dp[s - num]:
                dp[s] = True

    best = 0
    # Closest achievable sum to total/2 gives minimum partition difference.
    for s in range(target, -1, -1):
        if dp[s]:
            best = s
            break

    return total - 2 * best


if __name__ == "__main__":
    assert min_partition_difference([2, 3, 10, 7, 5]) == 1
    assert min_partition_difference([1, 6, 11, 5]) == 1
    assert min_partition_difference([1, 2, 3, 9]) == 3
    assert min_partition_difference([5]) == 5

    try:
        min_partition_difference([])
        raise AssertionError("Expected ValueError for empty input")
    except ValueError:
        pass

    print("All assertions passed.")
