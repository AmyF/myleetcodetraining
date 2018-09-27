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
    vector<int> preorder(Node* root) {
        vector<int> v;
        if (root == nullptr) return v; 
        myPreOrder(root, v);
        return v;
    }
    void myPreOrder(Node* root, vector<int> &v) {
        v.push_back(root->val);
        for (int i=0;i<root->children.size();i++) {
            if (root->children[i]) {
                myPreOrder(root->children[i], v);
            }
        }
    }
};
