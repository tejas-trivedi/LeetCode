func missingNumber(nums []int) int {

    n := len(nums);
    res := 0
    for i:=0; i<len(nums); i++ {
      res += nums[i]
   }
    totalSum := n*(n+1)/2;
    return totalSum - res;
}