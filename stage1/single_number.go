package stage1

func SingleNumber1(nums []int) int {
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := maps[nums[i]]; ok {
			maps[nums[i]] += 1
		} else {
			maps[nums[i]] = 1
		}
	}
	var result int
	for k, v := range maps {
		if v == 1 {
			result = k
		}
	}
	return result
}
func SingleNumber2(nums []int) int {
	var single int
	for _, num := range nums {
		single ^= num
	}
	return single
}
