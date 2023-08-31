func maxProfit(prices []int) int {
    maxProf := 0
    minBuy := prices[0]

    for i := 1; i < len(prices); i++ {
        if sell := prices[i]; sell < minBuy {
            minBuy = sell

        } else if profit := prices[i] - minBuy; profit > maxProf {
            maxProf = profit
        }
    }

    return maxProf
}

