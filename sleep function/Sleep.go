package main
import "fmt"
import "time"
 
func pinger(c chan string) {
	for i := 0; ; i++ {
		c<-"ping"
	}	
}

func printer(c chan string) {
	for{
		fmt.Println(<-c)
		Sleep(1)
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c<-"pong"
	}
}

func main() {
	var c chan string= make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}


func Sleep(x int) {
    <-time.After(time.Second * time.Duration(x))
}