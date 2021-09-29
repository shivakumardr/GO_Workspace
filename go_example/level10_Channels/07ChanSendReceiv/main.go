package main

import "fmt"

func main() {
	c := make(chan int)
	//send
	go sender(c)
	//receive
	receiver(c)
	fmt.Println("About to exit")

}

//send
func sender(c chan<- int) {
	c <- 42
}
//receiver
func receiver(c <-chan int) {
	fmt.Println(<-c)
}
