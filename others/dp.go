package main

import (
	"fmt"
	"math"
)

/**
 * DP (最小化問題)
 *
 * 参考URL
 * https://qiita.com/drken/items/dc53c683d6de8aeacf5a
 */
func Flog(n int) (int, error) {
	var h []float64
	
	for i := 0; i < n; i++ {
		fmt.Printf("Please input cost.[%d]: ", i)
		var cost float64
		fmt.Scan(&cost)
		h = append(h, cost)
	}

	fmt.Println("%v", h)

	var dp []float64

	dp = append(dp, 0)

	for i := 1; i < len(h); i++ {
		dp = append(dp, dp[i - 1] + math.Abs(h[i] - h[i - 1]))
		if i > 1 {
			cost := dp[i - 2] + math.Abs(h[i] - h[i - 2])
			if dp[i] > cost {
				dp[i] = cost
			}
		}
	}

	fmt.Println("%v", dp)

	return -1, nil
}

func main() {
	Flog(5)
}