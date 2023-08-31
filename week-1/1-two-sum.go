func quadratic(nums []int, target int) []int {
    res := make([]int, 2)

    for i, val1 := range nums {
        for j, val2 := range nums {
            if i != j && val1 + val2 == target {
                res[0], res[1] = i, j
                return res
            }
        }
    }

    return res 
}

func linear(nums []int, target int) []int {
    cache := map[int]int{}
    res := make([]int, 2)

    for i, val := range nums {
        if j, prs := cache[target-val]; prs {
            res[0], res[1] = i, j
            return res
        }
        cache[val] = i
    }

    return res

}



func twoSum(nums []int, target int) []int {
    
    //return quadratic(nums, target)
    return linear(nums, target)
}

