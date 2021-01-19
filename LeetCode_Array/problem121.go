func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	minPrice,maxProfit := prices[0],0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i]-minPrice > maxProfit) {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}