class Solution(object):
    def tr(self, x, f):
        v = 0
        while x > 0:
            (x,p) = divmod(x,10)
            v = v * 10 + p
        if f and v > ((1 << 31) - 1):
            return 0
        elif f == False and v > (1 << 31):
            return 0
        return v
    
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        if x >= 0:
            return self.tr(x,True)
        else:
            return - self.tr(-x,False)
        
    
