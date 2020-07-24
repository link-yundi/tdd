package ch05_struct_method_interface

type Triangle struct {
	Base   float64
	Height float64
}

func (triangle Triangle) Area() (area float64) {
	return 0.5 * triangle.Base * triangle.Height
}
