package main

import "fmt"
type person struct{
	First string
//	Last string
//	Age int
//	Saying []string
}
func (p *person)speak(){
	fmt.Println("Hello")
}
type human interface {
	speak()
}
func saySomething(h human){
	h.speak()
}
func main(){
	p1:= person{
		First: "james",
	}
	saySomething(&p1)
	//This don't work
//	saySomething(p1)
	p1.speak()
}