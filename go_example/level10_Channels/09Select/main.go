package main

import "fmt"

func main() {
	//Channels
	evn := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	//sender
	go send(evn, odd, quit)
	//receiver
	receiver(evn, odd, quit)
	fmt.Println("About to exit")
}
func receiver(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From Even channel\t:", v)
		case v := <-o:
			fmt.Println("Form Odd channel\t:", v)
		case v := <-q:
			fmt.Println("Form Quit channel\t:", v)
			return
		}
	}
}
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	q <- 0
}
