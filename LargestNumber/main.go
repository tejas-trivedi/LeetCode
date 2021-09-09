func largestNumber(nums []int) string {
    var number string
    sort.Slice(nums, func(i, j int) bool {
        str1 := strconv.Itoa(nums[i])
        str2 := strconv.Itoa(nums[j])

        if str1+str2 >= str2+str1 {
            return true
        }
        return false
    })

    for _, num := range nums {
        number += strconv.Itoa(num)
    }
    if number[0] == '0' {
        return number[:1]
    }
    return number
}
