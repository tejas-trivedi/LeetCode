class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int n = nums.size();

        int sum = accumulate(nums.begin(),nums.end(),0);
        int sumOnNum = (n)*(n+1)/2;
        // cout<<sum<<endl;
        // cout<<sumOnNum;
        return sumOnNum - sum;

    }
};