import os
import pprint
import sys
import textwrap
from typing import Any, Callable, ClassVar, List, Tuple

InputArgs = Tuple[Any]
ExpectedReturn = Any
TestCase = Tuple[InputArgs, ExpectedReturn]


def run_test_cases(solution_class: ClassVar,
                   cases: 'List[TestCase]',
                   method: Callable[[Any], Any] = None,
                   verbose: bool = None):
    if verbose is None:
        verbose = 'LEETCODE_PYTHON_TEST_VERBOSE' in os.environ
    if method is None:
        name = [key for key in dir(solution_class) if not str.startswith(key, '_')]
        assert name and len(name) == 1, \
            'should have one and only one public method of %s' % solution_class.__name__
        method = getattr(solution_class, name[0])
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
