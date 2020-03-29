# Alex and Lee continue their games with piles of stones.  There are a number of piles arranged
# in a row, and each pile has a positive integer number of stones piles[i].
# The objective of the game is to end with the most stones.
#
# Alex and Lee take turns, with Alex starting first.  Initially, M = 1.
#
# On each player's turn, that player can take all the stones in the first X remaining piles,
# where 1 <= X <= 2M.  Then, we set M = max(M, X).
#
# The game continues until all the stones have been taken.
#
# Assuming Alex and Lee play optimally, return the maximum number of stones Alex can get.
#
#
# Example 1:

# Input: piles = [2,7,9,4,4]
# Output: 10
# Explanation:  If Alex takes one pile at the beginning, Lee takes two piles, then Alex takes 2
# piles again. Alex can get 2 + 4 + 4 = 10 piles in total. If Alex takes two piles at the beginning,
# then Lee can take all three piles left. In this case, Alex get 2 + 7 = 9 piles in total.
# So we return 10 since it's larger
from typing import List


class Solution:
    def stoneGameII(self, piles: List[int]) -> int:
        return self._stoneGameII(piles, start=0, M=1, cached={})

    def _stoneGameII(self, piles: List[int], start: int, M: int, cached) -> int:
        if start >= len(piles):
            return 0
        if (start, M) in cached:
            return cached[(start, M)]
        remained_sum = sum(piles[start:])
        ret = max(
            remained_sum - self._stoneGameII(piles, start=start + x, M=max(M, x), cached=cached)
            for x in range(1, min(2 * M + 1, len(piles) - start + 1))
        )
        cached[(start, M)] = ret
        return ret


if __name__ == '__main__':
    cases = [
        (
            ([2, 7, 9, 4, 4],),
            10,
        ),
        (
            ([1, 2, 3, 4, 5, 100],),
            104,
        ),
    ]
    from test_utils import run_test_cases
    run_test_cases(Solution, cases)
