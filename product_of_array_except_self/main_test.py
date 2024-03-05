# Execute: python3 -m unittest -v main_test.py

from main import Solution

import unittest

class TestSolution(unittest.TestCase):

    def test_1(self):
        sol = Solution()
        input = [1,2,3,4]
        expected_output = [24,12,8,6]
        self.assertEqual(sol.productExceptSelf(input), expected_output)

    def test_2(self):
        sol = Solution()
        input = [-1,1,0,-3,3]
        expected_output = [0,0,9,0,0]
        self.assertEqual(sol.productExceptSelf(input), expected_output)

if __name__ == '__main__':
    unittest.main()