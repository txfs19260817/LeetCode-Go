package _966_Recover_the_Original_Array

import "sort"

func recoverArray(nums []int) []int {
	sort.Ints(nums)     // lower[0] == sorted nums[0]
	for i := 1; ; i++ { // iterate and validate higher[0]
		k2 := nums[i] - nums[0]   // doubled k
		if k2 == 0 || k2%2 == 1 { // ensure k is a positive integer
			continue
		}
		higherVis := make([]bool, len(nums)) // record visited higher[]
		higherVis[i] = true
		ans := []int{(nums[0] + nums[i]) / 2}
		for l, r := 0, i+1; r < len(nums); r++ { // two pointers: l -> `lower`, r-> `higher`
			for l++; l < len(nums) && higherVis[l]; l++ { // prevent l from pointing to `higher` elements
			}
			for ; r < len(nums) && nums[r]-nums[l] < k2; r++ { // find next higher element
			}
			if r == len(nums) || nums[r]-nums[l] > k2 { // it should be nums[r]-nums[l] == 2k
				break
			}
			higherVis[r] = true // mark visited `higher` element
			ans = append(ans, (nums[l]+nums[r])/2)
		}
		if len(ans) == len(nums)/2 {
			return ans
		}
	}
}
