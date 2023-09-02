func climbStairs(n int) int {
    if n <= 3 {
        return n
    }

    prev, prev2 := 3, 2
    for i := 4; i <= n; i++ {
        prev, prev2 = prev + prev2, prev
    }

    return prev
    
}

