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
	for i:=0;i<100;i++{
		c <- i
	}
	//close(c)
}
//receiver
func receiver(c <-chan int) {
	for v:=range c{
		fmt.Println(v)
	}

}
