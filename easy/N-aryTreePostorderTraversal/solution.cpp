/*
// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Solution {
public:
    vector<int> postorder(Node* root) {
        vector<int> v;
        if (root == nullptr) return v;
        myPostOrder(root,v);
        return v;
    }
    void myPostOrder(Node* root, vector<int> &v) {
        if (root->children.size() == 0) {
            v.push_back(root->val);
            return;
        }
        for (int i=0;i<root->children.size();i++) {
            myPostOrder(root->children[i], v);
        }
        v.push_back(root->val);
        return;
    }
};
