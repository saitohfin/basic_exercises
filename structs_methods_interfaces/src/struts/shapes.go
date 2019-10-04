package struts

//Perimeter calculate the perimeter of a rectangle
func Perimeter(witdh, height float64) float64 {
	return 2 * (witdh + height)
}

//Area calculate the area of a rectangle
func Area(witdh, height float64) float64 {
	return witdh * height
}
