from typing import List

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        output = []
        prefixes = [1]
        postfixes = [1]
        for i, x in enumerate(nums[:-1]):
            prefixes.append(prefixes[i]*x)
        for i, x in enumerate(nums[::-1][:-1]):
            postfixes.append(postfixes[i]*x)
        for i in range(len(nums)):
            output.append(prefixes[i]*postfixes[-i-1])
        return output