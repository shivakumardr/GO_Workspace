package main

import "fmt"
func main(){
	cs := make(chan int)
	//cr:=make(<-chan int)
	go func(){
		cs<-42
	}()
	fmt.Println(<-cs)
	fmt.Printf("cs:\t%T\n",cs)
	//fmt.Printf("cr:\t%T\n",cr)
}
