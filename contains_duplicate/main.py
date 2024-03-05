from typing import List

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        unique_dict = set()
        for x in nums:
            if x in unique_dict:
                return True
            unique_dict.add(x)
        
        return False