package main

import (
	"MetaNode/stage1"
	"fmt"
)

func main() {
	// single number
	//fmt.Println(stage1.SingleNumber1([]int{4, 1, 2, 1, 2}))
	//fmt.Println(stage1.SingleNumber2([]int{6, 1, 2, 1, 2}))

	// is palindrome
	//fmt.Println(stage1.IsPalindrome(12321))
	//fmt.Println(stage1.IsPalindrome(1221))
	//fmt.Println(stage1.IsPalindrome(-12321))
	//fmt.Println(stage1.IsPalindrome(12345))

	// is valid
	//fmt.Println(stage1.IsValid2("()[]{}"))
	//fmt.Println(stage1.IsValid2("(]]"))
	//fmt.Println(stage1.IsValid2("([])"))
	//fmt.Println(stage1.IsValid2("([)]"))

	// longest common prefix
	//fmt.Println(stage1.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	//fmt.Println(stage1.LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	//fmt.Println(stage1.LongestCommonPrefix([]string{"doorway", "door", "doomed"}))

	// plus one
	//fmt.Println(stage1.PlusOne([]int{1, 2, 3}))
	//fmt.Println(stage1.PlusOne([]int{4, 3, 2, 1}))
	//fmt.Println(stage1.PlusOne([]int{9, 9}))

	// remove duplicates
	//int1 := []int{1, 1, 2}
	//fmt.Println(stage1.RemoveDuplicates(&int1), int1)
	//int2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//fmt.Println(stage1.RemoveDuplicates(&int2), int2)
	//int3 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 6}
	//fmt.Println(stage1.RemoveDuplicates(&int3), int3)

	// merge
	//arr1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//fmt.Println(stage1.Merge(arr1))
	//arr2 := [][]int{{1, 4}, {4, 5}}
	//fmt.Println(stage1.Merge(arr2))
	//arr3 := [][]int{{4, 7}, {1, 4}}
	//fmt.Println(stage1.Merge(arr3))

	//two sum
	fmt.Println(stage1.TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(stage1.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println(stage1.TwoSum([]int{3, 3}, 6))

}
