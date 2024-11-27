package main

//寻找最长不含有重复字符的字串
func LengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if last, ok := lastOccurred[ch]; ok && last >= start {
			start = last + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i

	}
	return maxLength
}

func main() {

}
