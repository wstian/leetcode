class Solution:
    def strWithout3a3b(self, A: int, B: int) -> str:
        a = 'a'
        b = 'b'
        if A < B:
            A, B = B, A
            a, b = b, a
        slots = int(A / 2) + A % 2 - 1
        assert B >= slots, '%d should larger than %d' % (B, slots)
        ret = ''
        while B > 0:
            na = min(A, 2)
            nb = min(B, 2 if B > slots else 1)
            ret += a * na + b * nb
            A -= na
            B -= nb
            slots -= 1
        return ret + a * A


if __name__ == '__main__':
    cases = [
        (
            (1, 2),
            "bba",
        ),
        (
            (4, 1),
            "aabaa",
        ),
    ]
    from test_utils import run_test_cases
    run_test_cases(Solution, Solution.strWithout3a3b, cases)
