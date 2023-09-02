func longestPalindrome(s string) int {
    letters := map[rune]int{}

    for _, ru := range s {
        letters[ru]++
    }

    len := 0
    var hasOdd bool
    for _, count := range letters {
        if count % 2 == 1 {
            hasOdd = true
            len += count - 1
        } else {
            len += count
        }
    }

    if hasOdd {
        len += 1
    }
    return len
}

