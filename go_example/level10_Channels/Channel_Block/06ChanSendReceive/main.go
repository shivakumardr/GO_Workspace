package main

import "fmt"

func main() {
	c := make(chan int, 2)
	cr := make(chan<- int) //reciver
	cs := make(<-chan int) //sender

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
