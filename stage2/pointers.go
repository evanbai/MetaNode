package stage2

func pointers1(nums *int) {
	*nums += 10
}

func pointers2(nums *[]int) {
	for idx, _ := range *nums {
		(*nums)[idx] *= 2
	}
}
