package main
import "testing"
func Test123(t *testing.T) {
	r:= rectangle{length: 4, breadth: 6}
	
	if (r.perimeter() != 20){
	t.Error("Expected 20, got " , r.perimeter() )
	}
	if(r.area() != 24) {
		t.Error("Expected 24, got ", r.area())
	}
}
