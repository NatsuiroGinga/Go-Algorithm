package dp

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{2}
	amount := 3

	t.Log(coinChange(coins, amount))
}
