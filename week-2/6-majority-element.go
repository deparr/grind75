// Should also look at the linear time solutions,
// pretty tricky

func majorityElement(nums []int) int {
    sort.Ints(nums)

    count, prev := 1, nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] != prev {
            if count > len(nums) / 2 {
                return prev
            }
            count, prev = 1, nums[i]
        } else {
            count++
        }
    }

    if count > len(nums) / 2 {
        return nums[len(nums)-1]
    }

    return prev
}

