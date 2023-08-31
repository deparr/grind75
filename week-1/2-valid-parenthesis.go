func stack(s string) bool {
    stk := make([]rune, len(s))
    sidx := 0

    for _, c := range s {
        switch c {
            case '(', '[', '{':
                stk[sidx] = c
                sidx++
            case ')':
                if sidx==0 || stk[sidx-1] != '(' {
                    return false
                }
                sidx--
            case ']':
                if sidx==0 || stk[sidx-1] != '[' {
                    return false
                }
                sidx--
            case '}':
                if sidx==0 || stk[sidx-1] != '{' {
                    return false
                }
                sidx--
        }
    }

    return sidx == 0

}

func isValid(s string) bool {
    if len(s) % 2 != 0 {
        return false
    }
    return stack(s)
}

