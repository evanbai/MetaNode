package stage2

type Shape interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
}
