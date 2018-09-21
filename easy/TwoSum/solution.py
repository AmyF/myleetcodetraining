class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        leftindex = 0
        rightindex = 1
        while True:
            for i in range(len(nums) - 1 - leftindex):
                if target == nums[leftindex] + nums[rightindex]:
                    return leftindex,rightindex
                rightindex += 1
            leftindex += 1
            rightindex = leftindex + 1
        return leftindex,rightindex
            
