/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* insertionSortList(ListNode* head) {
        if (head == NULL || head->next == NULL) return head;
        ListNode sortedList(0);
        ListNode* cur = head;
        while (cur) {
            ListNode* next = cur->next;
            ListNode* node = &sortedList;
            while (node->next != nullptr && node->next->val < cur->val) {
                node = node->next;
            }
            cur->next = node->next;
            node->next = cur;
            cur = next;
        }
        return sortedList.next;
    }
};