package stage1

func SingleNumber(nums []int) int {
	var single int
	for _, num := range nums {
		single ^= num
	}
	return single
}
