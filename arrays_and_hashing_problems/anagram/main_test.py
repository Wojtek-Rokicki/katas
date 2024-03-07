# Execute: python3 -m unittest -v main_test.py

from main import Solution

import unittest

class TestSolution(unittest.TestCase):

    def test_1(self):
        sol = Solution()
        s = "anagram"
        t = "nagaram"
        expected_output = True
        self.assertEqual(sol.isAnagram(s, t), expected_output)

    def test_2(self):
        sol = Solution()
        s = "rat"
        t = "car"
        expected_output = False
        self.assertEqual(sol.isAnagram(s, t), expected_output)

if __name__ == '__main__':
    unittest.main()