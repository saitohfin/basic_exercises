package struts

//Perimeter calculate the perimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.witdh + rectangle.height)
}

//Area calculate the area of a rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.witdh * rectangle.height
}
