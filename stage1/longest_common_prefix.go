package stage1

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var key rune
	var i int = 0
	var lens int = 0
	var comp bool = true
	var smallLen int = len(strs[0])
	for ; comp && i < smallLen; i++ {
		for _, str := range strs {
			if len(str) < smallLen {
				smallLen = len(str)
			}

			if key == 0 {
				key = []rune(str)[i]
			} else if key != []rune(str)[i] {
				comp = false
			}
		}
		if comp {
			lens++
			key = 0
		}
	}
	return string([]rune(strs[0])[:lens])
}
