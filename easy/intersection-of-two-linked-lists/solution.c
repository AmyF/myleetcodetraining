/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode *getIntersectionNode(struct ListNode *headA, struct ListNode *headB) {
    struct ListNode *p = headA;
    struct ListNode *q = headB;
    
    int lenp = 0;
    int lenq = 0;
    while (p) {
        p = p->next;
        lenp ++;
    }
    while (q){
        q = q->next;
        lenq ++;
    }
    
    p = headA;
    q = headB;
    
    if (lenp > lenq) {
        for (int i=0;i<(lenp-lenq);i++) {
            p = p->next;
        }
    }
    if (lenp < lenq) {
        for (int i=0;i<(lenq-lenp);i++) {
            q = q->next;
        }
    }
    
    while (p != q) {
        p = p->next;
        q = q->next;
    }
    return p;
}