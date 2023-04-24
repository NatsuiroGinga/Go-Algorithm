package dp

import "testing"

func TestWordBreak(t *testing.T) {
	s := "cars"
	wordDict := []string{"car", "ca", "rs"}
	t.Log(wordBreak(s, wordDict))
}
