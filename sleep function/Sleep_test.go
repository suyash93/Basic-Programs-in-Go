package main
import "testing"
import "time"
import "fmt"
func Test123(t *testing.T) {
	fmt.Println("1")
	starttime:=time.Now()
	Sleep(1)
	elasped:= time.Since(starttime)
	if elasped<1 {
		t.Error("Expected 1, got " ,elasped)
	}
	}
