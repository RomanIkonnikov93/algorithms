package easy

/*
Вам дан массив цен, где цены[i] — цена данной акции на i-й день.
Вы хотите максимизировать свою прибыль, выбрав один день для покупки одной акции и выбрав другой день в будущем для продажи этой акции.
Верните максимальную прибыль, которую вы можете получить от этой сделки. Если вы не можете получить никакой прибыли, верните 0.
*/

/*
Input: []int{7,1,5,3,6,4}
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Input: []int{7,6,4,3,1}
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
*/

func MaxProfit(prices []int) int {
	var profit = 0
	var minPrice = prices[0]

	for i := 1; i < len(prices); i++ {
		// If we find any price which is lower than the current minPrice
		// update the minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if (prices[i] - minPrice) > profit {
			// If diff of current stock with minPrice is greater
			// update the profit
			profit = prices[i] - minPrice
		}
	}

	return profit
}
