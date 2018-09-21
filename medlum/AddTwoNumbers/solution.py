# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        v1, p1 = 0, 1
        v2, p2 = 0, 1
        while l1 != None:
            v1 += l1.val * p1
            p1 *= 10
            l1 = l1.next
        while l2 != None:
            v2 += l2.val * p2
            p2 *= 10
            l2 = l2.next
        (v,p) = divmod(v1+v2,10)
        root = ListNode(p)
        node = root
        while v > 0:
            (v,p) = divmod(v,10)
            node.next = ListNode(p)
            node = node.next
        return  root
