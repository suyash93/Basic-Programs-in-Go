
package main
import "fmt"

func Fib (input int) int {
	if (input==0 || input==1) {
		return input
	}else{
		return (Fib(input-1) +Fib(input -2))}
}
func main() {
	fmt.Println("Enter the digit: ")
	var input int
	fmt.Scanf("%v", &input)
	fmt.Println( Fib(input))
}

