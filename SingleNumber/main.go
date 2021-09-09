func singleNumber(nums []int) int {
    n := len(nums);
    sum := 0;
    for i:=0; i<n; i++ {
        sum = sum ^ nums[i]
    }
    return sum;
}