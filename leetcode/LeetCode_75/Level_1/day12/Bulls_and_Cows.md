# LeetCode 299. Bulls and Cows
## 题目描述
You are playing the Bulls and Cows game with your friend.

You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess, you provide a hint with the following info:

The number of "bulls", which are digits in the guess that are in the correct position.
The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.
Given the secret number secret and your friend's guess guess, return the hint for your friend's guess.

The hint should be formatted as "xAyB", where x is the number of bulls and y is the number of cows. Note that both secret and guess may contain duplicate digits.
## 思路
1. 同时遍历secret和guess, 使用一个哈希表fre[byte]int来记录secret和guess中每个数字出现的次数
2. 如果secret[i] == guess[i]，则bulls++
3. 否则, secret[i]出现的次数加1，guess[i]出现的次数减1
4. 同时判断, secret[i]出现的次数是否小于0，因为如果小于0，说明之前出现过，在第3步中已经减过1了，所以如果小于0，说明guess[i]出现过，cows++
5. 同理, guess[i]出现的次数是否大于0，如果是，则cows++
6. 最后返回bulls和cows的字符串

`例如`: secret = "1123", guess = "0111"
1. secret[0] != guess[0], 因为fre['1'] = fre['0'] = 0, 所以不进入if, 最后fre中数据为: fre['1'] = 1, fre['0'] = -1, bulls = 0, cows = 0
2. secret[1] == guess[1], 此时fre中数据为同上, bulls = 1, cows = 0
3. **重点**: secret[2] != guess[2], 因为fre['1'] = 1 > 0, 所以进入第二个if, cows++, **然后** fre['1']--, fre['2']++, 最后fre中数据为: fre['1'] = 0, fre['2'] = 1, bulls = 1, cows = 1
4. secret[3] != guess[3], 此时fre['3'] = 0, fre['1'] = 0, 所以不进入任何if, 最后fre中数据为: fre['1'] = 0, fre['3'] = 1, bulls = 1, cows = 1

## 代码
```go
func getHint(secret string, guess string) string {
    fre := make(map[byte]int, 10)
    bulls, cows := 0, 0
	
    for i := 0; i < len(secret); i++ {
        if secret[i] == guess[i] {
            bulls++
        } else {
            if fre[secret[i]] < 0 {
                cows++
            }
            if fre[guess[i]] > 0 {
                cows++
            }
            fre[secret[i]]++
            fre[guess[i]]--
        }
    }
    return fmt.Sprintf("%dA%dB", bulls, cows)
}
```