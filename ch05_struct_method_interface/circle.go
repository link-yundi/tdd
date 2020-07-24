package ch05_struct_method_interface

import "math"

type Circle struct {
	Radius float64
}


func (circle Circle)Area() (area float64) {
	return math.Pi * circle.Radius * circle.Radius
}