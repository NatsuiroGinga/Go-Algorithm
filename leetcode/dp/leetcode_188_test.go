package dp

import "testing"

func TestMaxProfit4(t *testing.T) {
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 2
	t.Log(maxProfit4(k, prices))
}
