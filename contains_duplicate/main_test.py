# Execute: python3 -m unittest -v main_test.py

from main import Solution

import unittest

class TestSolution(unittest.TestCase):

    def test_1(self):
        s = Solution()
        input = [1,2,3,1]
        expected_output = True
        self.assertEqual(s.containsDuplicate(input), expected_output)

    def test_2(self):
        s = Solution()
        input = [1,2,3,4]
        expected_output = False
        self.assertEqual(s.containsDuplicate(input), expected_output)

    def test_3(self):
        s = Solution()
        expected_output = True
        input = [1,1,1,3,3,4,3,2,4,2]
        self.assertEqual(s.containsDuplicate(input), expected_output)

if __name__ == '__main__':
    unittest.main()