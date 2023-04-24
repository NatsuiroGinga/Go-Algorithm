package dfs

var letterMap = map[byte]string{
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var (
	s   []byte
	ret []string
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	s = make([]byte, 0)
	ret = make([]string, 0)
	backtrack(digits, 0)
	return ret
}

func backtrack(digits string, cur int) {
	if len(digits) == cur {
		ret = append(ret, string(s))
		return
	}
	digit := digits[cur]
	letter := letterMap[digit]

	for i := 0; i < len(letter); i++ {
		s = append(s, letter[i])
		backtrack(digits, cur+1)
		s = s[:len(s)-1]
	}
}
