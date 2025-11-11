package stage1

func IsPalindrome(x int) bool {
	// 特殊情况：负数或末尾为0的非零数一定不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversedHalf := 0
	// 反转后半部分数字
	for x > reversedHalf {
		// 取出x的最后一位，拼接到reversedHalf
		reversedHalf = reversedHalf*10 + x%10
		// 移除x的最后一位
		x /= 10
	}

	// 当数字长度为奇数时，reversedHalf会比x多一位（中间位），需除以10再比较
	return x == reversedHalf || x == reversedHalf/10
}
