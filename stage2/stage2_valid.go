package stage2

import "fmt"

func Stage2Valid() {

	// Pointers
	//fmt.Println("----pointers----")
	//nums := 9
	//pointers1(&nums)
	//fmt.Println(nums)
	//numsArr := []int{4, 6, 9, 12}
	//pointers2(&numsArr)
	//fmt.Println(numsArr)

	// Goroutine
	//fmt.Println("----goroutine----")
	//ExeGoroutine()

	// Multiple task schedule
	// fmt.Println("----multiple task schedule----")
	// callFunc()

	// Object oriented programming
	// fmt.Println("----object oriented programming----")
	// rect := Rectangle{Width: 10, Height: 5}
	// circle := Circle{Radius: 7}
	// shapes := []Shape{rect, circle}
	// for _, shape := range shapes {
	// 	fmt.Printf("Area: %f\n", shape.Area())
	// 	fmt.Printf("Perimeter: %f\n", shape.Perimeter())
	// }
	// emp := Employee{
	// 	Person:     Person{Name: "Alice", Age: 30},
	// 	EmployeeID: "E12345",
	// }
	// emp.PrintInfo()

	// Channel
	// fmt.Println("----channel----")
	// chanMethod1()
	// chanMethod2()

	// Mutex
	fmt.Println("----mutex----")
	lockMethod1()
	lockMethod2()

}
