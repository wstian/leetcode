from typing import List


class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        seen = set()

        def area(_r, _c):
            if not(0 <= _r < len(grid) and 0 <= _c < len(grid[_r]) and
                           (_r, _c) not in seen and grid[_r][_c]):
                return 0
            seen.add((_r, _c))
            return 1 + area(_r-1, _c) + area(_r+1, _c) + area(_r, _c-1) + area(_r, _c+1)
        return max(area(r, c) for r in range(len(grid)) for c in range(len(grid[r])))


if __name__ == '__main__':
    cases = [
        (
            (
                [[0,0,1,0,0,0,0,1,0,0,0,0,0],
                 [0,0,0,0,0,0,0,1,1,1,0,0,0],
                 [0,1,1,0,1,0,0,0,0,0,0,0,0],
                 [0,1,0,0,1,1,0,0,1,0,1,0,0],
                 [0,1,0,0,1,1,0,0,1,1,1,0,0],
                 [0,0,0,0,0,0,0,0,0,0,1,0,0],
                 [0,0,0,0,0,0,0,1,1,1,0,0,0],
                 [0,0,0,0,0,0,0,1,1,0,0,0,0]],
            ),
            6,
        ),
        (
            (
                [[0,0,0,0,0,0,0,0]],
            ),
            0,
        ),
    ]
    from test_utils import run_test_cases
    run_test_cases(Solution, Solution.maxAreaOfIsland, cases)
