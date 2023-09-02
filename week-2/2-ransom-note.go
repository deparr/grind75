// can also use an array for cache since domain is known

func canConstruct(ransomNote string, magazine string) bool {
    letters := map[rune]int{}

    for _, c := range magazine {
        letters[c] += 1
    }

    for _, c := range ransomNote {
        if count, prs := letters[c]; !prs || count == 0 {
            return false
        }
        
        letters[c]--
    }

    return true
}

