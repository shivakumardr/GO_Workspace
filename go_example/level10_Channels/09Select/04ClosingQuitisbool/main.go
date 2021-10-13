package main

import "fmt"

func main() {
	evn := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)
	//sender
	go send(evn, odd, quit)
	//receiver
	receiver(evn, odd, quit)
	fmt.Println("About to exit")
}
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
		//	close(e)
		//	close(o)
	}
	//q<-true
	close(q)

}
func receiver(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("Value receive from EVEN channel:", v)
		case v := <-o:
			fmt.Println("Value receive from ODD channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok ", i, ok)
				return
			} else {
				fmt.Println("from comma ok ", i)
			}
		}
	}
}
