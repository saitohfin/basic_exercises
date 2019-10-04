package struts

//Rectangle is a shape with four sides and different width and height
type Rectangle struct {
	Witdh  float64
	Height float64
}

//Area calculate the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Witdh * r.Height
}
