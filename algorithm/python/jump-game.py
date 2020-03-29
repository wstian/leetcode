# Given an array of non-negative integers, you are initially positioned at the first index of the array.
#
# Each element in the array represents your maximum jump length at that position.
#
# Determine if you are able to reach the last index.
#
# Example 1:
#
# Input: [2,3,1,1,4]
# Output: true
# Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
# Example 2:
#
# Input: [3,2,1,0,4]
# Output: false
# Explanation: You will always arrive at index 3 no matter what. Its maximum
#              jump length is 0, which makes it impossible to reach the last index.
from typing import List


class Solution0:
    def canJump(self, nums: List[int]) -> bool:
        # reachable_set = {len(nums) - 1}
        final = len(nums) - 1
        for index in range(len(nums) - 2, -1, -1):
            jump = nums[index]
            if index + jump >= final:
                final = index
        return final == 0


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        return self._canJump(nums=nums, index=0, best_final=-1) >= len(nums) - 1

    def _canJump(self, nums, index, best_final):
        best_final = max(best_final, index + nums[index])
        length = len(nums)
        for jump in range(nums[index], 0, -1):
            next_index = index + jump
            if next_index >= length:
                return next_index
            if next_index + nums[next_index] <= best_final:
                continue
            best_final = max(best_final, self._canJump(nums, next_index, best_final))
            if best_final >= length:
                return best_final
        return best_final


if __name__ == '__main__':
    cases = [
        (
            ([2, 3, 1, 1, 4],),
            True,
        ),
        (
            ([3, 2, 1, 0, 4],),
            False
        ),
        (
            ([2, 0],),
            True
        ),
        (
            ([0],),
            True
        ),
    ]
    from test_utils import run_test_cases
    run_test_cases(Solution, cases)
