from __future__ import annotations


def max_equal_endpoint_subarray_sum(nums: list[int]) -> int:
    if not nums:
        return 0

    # 对于固定右端点 j:
    # sum(i..j) = prefix[j + 1] - prefix[i]
    # 如果 nums[i] == nums[j]，那我们只需要知道：
    # 之前所有值等于 nums[j] 的位置里，最小的 prefix[i] 是多少。
    min_prefix_before_value: dict[int, int] = {}

    prefix = 0
    answer = float("-inf")

    for value in nums:
        if value not in min_prefix_before_value:
            min_prefix_before_value[value] = prefix
        else:
            min_prefix_before_value[value] = min(min_prefix_before_value[value], prefix)

        prefix += value
        answer = max(answer, prefix - min_prefix_before_value[value])

    return int(answer)


if __name__ == "__main__":
    assert max_equal_endpoint_subarray_sum([1, -2, 3, 4, -1, 3]) == 9
    assert max_equal_endpoint_subarray_sum([5, -100, 5]) == 5
    assert max_equal_endpoint_subarray_sum([-3, -1, -3]) == -1
    assert max_equal_endpoint_subarray_sum([2, 1, 2, 1, 2]) == 8
    assert max_equal_endpoint_subarray_sum([]) == 0
