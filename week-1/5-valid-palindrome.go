func isNotAlphanumeric(b byte) bool {
    return b < '0' || (b > '9' && b < 'A') || (b > 'Z' && b < 'a') || b > 'z'
}

func isPalindrome(s string) bool {

    s = strings.ToLower(s)
    left, right := 0, len(s)-1

    var (
        rchar byte
        lchar byte
    )
    for left < right {
        rchar = s[right]
        for isNotAlphanumeric(rchar) {
           right-- 
           if right < left {
               return true
           }
           rchar = s[right]
        }

        lchar = s[left]
        for isNotAlphanumeric(lchar) {
           left++ 
           if left > right {
               return false
           }
           lchar = s[left]
        }

        if lchar != rchar {
            return false
        }

        left++
        right--
    }

   return true 
}

