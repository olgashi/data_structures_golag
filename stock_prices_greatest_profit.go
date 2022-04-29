package main

import (
	"fmt"
)

func find_greatest_profit(stockPrices []int) int {
// [3, 1, 4, 5, 10, 3, 1]

min := stockPrices[0]

largestProfit := 0;
index := 2;
for (index < len(stockPrices)) {
  if (stockPrices[index - 1] < min) {
		min = stockPrices[index - 1]
	}

	if largestProfit < stockPrices[index] - min {
		largestProfit = stockPrices[index] - min
	}
	index++
}

return largestProfit

}

func main() {
	fmt.Println(find_greatest_profit([]int{3, 1, 4, 5, 10, 3, 1}))
}