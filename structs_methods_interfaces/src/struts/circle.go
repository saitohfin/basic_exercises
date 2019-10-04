package struts

import "math"

//Circle shape with a radius
type Circle struct{
	Radius float64
}

//Area calculate area of the circle
func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}