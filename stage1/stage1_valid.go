package stage1

import (
	"fmt"
)

func Stage1Valid() {
	//single number
	fmt.Println("----single number----")
	fmt.Println(SingleNumber1([]int{4, 1, 2, 1, 2}))
	fmt.Println(SingleNumber2([]int{6, 1, 2, 1, 2}))

	//is palindrome
	fmt.Println("----is palindrome----")
	fmt.Println(IsPalindrome(12321))
	fmt.Println(IsPalindrome(1221))
	fmt.Println(IsPalindrome(-12321))
	fmt.Println(IsPalindrome(12345))

	//is valid
	fmt.Println("----is valid----")
	fmt.Println(IsValid1("()[]{}"))
	fmt.Println(IsValid1("(]]"))
	fmt.Println(IsValid2("([])"))
	fmt.Println(IsValid2("([)]"))

	//longest common prefix
	fmt.Println("----longest common prefix----")
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(LongestCommonPrefix([]string{"doorway", "door", "doomed"}))

	//plus one
	fmt.Println("----v----")
	fmt.Println(PlusOne([]int{1, 2, 3}))
	fmt.Println(PlusOne([]int{4, 3, 2, 1}))
	fmt.Println(PlusOne([]int{9, 9}))

	//remove duplicates
	fmt.Println("----remove duplicates----")
	int1 := []int{1, 1, 2}
	fmt.Println(RemoveDuplicates(&int1), int1)
	int2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(RemoveDuplicates(&int2), int2)
	int3 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 6}
	fmt.Println(RemoveDuplicates(&int3), int3)

	//merge
	fmt.Println("----merge----")
	arr1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(Merge(arr1))
	arr2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println(Merge(arr2))
	arr3 := [][]int{{4, 7}, {1, 4}}
	fmt.Println(Merge(arr3))

	//two sum
	fmt.Println("----two sum----")
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(TwoSum([]int{3, 3}, 6))

}
