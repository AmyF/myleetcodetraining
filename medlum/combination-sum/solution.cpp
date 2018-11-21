class Solution {
public:
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        sort(candidates.begin(), candidates.end());
        vector<vector<int>> res;
        vector<int> tmp;
        findCombination(res, candidates, tmp, target, 0);
        return res;
    }
    
    void findCombination(vector<vector<int>>& res, 
                         vector<int> candidates, 
                         vector<int> tmp,
                         int target, int pos) 
    {
        if (target == 0) {
            res.push_back(tmp);
        }
        else {
            for (int i = pos; i < candidates.size(); i++) {
                if (candidates[i] > target) break;
                tmp.push_back(candidates[i]);
                findCombination(res, candidates, tmp, target - candidates[i], i);
                tmp.pop_back();
            }
        }
    }
};