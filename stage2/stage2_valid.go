package stage2

import "fmt"

func Stage2Valid() {

	// pointers
	//fmt.Println("----pointers----")
	//nums := 9
	//pointers1(&nums)
	//fmt.Println(nums)
	//
	//numsArr := []int{4, 6, 9, 12}
	//pointers2(&numsArr)
	//fmt.Println(numsArr)

	// goroutine
	//fmt.Println("----goroutine----")
	//ExeGoroutine()

	// multiple task schedule
	// fmt.Println("----multiple task schedule----")
	// callFunc()

	// object oriented programming
	fmt.Println("----object oriented programming----")
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	shapes := []Shape{rect, circle}

	for _, shape := range shapes {
		fmt.Printf("Area: %f\n", shape.Area())
		fmt.Printf("Perimeter: %f\n", shape.Perimeter())
	}

}
