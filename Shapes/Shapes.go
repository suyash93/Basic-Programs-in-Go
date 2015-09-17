package main
import ("fmt"; "math")
type Shape interface {
	area () float64
	perimeter() float64
}
type circle struct {
	radius float64

}
type rectangle struct {
	length ,breadth float64

}
func (c circle) area() float64 {
	return math.Pi*c.radius*c.radius
}
func (c circle) perimeter() float64 {
	return 2*math.Pi*c.radius
}
func (r rectangle) area() float64 {
	return r.length*r.breadth
}
func (r rectangle) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}
func  takeshape (s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}
func main() {
	c:= circle{radius: 5}
	r:= rectangle{length: 5, breadth: 4}
	takeshape(c)
	takeshape(r)
}
