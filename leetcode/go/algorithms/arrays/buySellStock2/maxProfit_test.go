package buySellStock2

import "testing"

/*
122. Best Time to Buy and Sell Stock II
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

Say you have an array prices for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.

Example 2:

Input: [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
             Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
             engaging multiple transactions at the same time. You must sell before buying again.

Example 3:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.



Constraints:

    1 <= prices.length <= 3 * 10 ^ 4
    0 <= prices[i] <= 10 ^ 4
*/

/*
Time Complexity: O(n) - we must iterate through all the prices once
Space Complexity: O(1) - we do not require extra space proportional to n, just an int to track profit and an int to record buyPrice
*/

func TestMaxProfit(t *testing.T) {
	t.Run("should maximize profit in example 1", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}

		expected := 7
		actual := maxProfit(prices)

		if actual != expected {
			t.Errorf("Got %d, expected %d", actual, expected)
		}
	})

	t.Run("should maximize profit in example 2", func(t *testing.T) {
		prices := []int{1, 2, 3, 4, 5}

		expected := 4
		actual := maxProfit(prices)

		if actual != expected {
			t.Errorf("Got %d, expected %d", actual, expected)
		}
	})

	t.Run("should maximize profit in example 3 where no profit is possible", func(t *testing.T) {
		prices := []int{7, 6, 4, 3, 1}

		expected := 0
		actual := maxProfit(prices)

		if actual != expected {
			t.Errorf("Got %d, expected %d", actual, expected)
		}
	})
}

func maxProfit(prices []int) int {
	profit := 0
	buyPrice := -1

	for i := 0; i < len(prices)-1; i++ {
		priceToday := prices[i]
		priceTomorrow := prices[i+1]

		if buyPrice == -1 && priceToday < priceTomorrow {
			buyPrice = priceToday
			continue
		}

		if buyPrice != -1 && priceToday > priceTomorrow {
			profit += priceToday - buyPrice
			buyPrice = -1
			continue
		}
	}

	if buyPrice != -1 {
		profit += prices[len(prices)-1] - buyPrice
	}

	return profit
}
