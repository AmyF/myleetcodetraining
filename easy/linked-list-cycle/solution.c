/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
bool hasCycle(struct ListNode *head) {
    if (head == NULL) return false;
    struct ListNode *p = head;
    struct ListNode *q = head;
    while (q->next && p->next) {
        p = p->next;
        if (q->next->next == NULL) {
            return false;
        }
        q = q->next->next;
        if (p == q) {
            return true;
        }
    }
    return false;
}