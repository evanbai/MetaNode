package stage2

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	fmt.Println("Calculating area of rectangle")
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	fmt.Println("Calculating perimeter of rectangle")
	return 2 * (r.Width + r.Height)
}
func (c Circle) Area() float64 {
	fmt.Println("Calculating area of circle")
	return 3.14 * float64(c.Radius) * float64(c.Radius)
}

func (c Circle) Perimeter() float64 {
	fmt.Println("Calculating perimeter of circle")
	return 2 * 3.14 * float64(c.Radius)
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %s\n", e.Name, e.Age, e.EmployeeID)
}
