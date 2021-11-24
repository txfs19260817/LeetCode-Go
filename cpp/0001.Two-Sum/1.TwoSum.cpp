#include <algorithm>
#include <bitset>
#include <cmath>
#include <iostream>
#include <list>
#include <map>
#include <queue>
#include <set>
#include <stack>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <gtest/gtest.h>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int> &nums, int target) {
        unordered_map<int, int> m;
        for (int i = 0; i < nums.size(); ++i) {
            int numJ = target - nums[i];
            if (m.find(numJ) != m.end()) {
                return {m[numJ], i};
            }
            m[nums[i]] = i;
        }
        return {};
    }
};

int main() {
    testing::InitGoogleTest();
    return RUN_ALL_TESTS();
}

TEST(LeetCodeTest, twoSum0) {
    vector<int> num = {2, 7, 11, 15};
    vector<int> answer = Solution().twoSum(num, 9);
    vector<int> expected = {0, 1};
    sort(answer.begin(), answer.end());
    sort(expected.begin(), expected.end());
    EXPECT_TRUE(answer == expected);
}