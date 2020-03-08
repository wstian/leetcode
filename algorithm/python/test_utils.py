import os
import pprint
import sys
import textwrap
from typing import Any, Callable, ClassVar, List, Tuple

InputArgs = Tuple
ExpectedReturn = Any
TestCase = Tuple[InputArgs, ExpectedReturn]


def run_test_cases(solution_class: ClassVar,
                   method: Callable[[Any], Any],
                   cases: 'List[TestCase]', verbose: bool = None):
    if verbose is None:
        verbose = 'LEETCODE_PYTHON_TEST_VERBOSE' in os.environ
    for i, case in enumerate(cases, start=1):
        args, expected = case
        got = method(solution_class(), *args)
        if verbose or got != expected:
            print('=== case[%d]: %s ===' % (i, 'pass' if got == expected else 'fail'))
            d = {'expected': expected, 'got': got, 'input': args}
            print(textwrap.indent(pprint.pformat(d, indent=2), prefix='  '))
            print('')
        if got != expected:
            print('*** FAIL ***')
            sys.exit(1)
    print('*** PASS ***')
    return sys.exit(0)
