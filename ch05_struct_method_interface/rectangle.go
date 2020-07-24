package ch05_struct_method_interface

type Rectangle struct {
	Height float64
	Width float64
}

func (rectangle Rectangle) Area() (area float64) {
	return rectangle.Height * rectangle.Width
}

