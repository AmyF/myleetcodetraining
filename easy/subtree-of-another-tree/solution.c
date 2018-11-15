/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
bool isSame(struct TreeNode* s, struct TreeNode* t);

bool isSubtree(struct TreeNode* s, struct TreeNode* t) {
    if (s == NULL) {
        return false;
    }
    if (isSame(s,t)) {
        return true;
    }
    return isSubtree(s->left, t) || isSubtree(s->right, t);
}

bool isSame(struct TreeNode* s, struct TreeNode* t) {
    if (s == NULL && t == NULL) {
        return true;
    }
    if (s == NULL || t == NULL) {
        return false;
    }
    return (s->val == t->val) && isSame(s->left, t->left) && isSame(s->right, t->right);
}