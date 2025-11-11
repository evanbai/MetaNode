package stage1

import (
	"fmt"
	"strconv"
	"strings"
)

func PlusOne(digits []int) []int {

	var builder strings.Builder
	for _, d := range digits {
		builder.WriteString(strconv.Itoa(d))
	}

	var digitInt int
	if s, err := strconv.Atoi(builder.String()); err == nil {
		digitInt = s
	} else {
		fmt.Println(err)
		return digits
	}
	digitInt++
	result := []int{}
	for _, char := range strconv.Itoa(digitInt) {
		if s, er := strconv.Atoi(string(char)); er == nil {
			result = append(result, s)
		} else {
			fmt.Println(er)
			return result
		}
	}

	return result
}
