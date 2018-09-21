class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        m = {}
        i,count,ret = 0,0,0
        while i < len(s):
            tmp = s[i]
            if m.get(tmp,-1) == -1:
                m[tmp] = 1
                count += 1
                i += 1
            else:
                ret = max(count,ret)
                m.pop(s[i-count])
                count -= 1
                
        return max(count,ret)
