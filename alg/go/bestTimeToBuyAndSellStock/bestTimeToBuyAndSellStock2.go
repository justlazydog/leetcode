package bestTimeToBuyAndSellStock

func maxProfit2(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	minPrice := prices[0]
	res := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > res {
				res = profit
			}
		}
	}

	return res
}
