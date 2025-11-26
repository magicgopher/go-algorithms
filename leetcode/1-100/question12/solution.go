package question12

// 贪心
func intToRoman(num int) string {
	// 从大到小排列，所有可能出现的数值和对应的罗马符号
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string
	for i := 0; i < len(values); i++ {
		// 贪心：能减多少次就减多少次
		for num >= values[i] {
			result += symbols[i]
			num -= values[i]
		}
	}
	return result
}

// 按千、百、十、个位分别处理
func intToRoman1(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return thousands[num/1000] +
		hundreds[(num%1000)/100] +
		tens[(num%100)/10] +
		ones[num%10]
}
