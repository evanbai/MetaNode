package stage1

func RemoveDuplicates(nums *[]int) int {

	orgNums := make([]int, len(*nums))
	copy(orgNums, *nums)
	bMap := make(map[int]struct{})
	var keys int = 0
	for i := 0; i < len(*nums); {
		if i == 0 {
			keys = (*nums)[i]
			i++
		} else if keys == (*nums)[i] {
			*nums = append((*nums)[:i], (*nums)[i+1:]...)
			bMap[keys] = struct{}{}
		} else {
			keys = (*nums)[i]
			i++
		}
	}

	var result []int
	for _, v := range orgNums {
		if _, exists := bMap[v]; !exists {
			result = append(result, v)
		}
	}

	return len(result)
}
