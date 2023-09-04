func hashmap(nums []int) bool {
    seen := map[int]struct{}{}

    for _, v := range nums {
        if _, prs := seen[v]; prs {
            return true
        }
        seen[v] = struct{}{}
    }

    return false
}

func sorts(nums []int) bool {
    sort.Ints(nums)
    prev := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] == prev {
            return true
        }

        prev = nums[i]
    }

    return false
}

func containsDuplicate(nums []int) bool {
    return hashmap(nums)
    // return sorts(nums)
}

