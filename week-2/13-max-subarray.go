func maxSubArray(nums []int) int {
    runningSum, maxSum := nums[0], nums[0]

    for i := 1; i < len(nums); i++ {
        if runningSum + nums[i] > nums[i] {
            runningSum += nums[i]
        } else {
            runningSum = nums[i]
        }

        if runningSum > maxSum {
            maxSum = runningSum
        }
    }

    return maxSum
}

