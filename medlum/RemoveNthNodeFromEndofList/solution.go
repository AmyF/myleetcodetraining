/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    tmp := head
    count := 0
    for tmp != nil {
        count += 1
        tmp = tmp.Next 
    }
    if count - n == 0 {
        if head != nil {
            return head.Next
        } else {
            return nil
        }
    }
    tmp = head
    for i:=(count-n);i>1;i-- {
        tmp = tmp.Next
    }
    tmp.Next = tmp.Next.Next
    return head
}