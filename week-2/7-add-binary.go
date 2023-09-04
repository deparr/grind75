// try doing a simpler loop and just checking index valiidity
// each time
//
// for alen >= 0 || blen >= 0 {
//   if alen >=0 {
//    ...
//   }
// }


func addBinary(a string, b string) string {
    long, short := a, b
    if len(b) > len(long) {
        long = b
        short = a
    }

    var carry byte = '0'
    lendiff := len(long) - len(short)
    result := make([]byte, len(long)+1)
    resultIdx := len(result)-1
    i := len(long)-1
    for ; i >= lendiff; i-- {
        sum := carry + long[i] + short[i-lendiff]
        var nextDigit byte

        switch sum {
            // all zero
            case 144:
                nextDigit = '0' 
                carry = '0'
            // one '1'
            case 145:
                nextDigit = '1'
                carry = '0'
            // two '1'
            case 146:
                nextDigit = '0'
                carry = '1'
            // three '1'
            case 147:
                nextDigit = '1'
                carry = '1'
        }

        result[resultIdx] = nextDigit
        resultIdx--
    }

    if carry == '1' {

        for ; i > -1 && long[i] == '1'; i-- {
            result[resultIdx] = '0'
            resultIdx--
        }
        result[resultIdx] = '1'
        resultIdx--
        i--
    }
    fmt.Println(i, result)
    for ; i > -1; i-- {
        result[resultIdx] = long[i]
        resultIdx--
    }
    return string(result[resultIdx+1:])
}

