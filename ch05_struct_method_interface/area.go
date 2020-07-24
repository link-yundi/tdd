package ch05_struct_method_interface

func Area(rectangle Rectangle) (area float64) {
	return rectangle.Width * rectangle.Height
}
