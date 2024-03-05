from collections import defaultdict

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        counter = defaultdict(int)
        for x in s:
            counter[x] += 1
        for x in t:
            counter[x] -= 1

        for x in counter.values():
            if x != 0:
                return False

        return True