package stage1

func IsValid1(s string) bool {
	var al, ar, bl, br, cl, cr string = "(", ")", "[", "]", "{", "}"
	runes := []rune(s)

	for i := 0; i < len(runes); {
		switch string(runes[i]) {
		case al:
			if string(runes[i+1]) == ar {
				runes = append(runes[:i], runes[i+2:]...)
				i = 0
			} else {
				i++
			}
		case bl:
			if string(runes[i+1]) == br {
				runes = append(runes[:i], runes[i+2:]...)
				i = 0
			} else {
				i++
			}
		case cl:
			if string(runes[i+1]) == cr {
				runes = append(runes[:i], runes[i+2:]...)
				i = 0
			} else {
				i++
			}
		default:
			i++
		}
	}
	return len(runes) == 0
}

func IsValid2(s string) bool {
	// 用切片模拟栈
	stack := []rune{}
	// 定义括号匹配映射
	match := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 左括号入栈
			stack = append(stack, char)
		case ')', '}', ']':
			// 右括号：检查栈是否为空或栈顶不匹配
			if len(stack) == 0 || stack[len(stack)-1] != match[char] {
				return false
			}
			// 匹配则出栈
			stack = stack[:len(stack)-1]
		}
	}

	// 栈为空说明所有左括号都被匹配
	return len(stack) == 0
}
