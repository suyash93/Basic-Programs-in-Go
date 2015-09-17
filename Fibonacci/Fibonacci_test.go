package main 
import "testing"
func TestFib(t *testing.T) {
	var in int
	in= Fib(7)
	if in!=13 {
		t.Error("Expected 13, got ", in)
	}
}