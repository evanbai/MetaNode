package stage1

import "sort"

func Merge(intervals [][]int) [][]int {
	intArray := make([]int, 0)

	for _, interval := range intervals {
		var start int = interval[0]
		var end int = interval[1]
		for i := start; i <= end; i++ {
			intArray = append(intArray, i)
		}
	}

	sort.Slice(intArray, func(i, j int) bool {
		// 比较规则：若 nums[i] < nums[j]，则 i 位置的元素排在 j 前面（升序）
		return intArray[i] < intArray[j]
	})
	RemoveDuplicates(&intArray)
	result := make([][]int, 0)
	temp := make([]int, 0)
	for i := 0; i < len(intArray); i++ {
		if i == 0 {
			temp = append(temp, intArray[i])
		} else if intArray[i]-intArray[i-1] > 1 {
			temp = append(temp, intArray[i-1])
			result = append(result, temp)
			temp = make([]int, 0)
			temp = append(temp, intArray[i])
		} else if i == len(intArray)-1 {
			temp = append(temp, intArray[i])
			result = append(result, temp)
		}
	}
	return result
}
