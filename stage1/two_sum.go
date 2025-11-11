package stage1

func TwoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
