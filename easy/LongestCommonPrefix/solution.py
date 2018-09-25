class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if len(strs) == 0:
            return ""
        count = 0
        while True:
            flag = self.getChar(strs[0],count)
            for s in strs:
                if len(s) == 0 or len(s) == count or flag != self.getChar(s,count):
                    return strs[0][:count]
            count += 1
        return strs[0][:count]
                
    def getChar(self, s, i):
        if len(s) > i:
            return s[i]
        else:
            return ""
